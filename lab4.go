package main
import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
  data string
  position int
}

type Token struct {
	data string
	recipient int 
	ttl int 
}

func makePair(i int, N int) Token {
    var token Token
    token.recipient = rand.Intn(N)
	// для уменьшения шанса таймаута
    token.ttl = (rand.Intn(N)+(N/2))
    token.data = fmt.Sprintf("data for node ¹%d",token.recipient)
    return token
}

func show(){
  fmt.Print("\n")
  for i := range nodes {   
    if (nodes[i].data == "") {
        fmt.Print(i,"'th node: the message did not recieved\n")
        continue;
    }
        fmt.Print(i,"'th node: ",nodes[i].data,"\n")
    }
}

func drop (chanel chan Token, i int){
	token:= <- chanel
	for j := range nodes{
		if (token.recipient == nodes[j].position) {
			nodes[j].data = token.data
		}else{
			token.ttl--
			if(token.ttl<=0){
				break
			}
		}
	}
	defer close(chanel)
}

var nodes []Node;

func main() {
	var N int
	rand.Seed(time.Now().UnixNano())
	fmt.Scanf("%d", &N)
	chans := make([]chan Token, N)
	nodes = make([]Node, N)
	
	for i := range nodes {                
        nodes[i].position = i
	}
	
	for i := range chans {                
		chans[i] = make(chan Token,1) 
		newChan := makePair(i,N)
		chans[i]<- newChan
		go drop(chans[i],i)
	}
	
	show()
	time.Sleep(1 * time.Second)
}