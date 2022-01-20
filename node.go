package main

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
)

/* ---------------------------------------------------------------structure and methods used for building a Blockchain ---------------------------------------------------------------*/
// Structure of Block-data(data portion of a Block)
type BlockData struct {
	Title             string
	Description       string
	Owners            []string
	Problem           string
	Domain            string
	Technologies_used []string
	Viewing_price     float64
	Ownership_price   float64
	Pricing_history   []float64
}

// Same structure as the above, but in json format
type ideas struct {
	Ideas []BlockData `json:"ideas"`
}

// Structure of a Block
type Block struct {
	Data        BlockData
	PrevPointer *Block
	PrevHash    string
	CurrentHash string
}

// This method is used to print the whole Blockchain on console.
func ListBlocks(chainHead *Block) {
	var Temp *Block
	Temp = chainHead
	fmt.Print("\n")
	for ; Temp != nil; Temp = Temp.PrevPointer {
		fmt.Print(Temp.Data)
		fmt.Print("\n")
	}
}

// This method is used to create a Blockdata
func CreateBlockData(title string, desc string, owners string, prob string, dom string, tech []string, vp float64, op float64, ph float64) BlockData {
	var Idea BlockData
	Idea.Title = title
	Idea.Description = desc
	Idea.Owners = append(Idea.Owners, owners)
	Idea.Problem = prob
	Idea.Domain = dom
	Idea.Technologies_used = tech
	Idea.Viewing_price = vp
	Idea.Ownership_price = op
	Idea.Pricing_history = append(Idea.Pricing_history, ph)
	return Idea
}

// The head of the Blockchain i.e. The most recent block inserted
var mychain *Block

// Controls the whole program
var status_of_node = 1
var conn net.Conn
var err error

// This method always listens to the central node, so that the central node can send the updated Chain
func Always_listen_to_server(conn net.Conn) {
	for status_of_node != 0 {
		dec := gob.NewDecoder(conn)
		err := dec.Decode(&mychain)
		checkError(err)
		// fmt.Print("Updated Blockchain : ")
		// ListBlocks(mychain)
		// fmt.Print("\n\n\n")
	}
}

// Method to check error, if any
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

/* -------------------------------------------------------------------------------- main method --------------------------------------------------------------------------------*/
func main() {
	// Runs an http server on another thread
	go http_server()
	conn, err = net.Dial("tcp", "localhost:6000")
	checkError(err)
	// Constantly listen to server on another thread
	go Always_listen_to_server(conn)

	choice := 1
	gobEncoder := gob.NewEncoder(conn)
	for choice <= 1 {
		fmt.Print("------------------------------------------------------------------------------------------\n")
		fmt.Print("Press any number, greater than 1 to exit : \n")
		fmt.Print("------------------------------------------------------------------------------------------\n")
		fmt.Scan(&choice)
		err = gobEncoder.Encode(choice)
		checkError(err)
	}

	err = gobEncoder.Encode(10)
	checkError(err)
	status_of_node = 0
	conn.Close()
	fmt.Println("Connection closed.........................................")
	fmt.Println("Node is Closing................................")
}

/* --------------------------------------------------------------- methods used by http server---------------------------------------------------------------*/
// This method enables Cors, by which the front end in angular can send and recieve data to it.
func enableCors(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"
	if origin := r.Header.Get("Origin"); origin != "" {
		(*w).Header().Set("Access-Control-Allow-Origin", "*")
		(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		(*w).Header().Set("Access-Control-Allow-Headers", allowedHeaders)
		(*w).Header().Set("Access-Control-Expose-Headers", "Authorization")
	}
}

// Just to check the server
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server is up and running %s!", r.URL.Path[1:])
}

// Used by front-end to recieve data.
func ListIdeas(w http.ResponseWriter, r *http.Request) {
	enableCors(&w, r)
	B1 := mychain
	var list_ideas ideas

	for ; B1 != nil; B1 = B1.PrevPointer {
		list_ideas.Ideas = append(list_ideas.Ideas, B1.Data)
	}
	json.NewEncoder(w).Encode(list_ideas.Ideas)
}

// Used by front-end to send data.
func ProposeIdea(w http.ResponseWriter, r *http.Request) {
	enableCors(&w, r)
	json.NewEncoder(w).Encode("proposed successfully")
	fmt.Println("Request incoming!!!!")
	var res map[string]interface{}

	var title, description, owner, problem, domain string
	var view_p, own_p float64
	var t_u []string
	t_u = append(t_u, "MEAN STACK")
	t_u = append(t_u, "MERN STACK")
	json.NewDecoder(r.Body).Decode(&res)
	fmt.Println(res)
	resp := true

	if res["Title"] != nil {
		title = res["Title"].(string)
	} else {
		resp = false
	}

	if res["Description"] != nil {
		description = res["Description"].(string)
	} else {
		resp = false
	}
	owner = "Idea Market"
	if res["Ownership_price"] != nil {
		own_p = res["Ownership_price"].(float64)
	} else {
		resp = false
	}
	if res["Viewing_price"] != nil {
		view_p = res["Viewing_price"].(float64)
	} else {
		resp = false
	}
	if res["Problem"] != nil {
		problem = res["Problem"].(string)
	} else {
		resp = false
	}
	if res["Domain"] != nil {
		domain = res["Domain"].(string)
	} else {
		resp = false
	}

	if resp == true {
		json.NewEncoder(w).Encode("proposed successfully")
		nBlock := CreateBlockData(title, description, owner, problem, domain, t_u, view_p, own_p, own_p)

		fmt.Println(nBlock)

		gobEncoder := gob.NewEncoder(conn)
		err := gobEncoder.Encode(1)
		checkError(err)
		err = gobEncoder.Encode(nBlock)
		checkError(err)

	} else {
		json.NewEncoder(w).Encode("Record was not complete")
	}

}

// Used to configure the above methods according to the url specified by the front end
func http_server() {
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/listideas", ListIdeas)
	http.HandleFunc("/proposeidea", ProposeIdea)
	http.ListenAndServe(":8081", nil)
}
