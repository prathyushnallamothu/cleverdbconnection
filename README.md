# dbconnection
connect to a database in MySql using Golang
You can easily connect to a MySql database in Golang by importing this repository .
simply use go get github.com/prathyushnallamothu/dbconnection in your golang working directory.
after importing change the USERNAME,PASSWORD,and SCHEMA name in the file dbconnection.go with your USERNAME,PASSWORD and SCHEMA name respectively.
in my dbconnection.go you will have it as "root:pass@tcp(localhost:3306)/login" change it as "yourusername:yourpassword@tcp(localhost:3306)/yourschemaname" then only this will work to your mysql database.
and you can use the package to connect to database which is shown in example.go

