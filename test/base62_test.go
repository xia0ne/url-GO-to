package test

import (
	"ginLearnDemo/utils"
	"log"
	"testing"
)

func TestBase62(t *testing.T) {
	log.Println("result", utils.Base62("baidu.com"))
}
