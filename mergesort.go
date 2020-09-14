/*
Copyright 2020 Brian Pursley

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import "fmt"

func main() {
	data := []int{2, 8, 7, 1, 3, 5, 6, 4}
	fmt.Printf("Merge Sort: %v --> ", data)
	mergeSort(data)
	fmt.Println(data)
}

func mergeSort(data []int) {
	mergeSortInternal(data, 0, len(data)-1)
}

func mergeSortInternal(data []int, i, j int) {
	if i < j {
		mid := (i + j) / 2
		mergeSortInternal(data, i, mid)
		mergeSortInternal(data, mid+1, j)
		merge(data, i, j, mid)
	}
}

func merge(data []int, i, j, mid int) {
	left := make([]int, mid-i+1)
	copy(left, data[i:mid+1])
	right := make([]int, j-mid)
	copy(right, data[mid+1:j+1])
	var a, b int
	c := i
	for a < len(left) && b < len(right) {
		if left[a] <= right[b] {
			data[c] = left[a]
			a++
		} else {
			data[c] = right[b]
			b++
		}
		c++
	}
	for a < len(left) {
		data[c] = left[a]
		a++
		c++
	}
	for b < len(right) {
		data[c] = right[b]
		b++
		c++
	}
}
