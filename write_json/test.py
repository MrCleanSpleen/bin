import os
import time as t
import skilstak.colors as c
import json
def write_json():
    name = input(c.clear + "What is the name of your class? >>> ")
    health = float(input(c.clear + "What is your class's helath? >>> "))
    agility = float(input(c.clear + "What is your class's agility? >>> "))
    shield =float( input(c.clear + "What is your class's defense? >>> "))
    data = {
        "class":name,
        "health":health,
        "agility":agility,
        "shield":shield
            
        }
    
    json = name + ".json"
    with open(json, "w+") as f:
        out_file = open(json,"w")    
        json.dump(data,out_file, indent=4)
