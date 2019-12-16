## This week's question:
Given a positive number N, generate binary numbers between 1 to N using a queue-type structure in linear time.

Example:

> binaryQueue(10)
> 1 10 11 100 101 110 111 1000 1001 1010 1011 1100 1101 1110 1111 10000

## Solutions

`binaryQueueConvert` takes an int, loops from 1 -> `n`, converts them to binary, stores them on the list, and returns the list. Binary display is merely displaying the element with `fmt.Println("%s", element)`.

`binaryQueueDisplay` takes an int, loops from 1 -> `n`, stores them on the list, returns the list. Binary display is used by `fmt.Sprintf("%b", element)`.
