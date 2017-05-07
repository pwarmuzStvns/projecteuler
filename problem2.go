/*
   @Project: 			projecteuler
   @Author: 			Phil
   @Date:   			2017-05-06 15:58:27
   +Last Modified by:   Phil
   +Last Modified time: 2017-05-06 21:06:35
*/

package main

var sum uint

func problem2() uint {
	fib := fibonacci()
	//  1000 is a wild guess I'd like to get this accurate
	for i := 0; i < 1000; i++ {
		fib(4000000)
	}
	return sum
}

func fibonacci() func(uint) uint {
	/*  Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:
	    1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
	    By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.
	*/
	/*  Go makes a beautiful fibonacci seq*/
	var f, s uint
	f, s = 0, 1
	return func(limit uint) (ret uint) {
		if f >= limit {
			return
		}
		if f%2 == 0 {
			sum += f
		}
		ret, f, s = f, s, f+s
		return
	}
	//4613732
}
