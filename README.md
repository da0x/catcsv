# catcsv

View your CSV files in your terminal. The output is a formatted table.

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

