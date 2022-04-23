import json
from urllib import response
from flask import Flask, jsonify
from flask import request
import flask
from flask_cors import CORS, cross_origin

import requests
#from sklearn.feature_extraction.text import TfidfVectorizer
#from sklearn.metrics.pairwise import linear_kernel
# from sentence_transformers import SentenceTransformer, util
from flask_cors import CORS
import json
app = Flask(__name__)
cors = CORS(app)
app.config['CORS_HEADERS'] = 'application/json'

ideaList = []


@app.route("/")
def index():
    return jsonify({"name": "abubakar"})


@app.route('/nlp/<user_id>', methods=['POST'])
def getBlockChain(user_id):
    # validate user id
    if (validate(user_id=user_id or True)):
        print("IM in")
        # in last approach we sent data as query parameter with ?text_a=something&text_b=something
        # as you can see below next two commented lines
        # text_a = request.args.get("text_a")
        # text_b = request.args.get("text_b")
        # now we're sending data as body. so now we select body in thunder client rather than header.
        # header dont support sending of large data body does.

        response = requests.get('http://localhost:8081/listideas')
        # print(response.json()[0]['Description'])

        # dscrpt=[]

        for i in response.json():
            ideaList.append(i['Description'])
            print(i['Description'])

        # print(ideaList)
        # return dscrpt

        # return str(response.text)
        # print("helo")
        # print(str(response.data()))

        # request_data = request.get_json()
        # text_a = request_data["text_a"]
        # #text_array = request_data["text_array"]
        # print("Abubakar is the best")
        # print(text_a["BlockData"]["Description"])

        # return {
            # "data": response
            # None
            # "sameIdea":str(similarityAcrossIdeas(text_a["BlockData"]["Description"])[0]),
            # "score" :str(similarityAcrossIdeas(text_a["BlockData"]["Description"])[1])
        # }
        # [str(similarityAcrossIdeas(text_a["BlockData"]["Description"])[0]),
            # str(similarityAcrossIdeas(text_a["BlockData"]["Description"])[1])]
        # print(text_array)
        # return str(similarityCheck(text_a, text_array))

    return "Invalid User, Event will be logged"


def validate(user_id):
    # validation here
    if (user_id == str(1234)):
        print("User " + str(user_id) + " authentication")
        return True
    print("User " + str(user_id) + " unauthorized")
    logger("Log this bhaiya g")
    return False


def similarityCheck(ideaA, ideaB):
    # nlp algorithm here 😃
    # store similarity check score in 'score variable here'
    # Create TF-idf model...//comment///stop_words=token_stop,tokenizer=tokenizer
    vectorizer = TfidfVectorizer()
    doc_vectors = vectorizer.fit_transform([ideaA] + [ideaB])

    # Calculate similarity
    cosine_similarities = linear_kernel(doc_vectors[0:1],
                                        doc_vectors).flatten()
    document_scores = [item.item() for item in cosine_similarities[1:]]
    # [0.0, 0.287]
    score = 0
    for x in document_scores:
        score = x * 100
        # print(score)

    return score


def similarityAcrossIdeas(ideaA):
    #ideasScore = []
    # for idea in arrayIdea:
    # append score in an array
    #ideasScore.append(similarityCheck(ideaA, idea))
    # return ideasScore

    # dscrpt=[]
    # f = open('data.json',encoding="utf8")
    # data = json.load(f)
    # for i in data['ideas']:
    #     dscrpt.append(i['Description'])
    # f.close()

    # model = SentenceTransformer('sentence-transformers/multi-qa-MiniLM-L6-cos-v1')
    # print(ideaA)
    # print(ideaList)
    # Encode query and documents
    query_emb = model.encode(ideaA)
    doc_emb = model.encode(ideaList)

    # Compute dot score between query and all document embeddings
    # scores = util.dot_score(query_emb, doc_emb)[0].cpu().tolist()

    # Combine docs & scores
    doc_score_pairs = list(zip(ideaList, scores))

    # Sort by decreasing score
    doc_score_pairs = sorted(doc_score_pairs, key=lambda x: x[1], reverse=True)
    return doc_score_pairs[0]


def logger(loggedText):
    # logging code here
    print("Logging: " + str(loggedText))


@app.route('/proposeidea', methods=['POST'])
@cross_origin()
def abc():
    #response = requests.get('http://localhost:8081/listideas')
    #requests.post("http://localhost:8081/proposeidea", request.data,)
    # print(request.json['Description'])

    # print(request.data)

    print(request.json['Description'])
    ideaA = request.json['Description']
    # print(ideaA)
    # print(ideaList)
    # return request.data

    # score_txt = similarityAcrossIdeas(ideaA) #UNCOMMENT THIS PLEASE

    # print(score_txt)

    proposedIdea= {
      "Title": request.json['Title'],
      "Description": request.json['Description'],
      "Owners": request.json['Owners'],
      "Problem": request.json['Problem'],
      "Domain": request.json['Domain'],
      "Technologies_used": request.json['Technologies_used'],
      "Viewing_price": float(request.json['Viewing_price']),
      "Ownership_price": float(request.json['Ownership_price']),
      "Pricing_history": request.json['Pricing_history'],
    #   "Score_text": score_txt[0],
      "Score_text": score_txt[0],
      "Score": score_txt[1]
    }
    print(type(proposedIdea["Title"]))
    print(type(proposedIdea["Score_text"]))
    print(type(proposedIdea["Score"]))
    print(type(proposedIdea["Viewing_price"]))
    print(type(proposedIdea["Ownership_price"]))
    print(type(proposedIdea["Pricing_history"]))


    # proposedIdea = {
    #     'BlockData': {
    #         "Title": request.json['Title'],
    #         "Description": request.json['Description'],
    #         "Owners": request.json['Owners'],
    #         "Problem": request.json['Problem'],
    #         "Domain": [],
    #         "Technologies_used": request.json['Technologies_used'],
    #         "Viewing_price": request.json['Viewing_price'],
    #         "Ownership_price": request.json['Ownership_price'],
    #         "Pricing_history": request.json['Pricing_history']
    #     },
    #     # "SimIdea": score_txt[0],
    #     "SimIdea": "Place holder Idea",
    #     # "SimScore": score_txt[1]
    #     "SimScore": 0.12
    # }

    # proposedIdea = {
    #     "Title":       request.json['Title'],
    #     "Description": request.json['Description'],
    #     "Owners": request.json['Owners'],
    #     "Problem": request.json['Problem'],
    #     "Domain": request.json['Domain'],
    #     "Technologies_used": request.json['Technologies_used'],
    #     # "Viewing_price": request.json['Viewing_price'],
    #     "Viewing_price": 0,
    #     # "Ownership_price": request.json['Ownership_price'],
    #     "Ownership_price": 1,
    #     "Pricing_history": request.json['Pricing_history'],
    #     "Score_text": "Some score text",
    #     "Score": 1.28
    # }

    # this is was online that was required to convert it into the json and send
    proposedIdea = json.dumps(proposedIdea)

    print("Sending...")
    print(proposedIdea)

    z = requests.post("http://localhost:8081/proposeidea", data=proposedIdea)

    print("zzzz", z.text)
    return jsonify({"data": z.text})
    # http://localhost:8081/getPendingIdeas


if __name__ == "__main__":
    getBlockChain("1234")
    app.run(debug=True, host="localhost", port=8082)
