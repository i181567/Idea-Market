# IdeaMarket #
### Monetising your ideas ###
- - - -

IdeaMarket has three main components.
1. **An Angular (Node.js) based frontend** 
    1. Run npm install followed by ng serve from within the frontend folder
2. **A python backend module that makes use of a pretrained model from huggingface's sentence-transformers**
    1. install dependencies if not installed already such as pip install sentence_transformers and pip install flask 
    2. finally run by "python app.py" or "flask run"
3. **A go backend that handles all block-chain related operations**
    1. Run 'go mod init' followed by 'go mod tidy'
    2. Then use "go run api.go" to get the go backend working.
    3. The go backend is also connected to a mongodb database.

#### Componnet Connections ####
- The Frontend is communicating with both python and go backend.
- Python and go backend are communicating with each other.
- Go backend is communicating with mongodb.





