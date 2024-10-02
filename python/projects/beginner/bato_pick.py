import random

def evaluate(user_input, computer_input):
    draw = user_input == computer_input
    if draw:
        print("It's a Draw!")
        return 0,0

    user_wins = (user_input == "paper" and computer_input == "rock") or (user_input == "scissors" and computer_input == "paper") or (user_input == "rock" and computer_input == "scissors")
    if user_wins == True:
       print("You Won!")
       return 1,0
    else:        
       print("You Lost!")
       return 0,1

user_points = 0
computer_points = 0
options = ["rock", "paper", "scissors"]

while True:
    user_input = input("Rock Paper Scissors or Quit: ").lower()

    if user_input == "q":
        break

    if user_input not in options :
        continue

    random_number = random.randint(0,2)
    # rock: 0, paper:1, scissors: 2
    computer_pick = options[random_number]
    print("Computer picked", computer_pick + ".")
    user_point, computer_point = evaluate(user_input, computer_pick)  
    user_points += user_point
    computer_points += computer_point
    print(f"User: {user_points}, Computer: {computer_points}")

print("Bye Bye!")


    
