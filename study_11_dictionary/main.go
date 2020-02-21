package main

import (
	"fmt"
	"study_09_goLang/study_11_dictionary/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	word := "hello"
	dictionary.Add(word, "First")
	dictionary.Update(word, "Second")
	output, _ := dictionary.Search(word)
	fmt.Println(output)
	dictionary.Delete(word)
	fmt.Print(dictionary.Search(word))
}
