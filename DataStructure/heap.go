package main

import "fmt"

//Maxheap struct has a slice that holds the array
type MaxHeap struct{
	array []int
}

//insert adds an element to the heap
func (h *MaxHeap) insert(key int){
	h.array=append(h.array,key)

	h.maxHeapifyUp(len(h.array)-1)
}
//Extract returns the largest key, add removes it from the heap
func (h *MaxHeap) extract() int{
	extracted:=h.array[0]
	l:=len(h.array)-1
	h.array[0]=h.array[l]
	h.array=h.array[:l]
	if l==0{
		fmt.Println("cannot extract because array length is 0")
		return -1;
	}
	h.maxHeapifyDown(0)


	return extracted
}

//maxHeapifyUP will heapify from bottom top
func (h *MaxHeap) maxHeapifyUp(index int){
	for h.array[parent(index)]<h.array[index]{
		h.swap(parent(index),index)
		index=parent(index)
	}

}
//maxHeapifyDown will heapify from top bottom
func (h* MaxHeap) maxHeapifyDown(index int){
	//loop while index has at least one child
}

//get the parent index
func parent(i int) int{
	return (i-1)/2
}
//get the left child index
func left(i int) int{
	return 2*i+1
}
//get the right child index
func right(i int) int{
	return 2*i+2
}
//swap keys in the array
func (h*MaxHeap) swap(i1,i2 int){
	h.array[i1],h.array[i2]=h.array[i2],h.array[i1]
}
func main() {
	m:=&MaxHeap{}
	// fmt.Print(m)
	buildHeap:=[]int{10,20,30,6,8,40,3,2,70}
	for _ ,v:=range buildHeap{
		m.insert(v)
		fmt.Println(m)
	}
	
}
