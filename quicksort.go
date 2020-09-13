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
	fmt.Printf("%v --> ", data)
	quickSort(data)
	fmt.Println(data)
}

func quickSort(data []int) {
	if len(data) > 1 {
		pivotIndex := partition(data)
		quickSort(data[:pivotIndex])
		quickSort(data[pivotIndex:])
	}
}

func partition(data []int) int {
	pivotValue := data[len(data)-1]
	leftIndex := -1
	for i := 0; i < len(data); i++ {
		if data[i] <= pivotValue {
			leftIndex++
			data[leftIndex], data[i] = data[i], data[leftIndex]
		}
	}
	return leftIndex
}
