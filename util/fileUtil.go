package util

import (
	"os"
)

func Tracefile(str_content string)  {
	fd,_:=os.OpenFile(Config.LogFile.FileName,os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
	buf:=[]byte(str_content)
	fd.Write(buf)
	fd.Close()
}
