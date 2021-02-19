package module
// Yaml struct of yaml
type Yaml struct {
	Mysql struct {
		User string `yaml:"user"`
		Host string `yaml:"host"`
		Password string `yaml:"password"`
		Port string `yaml:"port"`
		Name string `yaml:"name"`

	}
	LogFile struct {
		FileName string `yaml:"filename"`
	}
	Export struct {
		TableName string `yaml:"tableName"`
		ExportSQL string `yaml:"exportSQL"`
		ExportOpenPaging bool `yaml:"exportOpenPaging"`
		ExportPagingStart int `yaml:"exportPagingStart"`
		ExportPagingEnd int `yaml:"exportPagingEnd"`
		ExportPagingLimit int `yaml:"exportPagingLimit"`
	}
}
