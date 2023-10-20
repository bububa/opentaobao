package miniapp

import (
	"sync"
)

// AfterSaleGetWorkTableListRecord 结构体
type AfterSaleGetWorkTableListRecord struct {
	// 工作表id
	TableId string `json:"table_id,omitempty" xml:"table_id,omitempty"`
	// 工作表名称
	TableName string `json:"table_name,omitempty" xml:"table_name,omitempty"`
	// 工作表创建时间
	CreateTime string `json:"create_time,omitempty" xml:"create_time,omitempty"`
}

var poolAfterSaleGetWorkTableListRecord = sync.Pool{
	New: func() any {
		return new(AfterSaleGetWorkTableListRecord)
	},
}

// GetAfterSaleGetWorkTableListRecord() 从对象池中获取AfterSaleGetWorkTableListRecord
func GetAfterSaleGetWorkTableListRecord() *AfterSaleGetWorkTableListRecord {
	return poolAfterSaleGetWorkTableListRecord.Get().(*AfterSaleGetWorkTableListRecord)
}

// ReleaseAfterSaleGetWorkTableListRecord 释放AfterSaleGetWorkTableListRecord
func ReleaseAfterSaleGetWorkTableListRecord(v *AfterSaleGetWorkTableListRecord) {
	v.TableId = ""
	v.TableName = ""
	v.CreateTime = ""
	poolAfterSaleGetWorkTableListRecord.Put(v)
}
