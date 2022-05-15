package main

import "fmt"

type stack []string

func (q *stack) push(name string) {
	*q = append(*q, name)
}
func (q *stack) size() int {
	return len(*q)
}
func (q *stack) pop() string {
	s := q.size()
	data := (*q)[s-1]
	*q = (*q)[0 : s-1]
	return data
}
func main() {
	q := make(stack, 0)
	q.push("1")
	q.push("3")
	q.push("5")
	q.push("7")
	fmt.Println("List")
	fmt.Println(q)

	r := q.pop()
	fmt.Println("Pop ", r)
	r1 := q.pop()
	fmt.Println("Pop ", r1)
	fmt.Println(q)
}
