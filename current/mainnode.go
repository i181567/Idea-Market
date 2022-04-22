package main

import (
	"crypto/sha256"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
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

type NewBlockData struct {
	Title                 string
	Description           string
	Owners                []string
	Problem               string
	Domain                string
	Technologies_used     []string
	Viewing_price         float64
	Ownership_price       float64
	Pricing_history       []float64
	Can_be_viewed_by      []string
	Bidding               bool
	Highest_bidder        string
	Highest_bidding_price float64
	Start_bidding_time    string
	End_bidding_time      string
	Hash_of_idea          string
}

// Structure of a Block
type NewBlock struct {
	Data        NewBlockData
	PrevPointer *NewBlock
	PrevHash    string
	CurrentHash string
}
type pending_idea struct {
	Title             string
	Description       string
	Owners            []string
	Problem           string
	Domain            string
	Technologies_used []string
	Viewing_price     float64
	Ownership_price   float64
	Pricing_history   []float64
	Score_text        string
	Score             float64
}

// Same structure as the above, but in json format
type ideas struct {
	Ideas []BlockData `json:"ideas"`
}

// Database_struct
type user struct {
	Username    string
	Password    string
	Balance     int
	Email       string
	Phonenumber string
}

type user_1 struct {
	Username    string
	Password    string
	Email       string
	Phonenumber string
}

type auth struct {
	Username string
	Password string
}
type bidding_param struct {
	Title        string
	Username     string
	Biddingprice float64
}

type bidding_param2 struct {
	Title    string
	Username string
	Password string
}

type deposit_param struct {
	Username string
	Password string
	Balance  float64
}

type user_json struct {
	User_info []user `json:"user_info"`
}

// Structure of a Block
type Block struct {
	Data        BlockData
	PrevPointer *Block
	PrevHash    string
	CurrentHash string
}
type title_username struct {
	Title    string
	Username string
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

// This method is used to print the whole Blockchain on console.
func ListNewBlocks(chainHead *NewBlock) {
	var Temp *NewBlock
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

func CreateNewBlockData(title string, desc string, owners string, prob string, dom string, tech []string, vp float64, op float64, ph float64) NewBlockData {
	var Idea NewBlockData
	Idea.Title = title
	Idea.Description = desc
	Idea.Owners = append(Idea.Owners, owners)
	Idea.Problem = prob
	Idea.Domain = dom
	Idea.Technologies_used = tech
	Idea.Viewing_price = vp
	Idea.Ownership_price = op
	Idea.Pricing_history = append(Idea.Pricing_history, ph)
	Idea.Bidding = false
	Idea.Can_be_viewed_by = append(Idea.Can_be_viewed_by, "")
	Idea.Highest_bidder = ""
	Idea.Highest_bidding_price = 0
	Idea.Start_bidding_time = ""
	Idea.End_bidding_time = ""
	Idea.Hash_of_idea = CalculateIdeaHash(Idea)

	return Idea
}

func CreateNewBlockData_from_blockdata(rhs BlockData) NewBlockData {
	var Idea NewBlockData
	Idea.Title = rhs.Title
	Idea.Description = rhs.Description
	Idea.Owners = rhs.Owners
	Idea.Problem = rhs.Problem
	Idea.Domain = rhs.Domain
	Idea.Technologies_used = rhs.Technologies_used
	Idea.Viewing_price = rhs.Viewing_price
	Idea.Ownership_price = rhs.Ownership_price
	Idea.Pricing_history = rhs.Pricing_history
	Idea.Bidding = false
	Idea.Can_be_viewed_by = append(Idea.Can_be_viewed_by, "")
	Idea.Highest_bidder = ""
	Idea.Highest_bidding_price = 0
	Idea.Start_bidding_time = ""
	Idea.End_bidding_time = ""
	Idea.Hash_of_idea = CalculateIdeaHash(Idea)

	return Idea
}

func CreateNewBlockData_from_pending_idea(rhs pending_idea) NewBlockData {
	var Idea NewBlockData
	Idea.Title = rhs.Title
	Idea.Description = rhs.Description
	Idea.Owners = rhs.Owners
	Idea.Problem = rhs.Problem
	Idea.Domain = rhs.Domain
	Idea.Technologies_used = rhs.Technologies_used
	Idea.Viewing_price = rhs.Viewing_price
	Idea.Ownership_price = rhs.Ownership_price
	Idea.Pricing_history = rhs.Pricing_history
	Idea.Bidding = false
	Idea.Can_be_viewed_by = append(Idea.Can_be_viewed_by, "")
	Idea.Highest_bidder = ""
	Idea.Highest_bidding_price = 0
	Idea.Start_bidding_time = ""
	Idea.End_bidding_time = ""
	Idea.Hash_of_idea = CalculateIdeaHash(Idea)

	return Idea
}

// This method is used to create a Blockdata
func CalculateHash(inputBlock *Block, phash string) string {
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

	hash := fmt.Sprintf("%x", sha256.Sum256([]byte(t+d+o+p+do+tu+vp+op+ph+phash)))
	return hash
}

func CalculatenewHash(inputBlock *NewBlock, phash string) string {
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

	hash := fmt.Sprintf("%x", sha256.Sum256([]byte(t+d+o+p+do+tu+vp+op+ph+phash)))
	return hash
}

func CalculateIdeaHash(inputBlock NewBlockData) string {
	var t, d, o, p, do, tu, vp, op, ph, canbeviewedby, bidding, start, end string
	t = fmt.Sprintf("%v", inputBlock.Title)
	d = fmt.Sprintf("%v", inputBlock.Description)
	o = fmt.Sprintf("%v", inputBlock.Owners)
	p = fmt.Sprintf("%v", inputBlock.Problem)
	do = fmt.Sprintf("%v", inputBlock.Domain)
	tu = fmt.Sprintf("%v", inputBlock.Technologies_used)
	vp = fmt.Sprintf("%v", inputBlock.Viewing_price)
	op = fmt.Sprintf("%v", inputBlock.Ownership_price)
	ph = fmt.Sprintf("%v", inputBlock.Pricing_history)

	canbeviewedby = fmt.Sprintf("%v", inputBlock.Can_be_viewed_by)
	bidding = fmt.Sprintf("%v", inputBlock.Bidding)
	start = fmt.Sprintf("%v", inputBlock.Start_bidding_time)
	end = fmt.Sprintf("%v", inputBlock.End_bidding_time)

	hash := fmt.Sprintf("%x", sha256.Sum256([]byte(t+d+o+p+do+tu+vp+op+ph+canbeviewedby+bidding+start+end)))
	return hash
}

// This method is used to insert a new block to the blockchain.
func InsertBlock(dataToInsert BlockData, chainHead *Block) *Block {
	NewBlock := Block{}
	if chainHead == nil {
		NewBlock.Data = dataToInsert
		NewBlock.PrevHash = "0"
		NewBlock.CurrentHash = CalculateHash(&NewBlock, NewBlock.PrevHash)
		NewBlock.PrevPointer = nil
	} else {
		NewBlock.Data = dataToInsert
		NewBlock.PrevHash = chainHead.CurrentHash
		NewBlock.CurrentHash = CalculateHash(&NewBlock, NewBlock.PrevHash)
		NewBlock.PrevPointer = chainHead
	}
	chainHead = &NewBlock
	return chainHead
}

func InsertnewBlock(dataToInsert NewBlockData, chainHead *NewBlock) *NewBlock {
	if dataToInsert.Title != "" && len(dataToInsert.Owners) > 0 {
		New_Block := NewBlock{}
		if chainHead == nil {
			New_Block.Data = dataToInsert
			New_Block.PrevHash = "0"
			New_Block.CurrentHash = CalculatenewHash(&New_Block, New_Block.PrevHash)
			New_Block.PrevPointer = nil
		} else {
			New_Block.Data = dataToInsert
			New_Block.PrevHash = chainHead.CurrentHash
			New_Block.CurrentHash = CalculatenewHash(&New_Block, New_Block.PrevHash)
			New_Block.PrevPointer = chainHead
		}
		chainHead = &New_Block
	} else {
		fmt.Println("No owner for IDea!")
	}
	return chainHead
}

// This method is used while creating a block by reading the json file
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

func CreateNewBlockData_J(title string, desc string, owners []string, prob string, dom string, tech []string, vp float64, op float64, ph []float64, can_be_viewed_by []string,
	bidding bool, highest_bidder string, highest_bidding_prcie float64, start_time string, end_time string, hash_of_idea string) NewBlockData {
	var Idea NewBlockData
	Idea.Title = title
	Idea.Description = desc
	Idea.Owners = owners
	Idea.Problem = prob
	Idea.Domain = dom
	Idea.Technologies_used = tech
	Idea.Viewing_price = vp
	Idea.Ownership_price = op
	Idea.Pricing_history = ph
	Idea.Can_be_viewed_by = can_be_viewed_by
	Idea.Bidding = bidding
	Idea.Highest_bidder = highest_bidder
	Idea.Highest_bidding_price = highest_bidding_prcie
	Idea.Start_bidding_time = start_time
	Idea.End_bidding_time = end_time
	Idea.Hash_of_idea = hash_of_idea
	return Idea
}

// This method is used to create a Blockdata
func CreatePendingIdea(title string, desc string, owners string, prob string, dom string, tech string, vp float64, op float64, ph float64, simidea string, simscore float64) pending_idea {
	var Idea pending_idea
	Idea.Title = title
	Idea.Description = desc
	Idea.Owners = append(Idea.Owners, owners)
	Idea.Problem = prob
	Idea.Domain = dom
	Idea.Technologies_used = append(Idea.Technologies_used, tech)
	Idea.Viewing_price = vp
	Idea.Ownership_price = op
	Idea.Pricing_history = append(Idea.Pricing_history, ph)
	Idea.Score_text = simidea
	Idea.Score = simscore
	return Idea
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

func Read_new_from_json(file_name string, newchain_head *NewBlock) *NewBlock {
	var ideas []NewBlockData
	jsonFile, err := os.Open(file_name)
	if err != nil {
		fmt.Println(err)
	}
	newchain_head = nil
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &ideas)
	for i := len(ideas) - 1; i >= 0; i-- {
		if len(ideas[i].Title) > 2 && len(ideas[i].Owners) > 0 {
			I1 := CreateNewBlockData_J(ideas[i].Title, ideas[i].Description, ideas[i].Owners, ideas[i].Problem, ideas[i].Description, ideas[i].Technologies_used, ideas[i].Viewing_price, ideas[i].Ownership_price, ideas[i].Pricing_history,
				ideas[i].Can_be_viewed_by, ideas[i].Bidding, ideas[i].Highest_bidder, ideas[i].Highest_bidding_price, ideas[i].Start_bidding_time, ideas[i].End_bidding_time, ideas[i].Hash_of_idea)
			newchain_head = InsertnewBlock(I1, newchain_head)
		}
	}

	//ListBlocks(chain_head)

	defer jsonFile.Close()

	// Title                 string
	// Description           string
	// Owners                []string
	// Problem               string
	// Domain                string
	// Technologies_used     []string
	// Viewing_price         float64
	// Ownership_price       float64
	// Pricing_history       []float64
	// Can_be_viewed_by      []string
	// Bidding               bool
	// Highest_bidder        string
	// Highest_bidding_price float32
	// Start_bidding_time    string
	// End_bidding_time      string
	// Hash_of_idea          string
	return newchain_head
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

func Write_New_to_json(file_name string, chain_head *NewBlock) {
	var data []NewBlockData
	temp := chain_head
	for temp != nil {

		data = append(data, temp.Data)
		temp = temp.PrevPointer

	}

	file, _ := json.MarshalIndent(data, "", " ")

	_ = ioutil.WriteFile(file_name, file, 0644)
}
func Read_pending_ideas_from_json(file_name string) {
	var ideas []pending_idea
	jsonFile, err := os.Open(file_name)
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &ideas)
	pending_ideas = nil
	for i := 0; i < len(ideas); i++ {
		pending_ideas = append(pending_ideas, ideas[i])
	}

	//ListBlocks(chain_head)

	defer jsonFile.Close()
}
func Write_pending_ideas_to_json(file_name string) {
	file, _ := json.MarshalIndent(pending_ideas, "", " ")
	_ = ioutil.WriteFile(file_name, file, 0644)
}

// Method to check error, if any
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

/* --------------------------------------------------------------- methods used by http server---------------------------------------------------------------*/
// This method enables Cors, by which the front end in angular can send and recieve data to it.
func enableCors(w *http.ResponseWriter, r *http.Request) {
	// (*w).Header().Set("Access-Control-Allow-Origin", "*")
	// allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"
	// if origin := r.Header.Get("Origin"); origin != "" {
	// 	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	// 	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	// 	(*w).Header().Set("Access-Control-Allow-Headers", allowedHeaders)
	// 	(*w).Header().Set("Access-Control-Expose-Headers", "Authorization")
	// }

	// (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	// (*w).Header().Set("Access-Control-Allow-Headers", allowedHeaders)
	// (*w).Header().Set("Access-Control-Expose-Headers", "Authorization")
	// (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Credentials", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "*")
}

// Just to check the server
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server is up and running %s!", r.URL.Path[1:])
}

func is_idea_present(l_ideas ideas, i1 NewBlockData) bool {
	for i := 0; i < len(l_ideas.Ideas); i++ {
		if l_ideas.Ideas[i].Title == i1.Title {
			return true
		}
	}
	return false
}

func is_idea_present_2(l_ideas []NewBlockData, i1 NewBlockData) bool {
	for i := 0; i < len(l_ideas); i++ {
		if l_ideas[i].Title == i1.Title {
			return true
		}
	}
	return false
}

// Used by front-end to recieve data.
func ListIdeas(w http.ResponseWriter, r *http.Request) {
	enableCors(&w, r)
	B1 := newchain_head
	var list_ideas ideas
	var bb1 BlockData

	for ; B1 != nil; B1 = B1.PrevPointer {
		if !(is_idea_present(list_ideas, B1.Data)) {
			// ideas is the same datastructure of an idea but in json format
			bb1.Title = B1.Data.Title
			bb1.Description = B1.Data.Description
			bb1.Owners = append(bb1.Owners, B1.Data.Owners...)
			bb1.Problem = B1.Data.Problem
			bb1.Domain = B1.Data.Domain
			bb1.Technologies_used = B1.Data.Technologies_used
			bb1.Viewing_price = B1.Data.Viewing_price
			bb1.Ownership_price = B1.Data.Ownership_price
			bb1.Pricing_history = append(bb1.Pricing_history, B1.Data.Pricing_history...)
			list_ideas.Ideas = append(list_ideas.Ideas, bb1)
		}
	}
	json.NewEncoder(w).Encode(list_ideas.Ideas)
}
func Ideas_in_auction(w http.ResponseWriter, r *http.Request) {
	enableCors(&w, r)
	var res map[string]interface{}
	json.NewDecoder(r.Body).Decode(&res)

	user := res["Username"]
	B1 := newchain_head
	var bb1 NewBlockData
	var all_ideas []NewBlockData
	var all_ideas_2 []NewBlockData

	for ; B1 != nil; B1 = B1.PrevPointer {
		if !(is_idea_present_2(all_ideas_2, B1.Data)) {
			if B1.Data.Bidding == true && len(B1.Data.Owners) > 0 {
				if user != B1.Data.Owners[len(B1.Data.Owners)-1] {
					bb1 = B1.Data
					Auth1 := 0
					for q := 0; q < len(B1.Data.Can_be_viewed_by); q++ {
						if B1.Data.Can_be_viewed_by[q] == user {
							Auth1 = 1
							break
						}
					}
					bb1.Problem = ""
					bb1.Technologies_used = bb1.Technologies_used[:0]
					bb1.Can_be_viewed_by = bb1.Can_be_viewed_by[:0]
					bb1.Highest_bidder = ""
					bb1.Can_be_viewed_by = bb1.Can_be_viewed_by[:0]
					if Auth1 == 0 {
						bb1.Description = ""

					}
					all_ideas = append(all_ideas, bb1)

				}
			}
		}
		all_ideas_2 = append(all_ideas_2, B1.Data)
	}
	json.NewEncoder(w).Encode(all_ideas)
}
func ProposedIdeas(w http.ResponseWriter, r *http.Request) {
	enableCors(&w, r)
	json.NewEncoder(w).Encode(pending_ideas)
}
func showideas(w http.ResponseWriter, r *http.Request) {
	enableCors(&w, r)
	var res map[string]interface{}
	json.NewDecoder(r.Body).Decode(&res)

	user := res["Username"]
	B1 := newchain_head
	var bb1 NewBlockData
	var all_ideas []NewBlockData

	for ; B1 != nil; B1 = B1.PrevPointer {
		if !(is_idea_present_2(all_ideas, B1.Data)) {
			// ideas is the same datastructure of an idea but in json format
			if len(B1.Data.Owners) > 0 {
				if user == B1.Data.Owners[len(B1.Data.Owners)-1] {
					all_ideas = append(all_ideas, B1.Data)
					// fmt.Println("Yess")
					// fmt.Println(B1.Data)
				} else {
					// fmt.Println("No")
					// fmt.Println(B1.Data)
					// fmt.Println(B1.Data.Owners[len(B1.Data.Owners)-1])
					// fmt.Println(res["Username"])
					bb1 = B1.Data
					Auth1 := 0
					for q := 0; q < len(B1.Data.Can_be_viewed_by); q++ {
						if B1.Data.Can_be_viewed_by[q] == user {
							Auth1 = 1
							break
						}
					}
					bb1.Problem = ""
					bb1.Technologies_used = bb1.Technologies_used[:0]
					//bb1.Can_be_viewed_by = bb1.Can_be_viewed_by[:0]
					bb1.Highest_bidder = ""
					//bb1.Can_be_viewed_by = bb1.Can_be_viewed_by[:0]
					if Auth1 == 0 {
						bb1.Description = ""

					}
					all_ideas = append(all_ideas, bb1)
				}
			}
		}
	}
	json.NewEncoder(w).Encode(all_ideas)

}
func myideas(w http.ResponseWriter, r *http.Request) {
	enableCors(&w, r)

	var res map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&res)
	if err != nil {
		fmt.Println("ERR")
		json.NewEncoder(w).Encode("ERR")
		return
	}

	user := res["Username"]
	B1 := newchain_head
	var all_ideas []NewBlockData

	for ; B1 != nil; B1 = B1.PrevPointer {
		if !(is_idea_present_2(all_ideas, B1.Data)) {
			// ideas is the same datastructure of an idea but in json format
			// fmt.Println(B1.Data.Title, "    ", B1.Data.Owners[len(B1.Data.Owners)-1])
			if len(B1.Data.Owners) <= 0 {
				fmt.Println("NO OWNER")
				json.NewEncoder(w).Encode("NO OWNER")
				return
			}
			if user == B1.Data.Owners[len(B1.Data.Owners)-1] {
				all_ideas = append(all_ideas, B1.Data)
			}
		}
	}
	json.NewEncoder(w).Encode(all_ideas)

}
func viewidea(w http.ResponseWriter, r *http.Request) {
	enableCors(&w, r)
	//{"Username,"Title"}

	// var res map[string]interface{}
	// json.NewDecoder(r.Body).Decode(&res)

	// fmt.Println(res)
	var ins1 title_username
	err := json.NewDecoder(r.Body).Decode(&ins1)
	if err != nil {
		fmt.Println("An Error occured!")
		return
	}
	var user_balance float64

	var Temp *NewBlock
	Temp = newchain_head
	var nb NewBlockData
	checker := 0
	fmt.Print("\n")
	for ; Temp != nil; Temp = Temp.PrevPointer {
		if Temp.Data.Title == ins1.Title {
			nb = Temp.Data
			checker = 1
			break
		}

	}
	if checker == 0 {
		fmt.Println("Invalid IDEA!")
		json.NewEncoder(w).Encode("Invalid IDEA!")
		return
	}

	// Check if the user can already view the idea or not
	for ww := 0; ww < len(nb.Can_be_viewed_by); ww++ {
		if nb.Can_be_viewed_by[ww] == ins1.Username {
			fmt.Println("You can already view the Idea's Description!")
			json.NewEncoder(w).Encode("You can already view the Idea's Description!")
			return
		}
	}
	// Check if the user is the owner of the idea or not
	if nb.Owners[len(nb.Owners)-1] != ins1.Username {

		db, err := sql.Open("mysql", mysql_str)
		if err != nil {
			panic(err.Error())
		}
		results, err := db.Query("select balance from user where username = '" + ins1.Username + "'")
		if err != nil {
			json.NewEncoder(w).Encode("ERR")
			json.NewEncoder(w).Encode("Invalid username!")
			fmt.Print("Invalid Username!")
			defer db.Close()
			return
		}
		for results.Next() {
			err = results.Scan(&user_balance)
			if err != nil {
				json.NewEncoder(w).Encode("An error occured!")
				fmt.Print("An error occured!")
				defer db.Close()
				break
			}
		}
		// checking the balance of user
		if user_balance > nb.Viewing_price {

			db, err := sql.Open("mysql", mysql_str)
			if err != nil {
				fmt.Println("An Errorr ouccured!")
			}

			// Giving cost to the owner
			results, err := db.Exec("update user set balance=balance+" + strconv.FormatFloat(nb.Viewing_price, 'E', -1, 32) + " where username ='" + nb.Owners[len(nb.Owners)-1] + "';")
			if err != nil {
				json.NewEncoder(w).Encode("Error")
				//panic(err.Error()) // proper error handling instead of panic in your app
				return
			}
			lastId, err := results.RowsAffected()
			if err != nil {
				fmt.Println("An Errorr ouccured!")
			}
			fmt.Printf("Affected rows: %d\n", lastId)

			// Deducting cost from user
			results, err = db.Exec("update user set balance=balance-" + strconv.FormatFloat(nb.Viewing_price, 'E', -1, 32) + " where username = '" + ins1.Username + "';")
			if err != nil {
				json.NewEncoder(w).Encode("Error")
				//panic(err.Error()) // proper error handling instead of panic in your app
				return
			}
			lastId, err = results.RowsAffected()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Affected rows: %d\n", lastId)

			nb.Can_be_viewed_by = append(nb.Can_be_viewed_by, ins1.Username)

			nb.Hash_of_idea = CalculateIdeaHash(nb)
			newchain_head = InsertnewBlock(nb, newchain_head)
			Write_New_to_json("newdata.json", newchain_head)

			json.NewEncoder(w).Encode("Operation completed successfully!")
			fmt.Println("Operation completed Successfully!")

		} else {
			json.NewEncoder(w).Encode("Not enough Balance!")
			fmt.Println("Not enough Balance!")
			defer db.Close()
		}
	} else {
		json.NewEncoder(w).Encode("You are already the owner!")
		fmt.Println("Owner cannot pay fee to view the idea!")
	}

}
func Updateuser(w http.ResponseWriter, r *http.Request) {
	enableCors(&w, r)

	var usr1 user_1
	json.NewDecoder(r.Body).Decode(&usr1)

	db, err := sql.Open("mysql", mysql_str)
	if err != nil {
		panic(err.Error())
	}

	results, err := db.Exec("update user set password='" + usr1.Password + "',email='" + usr1.Email + "',phone_number='" + usr1.Phonenumber + "'where username='" + usr1.Username + "';")
	if err != nil {
		json.NewEncoder(w).Encode("Error")
		//panic(err.Error()) // proper error handling instead of panic in your app
		return
	}
	lastId, err := results.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Affected rows: %d\n", lastId)
	json.NewEncoder(w).Encode("Account updated successfully!")

	defer db.Close()
}
func start_bidding(w http.ResponseWriter, r *http.Request) {
	enableCors(&w, r)

	var res bidding_param2
	json.NewDecoder(r.Body).Decode(&res)

	idea_title := res.Title

	var Temp *NewBlock
	Temp = newchain_head
	var nb NewBlockData
	fmt.Print("\n")
	for ; Temp != nil; Temp = Temp.PrevPointer {
		if Temp.Data.Title == idea_title {
			nb = Temp.Data
			break
		}
	}
	if nb.Bidding == false {
		nb.Bidding = true
		nb.Hash_of_idea = CalculateIdeaHash(nb)
		newchain_head = InsertnewBlock(nb, newchain_head)

		json.NewEncoder(w).Encode("Bidding started successfully!")
		fmt.Println("Bidding started successfully!")
		Write_New_to_json("newdata.json", newchain_head)
	} else {
		json.NewEncoder(w).Encode("Bidding is already started for you Idea!")
		fmt.Println("Bidding is already started for you Idea!")
	}

}
func stop_bidding(w http.ResponseWriter, r *http.Request) {
	enableCors(&w, r)

	var res bidding_param2
	json.NewDecoder(r.Body).Decode(&res)

	idea_title := res.Title

	var Temp *NewBlock
	Temp = newchain_head
	var nb NewBlockData
	fmt.Print("\n")
	ch := 0
	for ; Temp != nil; Temp = Temp.PrevPointer {
		if Temp.Data.Title == idea_title {
			nb = Temp.Data
			ch = 1
			break
		}
	}
	if ch == 0 {
		return
	}

	if nb.Bidding == true {
		if nb.Highest_bidding_price > 0.0 {

			db, err := sql.Open("mysql", mysql_str)
			if err != nil {
				panic(err.Error())
			}

			results, err := db.Exec("update user set balance=balance+" + strconv.FormatFloat(nb.Highest_bidding_price, 'E', -1, 32) + " where username ='" + nb.Owners[len(nb.Owners)-1] + "';")
			if err != nil {
				json.NewEncoder(w).Encode("Error")
				//panic(err.Error()) // proper error handling instead of panic in your app
				return
			}
			lastId, err := results.RowsAffected()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Affected rows: %d\n", lastId)

			results, err = db.Exec("update user set balance=0 where username = 'escrow';")
			if err != nil {
				json.NewEncoder(w).Encode("Error")
				//panic(err.Error()) // proper error handling instead of panic in your app
				return
			}
			lastId, err = results.RowsAffected()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Affected rows: %d\n", lastId)

			nb.Bidding = false
			nb.Owners = append(nb.Owners, nb.Highest_bidder)
			nb.Ownership_price = nb.Highest_bidding_price
			nb.Highest_bidder = ""
			nb.Highest_bidding_price = 0.0
			nb.Can_be_viewed_by = nb.Can_be_viewed_by[:0]
			nb.Bidding = false
			nb.Pricing_history = append(nb.Pricing_history, nb.Ownership_price)
			nb.Hash_of_idea = CalculateIdeaHash(nb)

			newchain_head = InsertnewBlock(nb, newchain_head)
			Write_New_to_json("newdata.json", newchain_head)

			json.NewEncoder(w).Encode("Bidding Stopped successfully and the ownership of Idea is transferred!")
			fmt.Println("Bidding Stopped successfully and the ownership of Idea is transferred!")
			defer db.Close()
		} else {
			nb.Bidding = false
			nb.Start_bidding_time = ""
			nb.End_bidding_time = ""
			nb.Hash_of_idea = CalculateIdeaHash(nb)
			newchain_head = InsertnewBlock(nb, newchain_head)
			Write_New_to_json("newdata.json", newchain_head)
			json.NewEncoder(w).Encode("Bidding Stopped successfully but the ownership is not changed!")
			fmt.Println("Bidding Stopped successfully but the ownership is not changed!")
		}
	} else {

		json.NewEncoder(w).Encode("Bidding is already stopped!")
		fmt.Println("Bidding is already stopped!")
	}
}
func bid(w http.ResponseWriter, r *http.Request) {

	// Title        string
	// Username     string
	// Biddingprice float64
	enableCors(&w, r)
	var bparam bidding_param
	json.NewDecoder(r.Body).Decode(&bparam)

	// user (bidder's) balance in his/her account
	var user_balance float64
	// Block of the Idea of interest
	var BlockData_of_interest NewBlockData

	// Finding the block of interest
	var Temp *NewBlock
	Temp = newchain_head
	fmt.Print("\n")
	for ; Temp != nil; Temp = Temp.PrevPointer {
		if Temp.Data.Title == bparam.Title {
			BlockData_of_interest = Temp.Data
			break
		}
	}
	if BlockData_of_interest.Bidding == true {

		db, err := sql.Open("mysql", mysql_str)
		if err != nil {
			panic(err.Error())
		}
		results, err := db.Query("select balance from user where username = '" + bparam.Username + "'")
		if err != nil {
			json.NewEncoder(w).Encode("ERR")
			json.NewEncoder(w).Encode("Invalid username!")
			fmt.Print("Invalid Username!")
			defer db.Close()
			return
		}
		for results.Next() {
			err = results.Scan(&user_balance)
			if err != nil {
				json.NewEncoder(w).Encode("An error occured!")
				fmt.Print("An error occured!")
				defer db.Close()
				break
			}
		}
		if user_balance > bparam.Biddingprice {

			if bparam.Biddingprice >= BlockData_of_interest.Ownership_price {
				if bparam.Biddingprice >= BlockData_of_interest.Highest_bidding_price {

					// Deduct amount from the current bidder.
					fmt.Println(strconv.FormatFloat(bparam.Biddingprice, 'E', -1, 32))
					results2, err2 := db.Exec("update user set balance = balance-" + strconv.FormatFloat(bparam.Biddingprice, 'E', -1, 32) + " where username = '" + bparam.Username + "';")
					if err2 != nil {
						json.NewEncoder(w).Encode("Error")
						fmt.Println("Error occured!")
						return
					}
					lastId, err := results2.RowsAffected()
					if err != nil {
						log.Fatal(err)
					}
					fmt.Printf("Affected rows: %d\n", lastId)

					// Update Amount in Escrow.
					results2, err2 = db.Exec("update user set balance = " + strconv.FormatFloat(bparam.Biddingprice, 'E', -1, 32) + " where username = 'escrow';")
					if err2 != nil {
						json.NewEncoder(w).Encode("Error")
						fmt.Println("Error occured!")
						return
					}
					lastId, err = results2.RowsAffected()
					if err != nil {
						log.Fatal(err)
					}
					fmt.Printf("Affected rows: %d\n", lastId)

					// Return the balance to the previous bidder
					if BlockData_of_interest.Highest_bidding_price > 0.0 {
						results2, err2 = db.Exec("update user set balance = balance +" + strconv.FormatFloat(BlockData_of_interest.Highest_bidding_price, 'E', -1, 32) + " where username = '" + BlockData_of_interest.Highest_bidder + "';")
						if err2 != nil {
							json.NewEncoder(w).Encode("Error")
							fmt.Println("Error!")
							return
						}
						lastId, err = results2.RowsAffected()
						if err != nil {
							log.Fatal(err)
						}
						fmt.Printf("Affected rows: %d\n", lastId)
					}

					BlockData_of_interest.Highest_bidder = bparam.Username
					BlockData_of_interest.Highest_bidding_price = bparam.Biddingprice
					BlockData_of_interest.Hash_of_idea = CalculateIdeaHash(BlockData_of_interest)
					newchain_head = InsertnewBlock(BlockData_of_interest, newchain_head)
					Write_New_to_json("newdata.json", newchain_head)
					json.NewEncoder(w).Encode("Bided successfully!")
					fmt.Println("Bidded Successfully!")

				} else {
					json.NewEncoder(w).Encode("A Bid of more price is already there!")
					fmt.Print("A Bid of more price is already there!")
				}

			} else {
				json.NewEncoder(w).Encode("Bidding price should be more than the ownership price of an idea!")
				fmt.Print("Bidding price should be more than the ownership price of an idea!")
			}

		} else {
			json.NewEncoder(w).Encode("Not enough balance!")
			fmt.Print("Not enough balance!")
		}
		defer db.Close()
	} else {
		json.NewEncoder(w).Encode("Idea is not open for bidding!")
		fmt.Print("Idea is not open for bidding!")
	}
}
func Signup(w http.ResponseWriter, r *http.Request) {
	enableCors(&w, r)

	//var res map[string]interface{}
	var usr1 user_1
	json.NewDecoder(r.Body).Decode(&usr1)
	if len(usr1.Username) > 3 {

		db, err := sql.Open("mysql", mysql_str)
		if err != nil {
			fmt.Println("Error while connecting to Database!")
			json.NewEncoder(w).Encode("Error while connecting to Database!")
			return
		}
		fmt.Println(r.Body)
		fmt.Println(usr1)

		results, err := db.Exec("insert into user(username,password,balance,email,phone_number) values ('" + usr1.Username + "','" + usr1.Password + "',0,'" + usr1.Email + "','" + usr1.Phonenumber + "');")
		if err != nil {
			json.NewEncoder(w).Encode("ERR")
			//panic(err.Error()) // proper error handling instead of panic in your app
			fmt.Print("Invalid Username!")
			defer db.Close()
			return
		}
		lastId, err := results.RowsAffected()
		if err != nil {
			json.NewEncoder(w).Encode("ERR")
			//panic(err.Error()) // proper error handling instead of panic in your app
			fmt.Print("ERR2!")
		}
		fmt.Printf("Affected rows: %d\n", lastId)
		json.NewEncoder(w).Encode("OK")

		defer db.Close()
	} else {
		fmt.Printf("Lenght of username is not enough")
		json.NewEncoder(w).Encode("Lenght of username is not enough")
	}
}

func Signin(w http.ResponseWriter, r *http.Request) {
	enableCors(&w, r)
	var usr1 user
	var aut1 auth
	json.NewDecoder(r.Body).Decode(&aut1)

	db, err := sql.Open("mysql", mysql_str)
	if err != nil {
		json.NewEncoder(w).Encode("Error while connecting to Database")
		fmt.Print("Invalid Username!")
		return
	}
	results, err := db.Query("select * from user where username = '" + aut1.Username + "'")
	if err != nil {
		json.NewEncoder(w).Encode("ERR")
		fmt.Print("Invalid Username!")
		return
	}
	authenticated := 0
	var usr2 user_json
	for results.Next() {
		err = results.Scan(&usr1.Username, &usr1.Password, &usr1.Balance, &usr1.Email, &usr1.Phonenumber)
		if err != nil {
			fmt.Print("An error occured!")
			json.NewEncoder(w).Encode("ERR")
			break
		}
		if usr1.Password == aut1.Password {
			usr2.User_info = append(usr2.User_info, usr1)
			// fmt.Println(usr1)
			// fmt.Println(usr2)
			// fmt.Println("\n\n")
			// fmt.Println(usr2.User_info)
			authenticated = 1
			break
		} else {
			fmt.Print("Invalid Password!")
			json.NewEncoder(w).Encode("ERR")
			break
		}
	}
	if authenticated == 0 {
		fmt.Print("Invalid Username or Password!")
		json.NewEncoder(w).Encode("ERR")
	} else {
		fmt.Println(usr2.User_info)
		fmt.Println("authenticated!")
		//json.NewEncoder(w).Encode("Authenticated!")
		json.NewEncoder(w).Encode(usr2.User_info)
	}

	defer db.Close()
}
func totalusers(w http.ResponseWriter, r *http.Request) {
	enableCors(&w, r)
	db, err := sql.Open("mysql", mysql_str)
	if err != nil {
		panic(err.Error())
	}
	results, err := db.Query("select count(*) from user;")
	if err != nil {
		json.NewEncoder(w).Encode("AN ERROR OCCURRED!")
		fmt.Print("AN ERROR OCCURRED!")
		defer db.Close()
		return
	}
	var number_of_users float64
	for results.Next() {
		err = results.Scan(&number_of_users)
		if err != nil {
			json.NewEncoder(w).Encode("An error occured!")
			fmt.Print("An error occured!")
			defer db.Close()
			break
		}
	}
	json.NewEncoder(w).Encode(&number_of_users)
	fmt.Print("Total number of users: ", number_of_users)
	defer db.Close()

}

func totalideasinauction(w http.ResponseWriter, r *http.Request) {
	fmt.Println("-----------Total IDEAS IN AUCTION-----------------------")
	enableCors(&w, r)
	number_of_ideas_in_the_auction = 0
	B1 := newchain_head
	var all_ideas []NewBlockData
	for ; B1 != nil; B1 = B1.PrevPointer {
		if !is_idea_present_2(all_ideas, B1.Data) {
			// ideas is the same datastructure of an idea but in json format
			if B1.Data.Bidding == true {
				number_of_ideas_in_the_auction += 1
			}
		}
		all_ideas = append(all_ideas, B1.Data)
	}
	json.NewEncoder(w).Encode(number_of_ideas_in_the_auction)
	fmt.Println("Total number of Ideas in Auction: ", number_of_ideas_in_the_auction)
}
func totalideas(w http.ResponseWriter, r *http.Request) {
	enableCors(&w, r)
	fmt.Println("-----------Total IDEAS-----------------------")
	total_number_of_ideas = 0
	B1 := newchain_head
	var all_ideas []NewBlockData
	for ; B1 != nil; B1 = B1.PrevPointer {
		if !(is_idea_present_2(all_ideas, B1.Data)) {
			fmt.Println(B1.Data)
			// ideas is the same datastructure of an idea but in json format
			total_number_of_ideas += 1
		}
		all_ideas = append(all_ideas, B1.Data)
	}
	json.NewEncoder(w).Encode(total_number_of_ideas)
	fmt.Println("Total number of Ideas in Auction: ", total_number_of_ideas)
}
func totalnumberofblocks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("-----------Total NUMBER OF BLOCKS-----------------------")
	enableCors(&w, r)
	total_number_of_blocks = 0
	B1 := newchain_head
	for ; B1 != nil; B1 = B1.PrevPointer {
		fmt.Println(B1.Data)
		total_number_of_blocks += 1
	}
	json.NewEncoder(w).Encode(total_number_of_blocks)
	fmt.Println("Total number of Blocks: ", total_number_of_blocks)
}
func Deleteuser(w http.ResponseWriter, r *http.Request) {
	enableCors(&w, r)
	var auth1 auth
	json.NewDecoder(r.Body).Decode(&auth1)
	B1 := newchain_head
	var all_ideas []NewBlockData

	for ; B1 != nil; B1 = B1.PrevPointer {
		if !(is_idea_present_2(all_ideas, B1.Data)) {
			// ideas is the same datastructure of an idea but in json format
			// fmt.Println(B1.Data.Title, "    ", B1.Data.Owners[len(B1.Data.Owners)-1])
			if auth1.Username == B1.Data.Owners[len(B1.Data.Owners)-1] {
				json.NewEncoder(w).Encode("Username cannot be deleted, as you own Ideas in the Blockchain!")
				fmt.Println("Username cannot be deleted, as you own Ideas in the Blockchain!")
				return
			}
		}
	}

	db, err := sql.Open("mysql", mysql_str)
	if err != nil {
		panic(err.Error())
	}
	results, err := db.Exec("delete from user where username ='" + auth1.Username + "' and password='" + auth1.Password + "'")
	if err != nil {
		json.NewEncoder(w).Encode("Error")
		defer db.Close()
		return
	}
	lastId, err := results.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Affected rows: %d\n", lastId)
	json.NewEncoder(w).Encode("Account deleted successfully!")

	defer db.Close()
}
func deposit_balance(w http.ResponseWriter, r *http.Request) {
	enableCors(&w, r)
	var dp1 deposit_param
	json.NewDecoder(r.Body).Decode(&dp1)

	db, err := sql.Open("mysql", mysql_str)
	if err != nil {
		fmt.Println("Error while connecting to Database!")
		json.NewEncoder(w).Encode("Error while connecting to Database!")
		return
	}
	strconv.FormatFloat(dp1.Balance, 'E', -1, 32)
	results, err := db.Exec("update user set balance = balance+" + strconv.FormatFloat(dp1.Balance, 'E', -1, 32) + " where username ='" + dp1.Username + "' and password ='" + dp1.Password + "';")

	if err != nil {
		json.NewEncoder(w).Encode("ERR")
		fmt.Println("ERR")
		defer db.Close()
		return
	}
	lastId, err := results.RowsAffected()
	if err != nil {
		json.NewEncoder(w).Encode("ERR")
		fmt.Println("ERR")
		defer db.Close()
		return
	}
	fmt.Printf("Affected rows: %d\n", lastId)
	fmt.Printf("Balance updated successfully!")
	json.NewEncoder(w).Encode("Balance updated successfully!")

	defer db.Close()
}

// Propose Idea.
func ProposeIdea(w http.ResponseWriter, r *http.Request) {
	enableCors(&w, r)
	var p_idea pending_idea
	err := json.NewDecoder(r.Body).Decode(&p_idea)
	if err == nil {
		fmt.Println(p_idea)
		fmt.Println(p_idea.Title)
		fmt.Println(p_idea.Description)
		pending_ideas = append(pending_ideas, p_idea)
		fmt.Println(pending_ideas)
		Write_pending_ideas_to_json("pending_ideas.json")
		json.NewEncoder(w).Encode("proposed successfully!")
	} else {
		json.NewEncoder(w).Encode("An Error occurred!")
		fmt.Println("An Error occurred in propose idea!")
	}
}

// Add Approved idea.
func AddIdea(w http.ResponseWriter, r *http.Request) {
	enableCors(&w, r)
	var p_idea pending_idea
	err := json.NewDecoder(r.Body).Decode(&p_idea)
	if err == nil {
		fmt.Println(p_idea)
		fmt.Println(p_idea.Title)
		fmt.Println(p_idea.Description)
		n_bdata := CreateNewBlockData_from_pending_idea(p_idea)
		newchain_head = InsertnewBlock(n_bdata, newchain_head)

		Write_New_to_json("newdata.json", newchain_head)

		for x := 0; x < len(pending_ideas); x++ {
			if pending_ideas[x].Title == p_idea.Title {
				pending_ideas[x] = pending_ideas[len(pending_ideas)-1]
				pending_ideas = pending_ideas[:len(pending_ideas)-1]
				break
			}
		}
		ListNewBlocks(newchain_head)
		fmt.Println(pending_ideas)
		Write_pending_ideas_to_json("pending_ideas.json")

		json.NewEncoder(w).Encode("Added to the Blockchain successfully!")
	} else {
		json.NewEncoder(w).Encode("An Error occurred!")
		fmt.Println("An Error occurred in Add idea!")
	}
}

func disapproveidea(w http.ResponseWriter, r *http.Request) {
	enableCors(&w, r)
	var p_idea pending_idea
	err := json.NewDecoder(r.Body).Decode(&p_idea)
	if err == nil {
		fmt.Println(p_idea)

		for x := 0; x < len(pending_ideas); x++ {
			if pending_ideas[x].Title == p_idea.Title {
				pending_ideas[x] = pending_ideas[len(pending_ideas)-1]
				pending_ideas = pending_ideas[:len(pending_ideas)-1]
				break
			}
		}
		fmt.Println(pending_ideas)
		Write_pending_ideas_to_json("pending_ideas.json")

		json.NewEncoder(w).Encode("Idea Rejected successfully!")
	} else {
		json.NewEncoder(w).Encode("ERR")
		fmt.Println("An Error occurred in Disapprove idea!")
	}
}

func insert_to_db() {
	p := 0
	db, err := sql.Open("mysql", mysql_str)
	if err != nil {
		panic(err.Error())
	}
	results, err := db.Exec("insert into user(username,password,balance,email,phone_number) values ('u8','p8',10,'u8@gmail.com','12345')")
	if err != nil {
		p++
		fmt.Println(p)
		panic(err.Error()) // proper error handling instead of panic in your app

	}

	lastId, err := results.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Affected rows: %d\n", lastId)

	defer db.Close()
}
func delete_user_from_db() {
	p := 0
	db, err := sql.Open("mysql", mysql_str)
	if err != nil {
		panic(err.Error())
	}
	results, err := db.Exec("delete from user where username = 'u8'")
	if err != nil {
		p++
		fmt.Println(p)
		panic(err.Error()) // proper error handling instead of panic in your app

	}

	lastId, err := results.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Affected rows: %d\n", lastId)

	defer db.Close()
}

// --------Global variables
var newchain_head *NewBlock
var chain_head *Block
var pending_ideas []pending_idea
var number_of_ideas_in_the_auction int
var total_number_of_ideas int
var total_number_of_blocks int

// Controls the whole program
var status_of_node = 1
var err error
var choice int
var mysql_str = "root:@tcp(127.0.0.1:3306)/ideamarket"

/* -------------------------------------------------------------------------------- main method --------------------------------------------------------------------------------*/
func main() {
	chain_head = Read_from_json("data.json", chain_head)
	Read_pending_ideas_from_json("pending_ideas.json")

	newchain_head = Read_new_from_json("newdata.json", newchain_head)
	// Temp := chain_head
	// fmt.Print("\n")
	// for ; Temp != nil; Temp = Temp.PrevPointer {
	// 	d1 := CreateNewBlockData2(Temp.Data)
	// 	newchain_head = InsertnewBlock(d1, newchain_head)
	// }

	ListNewBlocks(newchain_head)
	// Write_New_to_json("newdata.json", newchain_head)
	fmt.Print("\n\n----------------------------------------------------------------------------------------------------------------------\n\n")
	fmt.Print("\n\n                                          Centralized Blockchain is active\n\n")
	fmt.Print("\n\n----------------------------------------------------------------------------------------------------------------------\n\n")
	choice = 1
	// delete_user_from_db()
	// Runs an http server on another thread
	go http_server()

	for choice <= 1 {
		fmt.Print("------------------------------------------------------------------------------------------\n")
		fmt.Print("Press any number, greater than 1 to exit : \n")
		fmt.Print("------------------------------------------------------------------------------------------\n")
		fmt.Scan(&choice)
		if choice > 1 {
			break
		}
	}

	fmt.Println("Connection closed.........................................")
	fmt.Println("Server is Closing................................")
}

// Used to configure the above methods according to the url specified by the front end
func http_server() {
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/listideas", ListIdeas)
	http.HandleFunc("/proposeidea", ProposeIdea)
	http.HandleFunc("/signup", Signup)
	http.HandleFunc("/signin", Signin)
	http.HandleFunc("/deleteuser", Deleteuser)
	http.HandleFunc("/updateuser", Updateuser)
	http.HandleFunc("/proposedideas", ProposedIdeas)
	http.HandleFunc("/addidea", AddIdea)
	http.HandleFunc("/startbidding", start_bidding)
	http.HandleFunc("/stopbidding", stop_bidding)
	http.HandleFunc("/bid", bid)
	//
	http.HandleFunc("/showideas", showideas)
	http.HandleFunc("/viewidea", viewidea)
	http.HandleFunc("/myideas", myideas)

	http.HandleFunc("/totalusers", totalusers)
	http.HandleFunc("/totalideasinauction", totalideasinauction)
	http.HandleFunc("/totalblocks", totalnumberofblocks)
	http.HandleFunc("/totalideas", totalideas)

	http.HandleFunc("/disapproveidea", disapproveidea)

	http.HandleFunc("/deposit", deposit_balance)
	http.HandleFunc("/ideasinauction", Ideas_in_auction)
	//createuser

	http.ListenAndServe(":8081", nil)
}

// http://localhost:8081/getPendingIdeas
// http://localhost:8081/approveIdea
// http://localhost:8081/disapproveIdea
// https://github.com/abubakar2000/IdeaMarket-Backend

// func ProposeIdea(w http.ResponseWriter, r *http.Request) {
// 	enableCors(&w, r)
// 	//json.NewEncoder(w).Encode("proposed successfully")
// 	fmt.Println("Request incoming!!!!")
// 	var res map[string]interface{}

// 	var title, description, owner, problem, domain string
// 	var view_p, own_p float64
// 	var t_u []string
// 	t_u = append(t_u, "MEAN STACK")
// 	t_u = append(t_u, "MERN STACK")
// 	json.NewDecoder(r.Body).Decode(&res)
// 	fmt.Println(res)
// 	resp := true

// 	fmt.Println(res)

// 	if res["Title"] != nil {
// 		title = res["Title"].(string)
// 	} else {
// 		resp = false
// 		fmt.Println("title")
// 	}

// 	if res["Description"] != nil {
// 		description = res["Description"].(string)
// 	} else {
// 		resp = false
// 		fmt.Println("description")
// 	}
// 	owner = "Idea Market"
// 	if res["Ownership_price"] != nil {
// 		own_p = res["Ownership_price"].(float64)
// 	} else {
// 		resp = false
// 		fmt.Println("ownership price")
// 	}
// 	if res["Viewing_price"] != nil {
// 		view_p = res["Viewing_price"].(float64)
// 	} else {
// 		resp = false
// 		fmt.Println("viewing price")
// 	}
// 	if res["Problem"] != nil {
// 		problem = res["Problem"].(string)
// 	} else {
// 		resp = false
// 		fmt.Println("problem")
// 	}
// 	if res["Domain"] != nil {
// 		domain = res["Domain"].(string)
// 	} else {
// 		resp = false
// 		fmt.Println("doamin")
// 	}

// 	if resp == true {
// 		json.NewEncoder(w).Encode("proposed successfully")
// 		nBlock := CreateBlockData(title, description, owner, problem, domain, t_u, view_p, own_p, own_p)

// 		fmt.Println(nBlock)

// 		gobEncoder := gob.NewEncoder(conn)
// 		err := gobEncoder.Encode(1)
// 		checkError(err)
// 		err = gobEncoder.Encode(nBlock)
// 		checkError(err)

// 	} else {
// 		json.NewEncoder(w).Encode("Record was not complete")
// 	}

// }

//##########################################################
// Add approved idea to blockchain
// Change Blockdata by adding {start_time, end_time, bidding_on or off, can_be_viewed_by, bidder, bidding_price}
// Start bidding
// End bidding
// Check which bidding have ended

// Pending Idea Example
// {
// 	"Title": "T3",
// 	"Description": "d3",
// 	"Owners": [
// 	 "o3"
// 	],
// 	"Problem": "p",
// 	"Domain": "d3",
// 	"Technologies_used": [
// 	 "MEAN STACK",
// 	 "MERN STACK"
// 	],
// 	"Viewing_price": 34,
// 	"Ownership_price": 10,
// 	"Pricing_history": [
// 	 10
// 	],
// 	"Score_text": "t3",
// 	"Score": 21
//    }

// Total users
// Total Ideas in auction.
// Total value of blockchain.
