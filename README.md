# Idea-Market
An application that makes the use of custom Blockchain to protect the ownership of ideas.

# Overview
This application consists of a front-end, that is developed using Angular and a Back-end that is developed using Golang. The database connected to the backend is Mysql. The backend has 2 files named cnode.go and node.go. cnode.go is a central node and can be only one. On the otherhand node.go can be more than one. When the application starts, the cnode.go creates a blockchain of ideas and waits for a node for connection. when a node(node.go) connects with it, it sends the blockchain to the node and then the node shows the ideas and information of the blockchain on its front end. When someone want to add an idea, he/she can propose the frontend. then that idea is sent to the central node for approval. is the central node approves the idea, then it is added to the Blockchain and also shared with all connected nodes. Else it is ejected.

# Requirements
To run this application, you need Angular, Golang, Node and Mysql installed on your machine.

# Guidelines to run the application
First, run the cnode.go by the following command
  go run cnode.go
It will start listening for other nodes on port 6000, so make sure that port 6000 is not in use.
Now run node.go by the following command
  go run node.go
It will connect to the central node on port 6000 and also creates its own http server which will start listening at port 8081, so make sure port 8081 is not in use.
 Now to start front end, you need to download some node modules, for that, go to the directory of front end and write
  npm install
It will take some time and will install the required node modules.
Now to run the front end, write
  ng serve
Now go to your favourite browser and type
  http://localhost:4200
make sure that port 4200 is not in use. So in this way, you can now use Idea Market, which is based on a custom Blockchain that protects the ownership of ideas.
