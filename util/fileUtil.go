package util

import (
	"os"
)

func Tracefile(strContent string)  {
	fd,_:=os.OpenFile(Config.LogFile.FileName,os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
	buf:=[]byte(strContent)
	_, _ = fd.Write(buf)
	fd.Close()
}
