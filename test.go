package main

	import(
		"fmt"
		"net/http"
		"bytes"
		"encoding/json"
	)

	func main(){

		resp,err:=http.Get("http://api.3vbook.com/book/index/content?bookId=5509&chapterId=214")
		if err !=nil{
			return
		}

		defer resp.Body.Close()

		//headers:=resp.Header

		//for k,v:=range headers{
		//	fmt.Println(k,v)
		//}

		//fmt.Printf("resp status %s,statusCode %d\n", resp.Status, resp.StatusCode)

		//fmt.Println(resp)

		buf := bytes.NewBuffer(make([]byte, 0, 512))
		buf.ReadFrom(resp.Body)

		contents := string(buf.Bytes())
		fmt.Println(contents)
		//var jsonBlob = content


		type Animal struct {  
		    id  string  
		    book_id string  
		    chapter_id string
		    content string
		    title string
		}  

		aa := Animal{}

		err0 := json.Unmarshal([]byte(contents), &aa)
		
		fmt.Println(aa, err0)

		//fmt.Println(content)
	}
