package main

import (
	"fmt"
)

//交换函数的几种写法
/*
func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
func swap(arr []int, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp

}
func swap(arr []int, i, j int) {
	arr[i] = arr[i] + arr[j]
	arr[j] = arr[i] - arr[j]
	arr[i] = arr[i] - arr[j]

}
*/
func swap(arr []int, i, j int) {
	//交换元素（异或）
	arr[i] = arr[i] ^ arr[j]
	arr[j] = arr[j] ^ arr[i]
	arr[i] = arr[i] ^ arr[j]
}
func selectionSort(arr []int) {

	for i := 0; i < len(arr)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ { //因为minIndex起始是i，所有j比较从i+1开始
			if arr[j] < arr[minIndex] {
				minIndex = j //记录最小数的下标
			}
		}
		if minIndex != i { //因为采取的异或交换，若两个值相等，会变成0，所有进行判断
			swap(arr, i, minIndex)
		}
	}
}

func main() {
	arr := []int{3, 44, 38, 5, 47, 15, 36, 34}
	selectionSort(arr)
	fmt.Println(arr)
}
