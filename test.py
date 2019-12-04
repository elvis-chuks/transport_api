import requests
import json

params = {
    'depotcode':"ABA",
    "routeID":"15"
}

# params = json.dumps(params)

resp = requests.get("http://127.0.0.1:5000/v3/discount",headers=params).json()
print(resp)