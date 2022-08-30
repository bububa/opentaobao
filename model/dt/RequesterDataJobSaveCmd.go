package dt

// RequesterDataJobSaveCmd 结构体
type RequesterDataJobSaveCmd struct {
	// oss 导入文件路径
	OssKey string `json:"oss_key,omitempty" xml:"oss_key,omitempty"`
	// 有效值LABEL/MINE/TRAIL
	RecordType string `json:"record_type,omitempty" xml:"record_type,omitempty"`
	// 数据来源，OSS_CSV/ODPS_TABLE
	DataType string `json:"data_type,omitempty" xml:"data_type,omitempty"`
	// 操作方
	Operator string `json:"operator,omitempty" xml:"operator,omitempty"`
	// odps导入配置
	OdpsImportConfig *OdpsImportConfig `json:"odps_import_config,omitempty" xml:"odps_import_config,omitempty"`
	// 任务id
	TaskId int64 `json:"task_id,omitempty" xml:"task_id,omitempty"`
}
