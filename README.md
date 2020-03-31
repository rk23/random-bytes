# Random Bytes

This program attempts to mimic the linux command 

`cat /dev/random` 

by outputting random bytes to standard out. It uses CPU time and used memory polled at random intervals as sources of entropy and SHA512 as the hashing function. 

## Running

`go run ./cmd/dev-random/main.go`

Similar to `cat /dev/random`, you'll need to hit Control-C to stop the program.

## Further Development

Entropy can be added to the PRNG by simply adding another goroutine that dumps random bytes of your choice into the `input` channel specifed in `/cmd/dev-random/main.go`

## Resources
A lot of the information for  this project was pulled from the following lecture: https://www.youtube.com/watch?v=0DV8WnqhH2Y