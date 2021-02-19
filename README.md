# export-select-to-insert-golang


#### How to use

> clone project

```
git@github.com:verytalk/export-select-to-insert-golang.git
```


> Linux

- edit config file config.yaml
- running the file main For example:

```
./main
```

> Other execution ways

- Install golang environment
- running script main.go For example

```
go run main.go 
```


> If you don't want to use the default config file config.yaml

```
./main --config test.yaml

or

go run main.go --config test.yaml

```

> Explain config file

```
mysql:
  user : root  				# mysql user
  password : root 			# mysql password
  host : 127.0.0.1 			# mysql host
  port : 3306 				# mysql port
  name : test				# mysql database name
logfile:
  filename: Logs.sql   		# export SQL to the filename 
export:
  tableName: Logs 			# generate insert to table
  exportSQL: SELECT * FROM Logs  # SELECT statement
  exportOpenPaging: true		 # using paging true/false
  exportPagingStart: 0			 # start page
  exportPagingEnd: 400			 # end page
  exportPagingLimit: 20000 		 # length per page
```



