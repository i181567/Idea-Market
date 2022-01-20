package main

import (
	"crypto/sha256"
	"database/sql"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

/* ---------------------------------------------------------------structure and methods used for building a Blockchain ---------------------------------------------------------------*/
// structure and methods used for building a Blockchain //
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
type ideas struct {
	Ideas []BlockData `json:"ideas"`
}
type Block struct {
	Data        BlockData
	PrevPointer *Block
	PrevHash    string
	CurrentHash string
}

// This method is used to create a Blockdata
func CalculateHash(inputBlock *Block) string {
	var t, d, o, p, do, tu, vp, op, ph string
	t = fmt.Sprintf("%v", inputBlock.Data.Title)
	d = fmt.Sprintf("%v", inputBlock.Data.Description)
	o = fmt.Sprintf("%v", inputBlock.Data.Owners)
	p = fmt.Sprintf("%v", inputBlock.Data.Problem)
	do = fmt.Sprintf("%v", inputBlock.Data.Domain)
	tu = fmt.Sprintf("%v", inputBlock.Data.Technologies_used)
	vp = fmt.Sprintf("%v", inputBlock.Data.Viewing_price)
	op = fmt.Sprintf("%v", inputBlock.Data.Ownership_price)
	ph = fmt.Sprintf("%v", inputBlock.Data.Pricing_history)

	hash := fmt.Sprintf("%x", sha256.Sum256([]byte(t+d+o+p+do+tu+vp+op+ph)))
	return hash
}

// This method is used to insert a new block to the blockchain.
func InsertBlock(dataToInsert BlockData, chainHead *Block) *Block {
	NewBlock := Block{}
	if chainHead == nil {
		NewBlock.Data = dataToInsert
		NewBlock.CurrentHash = CalculateHash(&NewBlock)
		NewBlock.PrevHash = "0"
		NewBlock.PrevPointer = nil
	} else {
		NewBlock.Data = dataToInsert
		NewBlock.CurrentHash = CalculateHash(&NewBlock)
		NewBlock.PrevHash = CalculateHash(chainHead)
		NewBlock.PrevPointer = chainHead
	}
	chainHead = &NewBlock
	return chainHead
}

func ListBlocks(chainHead *Block) {
	var Temp *Block
	Temp = chainHead
	fmt.Print("\n")
	for ; Temp != nil; Temp = Temp.PrevPointer {
		fmt.Print(Temp.Data)
		fmt.Print("\n")
	}
}

// This method is used to create a block
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

// Another method to create a block
func CreateBlockData_J(title string, desc string, owners []string, prob string, dom string, tech []string, vp float64, op float64, ph []float64) BlockData {
	var Idea BlockData
	Idea.Title = title
	Idea.Description = desc
	Idea.Owners = owners
	Idea.Problem = prob
	Idea.Domain = dom
	Idea.Technologies_used = tech
	Idea.Viewing_price = vp
	Idea.Ownership_price = op
	Idea.Pricing_history = ph
	return Idea
}

//-----------------------------------------------------------------Some global variables--------------------------------------------------------------------------------------//
var chain_head *Block    // top block of the main Blockchain
var all_nodes []net.Conn // Contains all nodes connected

// --------------------------------------------------------------------------------------------------------------------------------------------------------------------------- //
// When blokchain is updated, this node sends a copy to all the nodes connected
func send_updated_blockchain(connection net.Conn) {
	for i := 0; i < len(all_nodes); i++ {
		gobEncoder := gob.NewEncoder(all_nodes[i])
		err := gobEncoder.Encode(chain_head)
		checkError(err)
	}
}

// This function is called through another thread when a new node is connected
func handleConnection(connection net.Conn, problems []string) {
	gobEncoder := gob.NewEncoder(connection)
	err := gobEncoder.Encode(chain_head)
	checkError(err)
	fmt.Println("A New Node has connected : ", connection.RemoteAddr())
	all_nodes = append(all_nodes, connection)
	dec := gob.NewDecoder(connection)

	stopper := 3
	var NBlockData BlockData
	for stopper > 0 {
		send_updated_blockchain(connection)

		err := dec.Decode(&stopper)
		if err != nil {
			fmt.Println(err)
		}

		if stopper == 1 {
			err := dec.Decode(&NBlockData)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Node ", connection.RemoteAddr(), " wants to add the following idea, \npress 1, if the idea is valid,\npress anyother button to reject the idea\n:")
			fmt.Println(NBlockData)
			var idea_option int
			fmt.Scan(&idea_option)
			if idea_option == 1 {
				chain_head = InsertBlock(NBlockData, chain_head)
				Write_to_json("data.json", chain_head)

			} else {
			}

		} else {
			for x := 0; x < len(all_nodes); x++ {
				if connection == all_nodes[x] {
					all_nodes[x] = all_nodes[len(all_nodes)-1]
					all_nodes[len(all_nodes)-1] = nil
					all_nodes = all_nodes[:len(all_nodes)-1]
					stopper = 0
					break
				}
			}

			fmt.Println(all_nodes)

			connection.Close()
			break
		}
	}

}

// This method reads ideas from json file and creates a Blockchain
func Read_from_json(file_name string, chain_head *Block) *Block {
	var ideas ideas
	jsonFile, err := os.Open(file_name)
	if err != nil {
		fmt.Println(err)
	}
	chain_head = nil

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &ideas)
	for i := 0; i < len(ideas.Ideas); i++ {
		I1 := CreateBlockData_J(ideas.Ideas[i].Title, ideas.Ideas[i].Description, ideas.Ideas[i].Owners, ideas.Ideas[i].Problem, ideas.Ideas[i].Description, ideas.Ideas[i].Technologies_used, ideas.Ideas[i].Viewing_price, ideas.Ideas[i].Ownership_price, ideas.Ideas[i].Pricing_history)
		chain_head = InsertBlock(I1, chain_head)
	}

	//ListBlocks(chain_head)

	defer jsonFile.Close()

	return chain_head
}

// This method writes all ideas from blockchain to the json file
func Write_to_json(file_name string, chain_head *Block) {
	var data ideas
	temp := chain_head
	for temp != nil {

		data.Ideas = append(data.Ideas, temp.Data)
		temp = temp.PrevPointer

	}

	file, _ := json.MarshalIndent(data, "", " ")

	_ = ioutil.WriteFile(file_name, file, 0644)
}

// main method
func main() {
	test_db()

	var unsolved_problems []string
	unsolved_problems = append(unsolved_problems, "Problem1")
	unsolved_problems = append(unsolved_problems, "Problem2")

	fmt.Println("Reading data------------")
	chain_head = Read_from_json("data.json", chain_head)

	fmt.Print("\n\n----------------------------------------------------------------------------------------------------------------------\n\n")
	fmt.Print("\n\n                                          Centralized Blockchain is active\n\n")
	fmt.Print("\n\n----------------------------------------------------------------------------------------------------------------------\n\n")
	//ListBlocks(chain_head)
	ln, err := net.Listen("tcp", ":6000")
	checkError(err)
	fmt.Println(all_nodes)
	for {
		conn, err := ln.Accept()
		ListBlocks(chain_head)
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn, unsolved_problems)
	}

}

// checks error if any
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// sends blockchain to all nodes
func send_updated_blockchain_to_all_nodes(conn []net.Conn, chain_head *Block) {
	for i := 0; i < len(conn); i++ {
		gobEncoder := gob.NewEncoder(conn[i])
		err := gobEncoder.Encode(chain_head)
		checkError(err)
	}
}

// Database_methods
type user struct {
	username    string
	password    string
	balance     int
	email       string
	phonenumber string
}

// this method is used to create connection with mysql database
func test_db() {
	db, err := sql.Open("mysql", "root:wasif@tcp(127.0.0.1:3306)/ideamarket")
	if err != nil {
		panic(err.Error())
	}
	results, err := db.Query("SELECT * from user")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var tmp_user user
		// for each row, scan the result into our tag composite object
		err = results.Scan(&tmp_user.username, &tmp_user.password, &tmp_user.balance, &tmp_user.email, &tmp_user.phonenumber)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		fmt.Println(tmp_user)
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()
}
