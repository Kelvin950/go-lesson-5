package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	s1 := "hi there go!"
   
	 fmt.Println(s1) ;

	 s2:= `I like go `
	 fmt.Println(s2) 
   
	 var s3 = "Iaove " + "Go" + "programming" 
	 fmt.Printf(s3 +  `!`)	

	  fmt.Println("Element at index 0: " , s3[1])
	fmt.Println(s3[0]- 65)
	  
  
	str :=  "ţară"

 
	 fmt.Println(len(str))
	 fmt.Printf( "%v",str[1])
 

 for i:=0 ; i< len(str);i++{

 fmt.Printf("%c\n" , str[i])
 } 

  

 for i:=0 ; i< len(str);{
    r,size := utf8.DecodeRuneInString(str[i:]) 

	fmt.Printf("%c" , r) ;
	i += size 
 } 

 str2 :=  "casablanca" 

 for i:=0 ; i< len(str2);{
 

	r, size := utf8.DecodeRuneInString(str2[i:])
  
 
	fmt.Printf("%c" ,r) ;
	i+= size ;
 }
  

 r,size := utf8.DecodeRuneInString(str2[1:])

 fmt.Printf("%c , %v", r, size) 

 n := utf8.RuneCountInString(str) 
 m := utf8.RuneCountInString(str2) 

 fmt.Println(n ,m) 
	}