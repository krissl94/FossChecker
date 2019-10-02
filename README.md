# FossChecker
The FossChecker is a program that can be used to retrieve student data from your university's database based on their student number.

# How to use
## Installing golang
In order to compile the program, you will need golang installed on your system. Follow [this guide](https://golang.org/doc/install) for installing golang. 

## Compiling
You can compile the program by running 
```bash
go build fosschecker.go
./fosschecker.exe
```
or 
```bash
go run fosschecker.go
```

## Using the FossChecker
The FossChecker uses .csv files as input, with a student number on each row. The easiest way to point to the input file is to put it in the same directory as the FossChecker executable, or put it in a path that's easy to remember.

In order to request data from your university, you have to log in to university's MS Office environment and copy a cookie value.
The required cookie can be retrieved by logging in to eur.delve.office.com and copying the value in the X-Delve-AuthEur field.
Use the Application tab in Chrome's devTools to read cookies. 
The program will ask you for this value on each run, it's not stored anywhere. 

The output of the program is another .csv file, which is put in the same directory as the FossChecker executable. It contains all the data that your university was able to provide. Open it in Excel to manipulate the data as needed. 

To run the fosschecker, simply execute the executable, input your input.csv and the cookie value and wait for the program to do it's job.

# Troubleshooting
## Can't open input file
Your path might be incorrect. Try putting the input file in the same directory as the executable and type only "[FILENAME].csv", omitting the rest of the path.

## Can't write to output file
If you have the output.csv open in MS Excel, it is locked and the program can't write output to the file. Close output.csv and try again.