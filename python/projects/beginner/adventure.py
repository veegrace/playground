import random

print("The free service of planetary system 8612 has been terminated...")
print("The main scenario starts now...")

print("\n\n[NOTICE] #BI-7623 Channel is open.")
print("[NOTICE] The Constellations have entered.")

player = input("Player Name: ")
player+="!"
print("Welcome incarnator", player, "You have to complete the main scenario in order to survive this realm!")

print("\n[MAIL] The Main scenario has arrived")
print("\n\n[MAIL] MAIN SCENARIO #1\n\n")
print("       [PROOF OF VALUE]\n")
print("      Kill the DEMON KING.\n\n")
print("           Category: MAIN")
print("           Dificulty: F")
print("           Time Limit: 30m")
print("           Reward: 300 Coins")
print("           Failure: DEATH\n\n")

new_demon_king_options = ["kim dokja", "han sooyoung", "yoo joonhyuk", "yoo sangah", "jung heewon"]
random_number = random.randint(0,2)
new_demon_king = new_demon_king_options[random_number]
print(f"{player}! It seems that the DEMON KING have been dethroned by incarnator {new_demon_king}")
action = input("Will you attack?(y/n): ")

if new_demon_king == "kim dokja" and action == "yes":
    print("You have killed the DEMON KING!")

elif new_demon_king == "kim dokja" and action == "no":
    print("You have died!")

elif new_demon_king != "kim dokja": 
    print("Scenario have been concluded!")