This is simple program which reads provided data file into memory and then calculates sha512 hash for each line in file.
Result is written into tmp file. Also single.out, pool.out and multi.out trace files are produced.

run programs:
```
go run samples/single/main.go some_data_file

go run samples/multiple/main.go some_data_file

go run samples/pool/main.go some_data_file
```


Then run for each \*.out file:
```
go tool trace single.out
```

And open chrome browser to see trace ui interface.
