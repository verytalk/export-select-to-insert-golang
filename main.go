package main

import (
	"flag"
	"fmt"
	"golaer/util"
	"log"
	"strconv"
)


func foreachData(result []map[string]string, columns []string){

	fmt.Println("columns: ",columns)
	for _, r := range result {
		exportColumns := "("
		values := "("
		sql := ""
		dataLength := len(r)
		i := 0
		//sortedKeys := make([]string, 0)
		//for k, _ := range r {
		//	sortedKeys = append(sortedKeys, k)
		//}
		//
		////sort 'string' key in increasing order
		//sort.Strings(sortedKeys)
		for _, k := range columns {
			i++
			exportColumns += k
			values += "'"+ util.ReplaceSQLStr(r[k])+"'"
			if dataLength != i {
				exportColumns += ","
				values += ","
			}
		}
		exportColumns += ")"
		values += ")"
		exportTable:= util.Config.Export.TableName
		sql = "INSERT "+exportTable+" " + exportColumns + " VALUES " + values + "; \n"
		//fmt.Print(sql)
		util.Tracefile(sql)
	}

}


func getData(exportSQL string) int{
	fmt.Println("exportSQL: ",exportSQL)
	var result, columns = util.RunSelect(exportSQL)
	fmt.Println("columns: ",columns)
	resultLength := len(result)
	fmt.Println("dataLength: ", resultLength)
	foreachData(result,columns)
	return resultLength
}

func task() {
	log.Println("running ...")
	//fmt.Println(util.Config.Crontab.Period)
	//util.Run()
	exportSQL:= util.Config.Export.ExportSQL
	exportOpenPaging:= util.Config.Export.ExportOpenPaging

	if exportOpenPaging {
		exportPagingEnd:= util.Config.Export.ExportPagingEnd
		exportPagingStart:= util.Config.Export.ExportPagingStart
		exportPagingLimit:= util.Config.Export.ExportPagingLimit

		for i := exportPagingStart; i <= exportPagingEnd; i++ {
			limitStart := i * exportPagingLimit
			fmt.Println("current page: ",i,"page limit: ",exportPagingLimit)
			exportSQLPaging := exportSQL+" LIMIT " + strconv.Itoa(limitStart) + ","+ strconv.Itoa(exportPagingLimit)
			dataLength := getData(exportSQLPaging)
			if 0 == dataLength {
				fmt.Println("export complete , the page is:",i,"data length is :",dataLength)
				break
			}
		}
	}else{
		getData(exportSQL)
	}
	log.Println("export done please check logfile " + util.Config.LogFile.FileName)
}



func main() {
	config := flag.String("config", "./config.yaml", "the config")
	flag.Parse()
	fmt.Println("config:", *config)
	util.ConfigFile = *config
	util.InitConfig()
	task()
}
