print("Welcome to Riddle Game!")

wantPlay = input("Do you want to play? (Y/N) ")
if wantPlay.lower() != "y":
    print("Too bad! Let's play next time!")
    quit()
print("Let's Play!")

score = 0

# First Round
question = "I am a color, but you can eat me. What am I? "
answer = "orange"
if input(question).lower() != answer:
    print("Incorrect!")
    tries = 4      
    print(f"Let's try again! You have {tries} left")
    while input(question) != answer:
        tries -= 1
        print("Let's try again! You have ", tries)
        if tries == 0:
            print("Your current score is ", score)
            print("You lose~ Better luck next time!")
            quit()
print("Correct!")
score += 10

# Second Round
wantPlay = input("Do you want to play another round? (Y/N) ")
if wantPlay.lower() != "y":
    print("Too bad! Let's play again next time!")
    quit()
print("Let's Play! Round 2")
print("Your current score is ", score)

question = "I am Unstoppable, but I am easy to waste. What am I? "
answer = "time"
if input(question).lower() != answer:
    print("Incorrect!")
    tries = 4      
    print(f"Let's try again! You have {tries} left")
    while input(question) != answer:
        tries -= 1
        print("Let's try again! You have ", tries)
        if tries == 0:
            print("You lose~ Better luck next time!")
            quit()
print("Correct!")
score += 10

print("New High Score ", score)