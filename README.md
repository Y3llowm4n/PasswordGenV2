# PasswordGen

How to use this application:

Use -l to specify amount of lowercase letters in password
Use -u to specify amount of uppercase letters in password
Use -s to specify amount of special characters in password
use -p to specify total length of password. There is a bug with this feature. Read more about it at the bugs chapter.

Example in Terminal: go run . -l 3 -u 3 -s 5 -p 12
after this enter username and press enter to recieve a password.


`Bugs

1. If given more characters in total with lowercount, uppercount and specialcount than lengtpassword it'll generate the password length based on the lower-, upper- and specialcount.

