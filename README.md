# catcsv

This is a simple command line tool to view CSV files in your terminal. The output will be formatted into an easy to view box.

## Installation

Run the following commands:

```
go get github.com/da0x/catcsv
go install github.com/da0x/catcsv
```

## Usage

Create a CSV File.
```
$ cat file.csv 
hello,world
1,2
10,20
```

Now you can read it directly or use it in a pipe.
```
$ catcsv file.csv 
┌───────────────┐
│ file.csv      │
├───────┬───────┤
│ hello │ world │
│ 1     │ 2     │
│ 10    │ 20    │
└───────┴───────┘
$ cat file.csv | catcsv -
┌───────┬───────┐
│ hello │ world │
│ 1     │ 2     │
│ 10    │ 20    │
└───────┴───────┘
```

