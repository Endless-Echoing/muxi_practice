package main

import (
	"fmt"
)

func printnum(letterCh chan struct{}, numberCh chan struct{}, done chan struct{}) {
	for c := 'A'; c <= 'Z'; c++ {
		<-letterCh
		fmt.Printf("%c", c)
		numberCh <- struct{}{}
	}
}

func printletter(letterCh chan struct{}, numberCh chan struct{}, done chan struct{}) {
	for i := 1; i <= 26; i++ {
		<-numberCh
		fmt.Printf("%d", i)
		if i < 26 {
			letterCh <- struct{}{}
		} else {
			close(done)
		}
	}
}
