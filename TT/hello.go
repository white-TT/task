package main 
import(
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	
	
	
	
)

type WebSite struct{
	url string
	header map[string]string
}
func (keyword WebSite) get_html_header()string{
	client:=&http.Client{}
	req,err:=http.NewRequest("GET,keyword.url,nil")
	
}

func main(){
	resp,err :=http.Get("https://blog.lenconda.top/")
	if err != nil{
		fmt.Printf("get failed,err:%v\n",err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err !=nil{
		fmt.Printf("read from resp.Body failed,err:%V\n",err)
		return
	}
	fmt.Print(string(body))
}