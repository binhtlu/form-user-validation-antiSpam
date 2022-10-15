package User

import (
	"bufio"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

//var counter HandleIpRequest.Counter

func AddNewUser(c *gin.Context) {
	//var counter HandleIpRequest.Counter neu khai bao scyn.map o day thi moi lan request thi se tao ra mot map moi

	var NewUser user
	if err := c.ShouldBindJSON(&NewUser); err != nil {
		//var tx *redis.Tx
		//xForwardedAddress := c.Request.Header.Get("X-Forwarded-For")
		//clientIp := c.ClientIP()
		//tx.ClientKill(xForwardedAddress)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	//luu vao mang
	UserList = append(UserList, NewUser)
	c.JSON(http.StatusOK, NewUser)
	// luu xuong file txt
	SaveUser(c, NewUser)
}

func SaveUser(c *gin.Context, NewUser user) {
	sampledata := []string{fmt.Sprintf("%v", NewUser)}
	file, err := os.OpenFile("save-user/file-save.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	datawriter := bufio.NewWriter(file)
	for _, data := range sampledata {
		_, _ = datawriter.WriteString(data + "\n")
	}

	datawriter.Flush()
	file.Close()
}
