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
	fmt.Printf("Insertion Sort: %v --> ", data)
	insertionSort(data)
	fmt.Println(data)
}

func insertionSort(data []int) {
	for i := 0; i < len(data); i++ {
		for j := i; j > 0 && data[j - 1] > data[j]; j-- {
			data[j - 1], data[j] = data[j], data[j - 1]
		}
	}
}

