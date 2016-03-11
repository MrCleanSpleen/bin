import os
import time as t
import skilstak.colors as c
import json
def write_json():
    name = input(c.clear + "What is the name of your class? >>> ")
    health = float(input(c.clear + "What is your class's helath? >>> "))
    agility = float(input(c.clear + "What is your class's agility? >>> "))
    shield =float( input(c.clear + "What is your class's defense? >>> "))
    move1 = input(c.clear + "What is the name of your first move? >>> ")
    damage1 = float(input(c.clear + "What is the damage of the move? >>> "))
    accuracy1 = float(input(c.clear + "What is the accuracy of the move? >>> "))
    speed1 = float(input(c.clear + "What is the speed of the move? >>> "))
    move2 = input(c.clear + "What is the name of your second move? >>> ")
    damage2 = float(input(c.clear + "What is the damage of the move? >>> "))
    accuracy2 = float(input(c.clear + "What is the accuracy of the move? >>> "))
    speed2 = float(input(c.clear + "What is the speed of the move? >>> "))
    data = {
        "class":name,
        "health":health,
        "agility":agility,
        "shield":shield
            
        }
    
    json = name + ".json"
    with open(json, "w+") as f:
        pass
    out_file = open(json,"w")    
    json.dump(data,out_file, indent=4)

    
    
