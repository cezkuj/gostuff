package main

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
	"strings"
	"testing"
)

/*
change
e -> d
d -> c
d2 -> c
c -> b
c2 -> b
c3 -> b
b -> a

into

a
  b
    c
      d
        e
      d2
    c2
    c3
*/

type Page struct {
	name     string
	subPages map[string]Page
}

func getKey(m map[string]string) string {
	for key := range m {
		return key
	}
	return ""
}

func createStack(page_name string, m map[string]string) (*stack.Stack, string) {
	st := stack.New()
	for ; m[page_name] != ""; page_name = m[page_name] {
		st.Push(page_name)
	}
	return st, page_name
}
func TestMigrate(*testing.T) {
	m := map[string]string{
		"e":  "d",
		"d":  "c",
		"d2": "c",
		"c":  "b",
		"c2": "b",
		"c3": "b",
		"b":  "a",
	}
	key := getKey(m)
	_, masterPage := createStack(key, m)
	p_head := Page{masterPage, make(map[string]Page)}
	for key := range m {
		p := p_head
		st, _ := createStack(key, m)
		for st.Len() > 0 {
			page_name, ok := st.Pop().(string)
			if !ok {
				fmt.Println("not ok")
			}
			if _, present := p.subPages[page_name]; !present {
				page := Page{page_name, make(map[string]Page)}
				p.subPages[page.name] = page
				p = page
			} else {
				p = p.subPages[page_name]
			}
		}
	}
	fmt.Println(p_head)
	printPages(p_head, 0)
}
func printPages(page Page, depth int) {
	depth++
	fmt.Println(strings.Repeat(" ", depth*4), depth, page.name)
	for _, subpage := range page.subPages {
		printPages(subpage, depth)
	}
}
