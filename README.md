# Task Tracker (REFACTORING NOW!)

My First solution for the  [task-tracker](https://roadmap.sh/projects/expense-tracker) challenge from [roadmap.sh](https://roadmap.sh/).

# ID  Date       Description          Amount
# 1   30353-100-01  buy cofe              $3
# 2   30353-100-01  buy dinner            $5
# 3   30354-100-01  buy  pencil           $1

## How to run

Clone the repository, and use following command

```bash
git clone https://github.com/itocode21/ExpenseTracker.git
cd main
```




### Run the following command to build:

```bash
go build main.go
```
-----------------------------

## Run the following command to  run the project:
```bash
# To add a expense
./main add "Description" 10  #10 --> amount 

# To delete expense
./main delete n #|n --> expense id

# To delete a task
./main summary #|list summary all expense

# To list
./main list #|list all expense

# To list expense by month
./main month n #|n --> month 1-12

# export json file with data
./main export #|u get link http://localhost:8080/download ctrl+rmb for dowloand. ctrl+c for exit


```