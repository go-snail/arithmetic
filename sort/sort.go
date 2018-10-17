package sort

import (

)

//插入排序算法
func InsertionSort(in []int)  (e error) {
	var temp int
	for i := 1; i < len(in); i++ {
		for j := i; j > 0 && in[j] < in[j-1]; j-- {
			temp = in[j]
			in[j] = in[j-1]
			in[j-1] = temp
		}
	}
	return nil
}


//希尔排序算法
func ShellSort(in []int)  (e error) {
	//使用希尔推荐的增量 h(k) = N/2
	var temp int
	h := len(in)
	for h >= 1 {
		for i := h;i < len(in);i += h {
			for j := i;j >= h && in[j] < in[j-h]; j -= h {
				temp = in[j]
				in[j] = in[j-h]
				in[j-h] = temp
			}
		}

		h /= 2
	}
	return nil
}

//冒泡排序
func BubbleSort(in []int) {
	var max int
	for i:=0; i< len(in) - 1; i++ {
		for j:=0; j < len(in)-i-1;j++ {
			if in[j] > in[j+1] {
				max = in[j]
				in[j] = in[j+1]
				in[j+1] = max
			}
		}
	}
}

//快速排序
func QuickSort(in []int, l, r int)  {

	if l < r {
		var i,j,x int
		i = l
		j = r
		x = in[i]
		for i < j {
			for i < j && in[j] > x {
				j--
			}
			if i < j {
				in[i]=in[j]
				i++
			}
			for i < j && in[i] < x {
				i++
			}
			if i < j {
				in[j] = in[i]
				j--
			}

		}
		in[i] = x
		QuickSort(in,l, i-1)
		QuickSort(in, i+1,r)
	}
}

//选择排序
func SelectionSort(in []int)  {
	var minIndex int
	var temp int
	for i := 0;i< len(in) - 1; i++ {
		minIndex = i
		for j := i+1; j < len(in); j++  {
			if in[minIndex] > in[j] {
				minIndex = j
			}
		}
		temp = in[i]
		in[i] = in[minIndex]
		in[minIndex] = temp
	}
}