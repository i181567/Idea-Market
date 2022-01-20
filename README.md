# Idea-Market
An application that makes the use of custom Blockchain to protect the ownership of ideas.
# Pre-installed
To run the application, you need golang and node installed on your machine.
# Commands to run the application
First of all, you need to run the cnode.go by using <br />
go run cnode.go <br/>
Now the central node is running. After that, you can run as many nodes as you want by using the followiing command for each node.<br/>
go run node.go<br />
Now once the Blockchain is built and the api is ready, you can run the front end for each node by first going to the directory of the front end and then writing<br/>
npm install<br/>ng serve<br/>
Now open port 4200 and you can see the interface.
