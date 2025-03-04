# concurrency-patterns
Three main concurrency patterns in go are:

## For-select loop 
* Buffered channel is made if we give a capacity in the make
* For asynchronous communication we need to have a buffered channel
* An unbuffered channel provides a guarante that an exchange between 2 go-routines is performed at the instant the send and receive take place : a receiving goroutine have to wait for the sending goroutine (syncronous communication)
* With buffered channel we use a queue like functionality where we can send data to channel and forget upto the alloted capacity 

## Done channels 
* Mechanism for main go routine to cancel the infinitely running go routine
<img src="channels.png"/>

## Pipelines
* Communication is syncronous in a pipeline 
<img src="pipeline.png"/>