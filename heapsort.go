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
	fmt.Printf("Heap Sort: %v --> ", data)
	heapSort(data)
	fmt.Println(data)
}

func heapSort(data []int) {
	buildMaxHeap(data)
	for i := len(data) - 1; i > 0; i-- {
		data[0], data[i] = data[i], data[0]
		maxHeapify(data[:i], 0)
	}
}

func buildMaxHeap(data []int) {
	for i := len(data)/2 - 1; i >= 0; i-- {
		maxHeapify(data, i)
	}
}

func maxHeapify(data []int, i int) {
	leftChildIndex := i*2 + 1
	rightChildIndex := i*2 + 2
	var largestIndex int
	if leftChildIndex < len(data) && data[leftChildIndex] > data[i] {
		largestIndex = leftChildIndex
	} else {
		largestIndex = i
	}
	if rightChildIndex < len(data) && data[rightChildIndex] > data[largestIndex] {
		largestIndex = rightChildIndex
	}
	if largestIndex != i {
		data[i], data[largestIndex] = data[largestIndex], data[i]
		maxHeapify(data, largestIndex)
	}
}
