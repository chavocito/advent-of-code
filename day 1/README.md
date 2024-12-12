### I'll develop this later

- to solve the challenge, I went traditional
	- i kept the input in a txt file.
	- i read the file using the os package
	- i wrote a helper to read the file line by line, and split the strings around each instance of one or more consecutive white space characters
	- inside of the scan loop, i push the parts into a left and right slice if I have two parts
	- at the end of the loop, i return the two slices
	- now with the two arrays, i do a sort, peep their length, and do the difference accumulation

### revisit this in future and do some refactor
