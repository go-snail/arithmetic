package leecode

// 经典的水王问题
// 某一个数在给定的系列数中出现的次数大于一半及一半以上，那么这个数为水王
// []int{1,2,3,3,34,4,2,3,3}
func waterKing(src []int) int {
	var condidate = -1
	var blood = 0
	for _,v := range src {
		if v == condidate {
			blood++
		} else {
			if blood == 0 {
				condidate = v
			} else {
				blood--
			}
		}
	}
	return condidate
}


