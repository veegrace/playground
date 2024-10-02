def view():
    with open('passwords.txt', 'r') as f:
        for line in f.readlines():
            print(line.rstrip())

def add():
    name = input("Account Name: ")
    password = input("Password: ")

    # w - writes
    # r - read only
    # a - pen mode, creates file if not exists, else read 
    with open("passwords.txt", "a") as f:
        f.write(name + "|" + password + "\n")
    


password = input("What is the master password?")

while True:    
    mode = input("Would you like to add a new password or view existing ones? (view/add/q)")

    if mode == "q":
        break
    if mode == "view":
        view()
    elif mode == "add":
        add()
    else:
        print("Invalid Mode.")
        continue