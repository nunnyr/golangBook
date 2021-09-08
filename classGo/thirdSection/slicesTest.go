package main


fruits = ["apple", "banana", "grape", "cherry"]

fruits[startIndexIncluding: upToNotIncluding]

fruits[0:2] => ["apple", "banana"]

//we can also optionally leave out the second number
so you can also write fruits[:2]
//Go interprets this as the very beginning of the slice
//up to but not including 2
you can also write fruits [2:]
//which would print out ["grape", "cherry"]
//we can use this range syntax to select a subset of cards within
//our deck 

