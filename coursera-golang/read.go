package main

import ("os")

func f (){

f, err := os.Open("text.txt")
sli := make([]byte,20)
nb, err := f.Read(sli)

	type names struct{
		fname string
		lname string
	} 

	func (f *File) Readdir(n, int) ([]FileInfo, error) {
		if f == nil {
			return nil, ErrInvalid
		}
		return f.readdir(n)
	}

	func (f *File) Readdirnames(n int) (names []string, err error){
		if f == nil{
			return nil, ErrInvalid
		}
		return f.readdirnames(n)
	}


