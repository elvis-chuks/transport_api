import requests
import json

params = {
    'depotcode':"ABA",
    "routeID":"23",
    "busseatarrangementid":"1",
    'busQueueID':'2833',
    "busclassid":'2',
    'departuredate':'2019-04-04',
    "phonenumber":'08100726139'
}

# params = json.dumps(params)

resp = requests.get("http://127.0.0.1:5000/v3/checkbook",headers=params).json()
print(resp)