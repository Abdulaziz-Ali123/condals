# condals
A simple cli to saving conda enviorment paths to enable faster activation of any conda enviorments.

## Requirements
Golang version 1.24.6


## How to Setup 
In your terminal run

``` bash
git clone https://github.com/Abdulaziz-Ali123/condals.git
cd condals
go mod tidy
```

## How to run
Run the following in your terminal:
``` bash
go run . [flags can go here]
```
```
```


## How to Build and Add to paths
After ensuring the program runs properly

for linux run the following in your terminal:
``` bash
go build builds/condals
sudo mv buils/condals /usr/local/bin/ 
```

Now you will be able to use the `condals` command from anywhere in your terminal
