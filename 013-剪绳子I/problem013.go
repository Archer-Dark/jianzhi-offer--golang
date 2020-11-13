package main

//数论法1：
//	任何大于1的数都可由2和3相加组成（根据奇偶证明）
//	因为2*2=1*4，2*3>1*5, 所以将数字拆成2和3，能得到的积最大
//	因为2*2*2<3*3, 所以3越多积越大 时间复杂度O(n/3)，用幂函数可以达到O(log(n/3)), 因为n不大，所以提升意义不大，我就没用。 空间复杂度常数复杂度O(1)

func cuttingRope(n int) int {
	if n <= 3{
		return n-1
	}
	div := n/3
	rem := n%3
	var result int64 = 1
	for i := 0;i<div;i++{
		if i < div-1 {
			result *= 3
		}else{
			if rem == 2{
				result *= (int64)(3*rem)
			}
			result *= (int64)(3+rem)
		}
		if result >= 1000000007{
			result = result&1000000007
		}
	}
	return int(result)
}

//数论法2：
//	绳子段切分的越多，乘积越大，且3做因数比2更优。
//	不断剪去 长度3 并用sum累乘
//	把绳子切为多个长度为 3 的片段，则留下的最后一段绳子的长度可能为 0,1,2 三种情况。
//	循环结束的三种情况：
//	    n=n-3==2
//	n长度剩下2，因 n > 4 跳出循环，return 乘积为sum*2。
//	    n=n-3==3
//	长度刚好可以被剪完，因 n > 4 跳出循环，return 乘积为sum*3。
//	    n=n-3==4
//	再剪一次的话，长度剩下1，则最后一段绳子长度为 1； 需要把最后的 3 和 1 替换为 2 和 2，因为 2 * 2 > 3 * 1； 因 n > 4 跳出循环，return 乘积 sum*4 即可。
func cuttingRope2(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	sum := 1
	for n > 4{
		sum *= 3
		n -= 3
	}
	return sum * n
}