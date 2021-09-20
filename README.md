How to run the backend:

1) https://golang.org/doc/install
2) After installation, run the command "go env" in terminal/cmd/powershell to see if Go installed properly
3) Clone the github repo (https://github.com/adecenzo/picklejar)
4) Navigate to the picklejar directory with the "main.go" file
5) to start the web server, run the command "go run main.go"
6) web server is configured to point to http://localhost:8080/

How to run the frontend:

1) install npm
2) navigate to the cloned picklejar directory then the pj-frontend folder
3) if the first time running the frontend, run the "npm install" cmd. If not, run "npm start"
4) frontend is currently configured to point to http://localhost:3000/


How to configure the AWS SNS service:
(spoke with TA, it is fine if this feature does not work on the TAs end because of the access keys required)

1) install AWS CLI
2) open terminal/powershell/cmd and run "aws configure"
3) put your access key & secret key into the fields
4) for region, input "us-east-2"
5) ensure the ~/.aws files exist

Tools to test REST API:
- postman

JSON Formats:

For validate-pin endpoint:

// pin and user_id are string types
// connect to the database using DBeaver or some other database manager with
// the credentials in the db.go file
{"pin": "1234", "user_id" : "103"}

For register endpoint:

{"user": {"first_name": "", "last_name": "", "phone_number": "", "email": "", "password": ""}}

For login endpoint:

{"email": "", "password": ""}


CORS Browser

- If CORS error with the post/get requests, please allow CORS in your web browser.
