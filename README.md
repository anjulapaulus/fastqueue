[![Build Status](https://travis-ci.com/anjulapaulus/fastqueue.svg?branch=master)](https://travis-ci.com/anjulapaulus/fastqueue)
[![codecov](https://codecov.io/gh/anjulapaulus/fastqueue/branch/master/graph/badge.svg)](https://codecov.io/gh/anjulapaulus/fastqueue)

# fastqueue
Fast numeric-based binary heap priority Queue in golang. It stores the queue as two flat arrays of item ids and and their numeric priority values.


#### Installation
````
go get -u github.com/anjulapaulus/fastqueue
````


#### Acknowledgement
The current implementation is inspired on the javascript library [fastqueue](https://github.com/mourner/flatqueue) by Vladimir Agafonkin.

#### Implementation

````
package main

import queue "github.com/anjulapaulus/fastqueue"


func main(){
	q:= queue.Queue{}   //Initialize Queue

	q.Push(2)      //Insert Items to queue
	q.Push(5)
	q.Push(10)

	map := q.All()      //returns map of elemeents with ids and values in queue

	q.Peek()            //returns the top item id without removal

    q.PeekValue()       //returns the top item value without removal

	q.Len()             //length of queue

	q.Pop()             //returns the top item id and value and removes
}
