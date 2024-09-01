import random

def lost(rand):
    print("Wrong! It's", rand)
    print("You lose~ Better luck next time!")
    quit()

def check_digit(guess):
    if guess.isdigit():
        guess = int(guess)

        if guess <= 0:
            print("Please type a number larger than 0 next time.")
    else:
        print("Please type a number next time.")

print("Welcome to Number Guessing Game!")
print("Try to guess a number from 1-10.")

rand = random.randrange(1,10)
question = "Your Guess: "
tries = 5
     
while True:
    guess = input(question)
    check_digit(guess)

    if guess == rand:
        print("Correct! It's", rand)
        break

    tries -= 1
    print(f"Let's try again! You have {tries} left")

    if tries == 0:
        lost(rand)


