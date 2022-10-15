package tests

import (
	"controllers/HandleIpRequest"
	"fmt"
	"testing"
)

func TestQP_AnswerAfterSelectJoinWithTokenCreate(t *testing.T) {
	var counter HandleIpRequest.Counter
	counter.Add("binh", 2)
	counter.Add("binh", 2)

	fmt.Println(counter.Get("binh"))

	//fmt.Println(counter.DeleteAndGetLastValue("key"))

	//fmt.Println(counter.Get("key"))
}
