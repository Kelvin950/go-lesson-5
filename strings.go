package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"

	"unicode/utf8"
)

func  comma(s string) string{
 var buf bytes.Buffer 
  
 i := 1
 for  _, v :=range s {

	 if !unicode.IsDigit(v){
		  buf.WriteRune(v)
		continue ;
	 }
   if  i == 3 {
	buf.WriteString(",")
	 i = 0 
   }
   buf.WriteRune(v)
 i++ ;

 }

 
 return buf.String()

}

func commaWithDec(s string) string{


 parts := strings.Split(s, ".") 
intpart :=  parts[0] 


decimalPart :=""
 
if len(parts) > 1 {

	decimalPart = "."+parts[1] ;
}

var buf bytes.Buffer
n := len(intpart) ;

 if n <=3 {
     
	buf.WriteString(intpart + decimalPart) ;

 } else {

   
	for i , v:=range intpart {
 

		if i > 0  && (n-i)%3 ==0 {
      
			buf.WriteString(",")

		} 

		buf.WriteRune(v)


	}


	buf.WriteString(decimalPart)

 }

 return buf.String()


}
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

 
	 fmt.Println(len(str), "fd")
	 fmt.Printf( "%v",str[1])
 

 for i:=0 ; i< len(str);i++{

 fmt.Printf("kk %c\n " , str[i])
 } 



 for i:=0 ; i< len(str);{
    r,size := utf8.DecodeRuneInString(str[i:]) 

	fmt.Printf("%c" , r) ;
	i += size 
 } 

 y := "abc" 
 b := []byte(y) 


 s4 := string(b)

fmt.Println(b , s4) 

fmt.Println(strings.Count(y ,"bc"))

x :=commaWithDec("1234567890.12345")
fmt.Println(x)





//  str2 :=  "casablanca" 

//  for i:=0 ; i< len(str2);{
 

// 	r, size := utf8.DecodeRuneInString(str2[i:])
  
 
// 	fmt.Printf("%c" ,r) ;
// 	i+= size ;
//  }
  

//  r,size := utf8.DecodeRuneInString(str2[1:])

//  fmt.Printf("%c , %v", r, size) 

//  n := utf8.RuneCountInString(str) 
//  m := utf8.RuneCountInString(str2) 

//  fmt.Println(n ,m) 

//  result := strings.Contains("I love programming" , "I") 

// fmt.Println(result) ;

// result =  strings.ContainsRune("I love orgram" ,'i')
// fmt.Println(result)


//  number :=  strings.Count("cheese" ,"e") ;
// fmt.Println(number)

    
//  fmt.Println(strings.Compare(s1 , s2))
  
//  myStr :=  strings.Repeat(str2 , 10) ;
//  fmt.Println(myStr)


//  s :=  strings.Split("hel/loxw/orl/d" , "/") 
 
//  ver :=  strings.Join(s,"/") 
//  fmt.Println(ver)

//  somstring :=  "Orange Green \n Blue Yellow" 
//  fields :=  strings.Fields(somstring) ;
//  fmt.Println(fields[0])
 
//  fmt.Println(strings.Trim("   fdf   " , " "))
// 	
vv := "amas"
zz:= "maz"
bytess := []byte(zz)
_ = vv 
fmt.Println(bytess)



fmt.Println(anagrams(vv ,zz))

}


func anagrams(s , q string) bool {


	if len(s) != len(q){
		return false
	} 
	  first  := make(map[rune]int) 
	second :=  make(map[rune]int) 
  

	 for _ , v  := range s {

 
		  _ ,ok := first[v] 
		   if ok {
 
			  first[v]++
		   }else {
			first[v] = 1 
		   }
		   


	 }

 

for _ , v  := range  q {

 
		  _ ,ok := second[v] 
		   if ok {
 
			  second[v]++
		   }else {
			second[v] = 1 
		   }
		   


	 }


	 for  _ ,v  :=range  q {


		_ ,ok:=  first[v] ;

		if !ok || second[v] != first[v]{
			return false
		}

	 }

	 return true



}