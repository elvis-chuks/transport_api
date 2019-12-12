# import requests
# import json

# params = {
#     'depotcode':"ABA",
#     "routeID":"23",
#     "busseatarrangementid":"1",
#     'busQueueID':'2833',
#     "busclassid":'2',
#     'departuredate':'2019-04-04',
#     "phonenumber":'08100726139'
# }

# # params = json.dumps(params)

# resp = requests.get("http://127.0.0.1:5000/v3/checkbook",headers=params).json()
# print(resp)

def magic(card, num):
    cardinal = ''
    for i in card:
        if i == ' ':
            pass
        else:
            cardinal += i
    number = ''
    for i in num:
        if i == ' ':
            pass
        else:
            number += i
    seatCardinal = cardinal.split(',')
    seatNumber = number.split(',') #with ',' as the delimiter
    resultDict = {}
    for i in range(len(seatNumber)):
        resultDict.update({seatNumber[i]: seatCardinal[i]})
    return resultDict

rid = magic("1,2,3,4", "1_1, 1_2, 1_3,1_4")
print(rid,"yes")