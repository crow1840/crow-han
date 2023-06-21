package pkg

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/toolkits/pkg/logger"
	"io"
	"os"
	"strconv"
)

func WriteFile(info string, path string) {

	fi, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0664)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fi.Close()

	_, err = fi.WriteString(info)
	if err != nil {
		return
	}

}

func EncryptWithMd5(srcString string) (result string) {
	h := md5.New()
	_, err := io.WriteString(h, srcString)
	if err != nil {
		logger.Error(err)
	}
	b := h.Sum(nil)

	result = hex.EncodeToString(b)
	return result
}

func StringToInt64(s string) int64 {
	intResult, _ := strconv.Atoi(s)
	return int64(intResult)
}
