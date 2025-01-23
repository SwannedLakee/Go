/**
* Problem 2 - Even Fibonacci numbers
* @see {@link https://projecteuler.net/problem=2}
*
* Each new term in the Fibonacci sequence is generated by adding the previous two terms.
* By starting with 1 and 2, the first 10 terms will be:
*
* 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
*
* By considering the terms in the Fibonacci sequence whose values do not exceed four million,
* find the sum of the even-valued terms.
*
* @author ddaniel27
 */
package problem2

func Problem2(n uint) uint {
	sum := uint(0)
	a, b := uint(1), uint(2)

	for b < n {
		if b%2 == 0 {
			sum += b
		}

		a, b = b, a+b
	}

	return sum
}