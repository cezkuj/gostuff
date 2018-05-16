package main

import (
	"fmt"
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
	subPages *[]Page
}

func getUndetermisticKey(m map[string]string) string {
	for k := range m {
		return k
	}
	return ""
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
	queue := []string{}
	key := getUndetermisticKey(m)
	for ; m[key] != ""; key = m[key] {
		queue = append(queue, key)
	}
	fmt.Println("master page: " + key)
	p := Page{key, &[]Page{}}
	p_head := p
	for i := 0; i != len(queue); i++ {
		page := Page{queue[len(queue)-1-i], &[]Page{}}
		*p.subPages = append(*p.subPages, page)
		p = page
	}
	fmt.Println((*p_head.subPages)[0].subPages)
}
