package miniapp

// AfterSaleGetWorkTableListRecord 结构体
type AfterSaleGetWorkTableListRecord struct {
	// 工作表id
	TableId string `json:"table_id,omitempty" xml:"table_id,omitempty"`
	// 工作表名称
	TableName string `json:"table_name,omitempty" xml:"table_name,omitempty"`
	// 工作表创建时间
	CreateTime string `json:"create_time,omitempty" xml:"create_time,omitempty"`
}
