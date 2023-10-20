package idle

import (
	"sync"
)

// ReportUploadTopCmd 结构体
type ReportUploadTopCmd struct {
	// 唯一码，手机传imei，奢品传溯源码
	TraceCode string `json:"trace_code,omitempty" xml:"trace_code,omitempty"`
	// 产品标识，验货宝：YHB；回收：RECYCLE
	ProductCode string `json:"product_code,omitempty" xml:"product_code,omitempty"`
	// 报告内容JSON
	Report string `json:"report,omitempty" xml:"report,omitempty"`
	// 订单号
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 价格(分)
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// SPUID
	SpuId int64 `json:"spu_id,omitempty" xml:"spu_id,omitempty"`
}

var poolReportUploadTopCmd = sync.Pool{
	New: func() any {
		return new(ReportUploadTopCmd)
	},
}

// GetReportUploadTopCmd() 从对象池中获取ReportUploadTopCmd
func GetReportUploadTopCmd() *ReportUploadTopCmd {
	return poolReportUploadTopCmd.Get().(*ReportUploadTopCmd)
}

// ReleaseReportUploadTopCmd 释放ReportUploadTopCmd
func ReleaseReportUploadTopCmd(v *ReportUploadTopCmd) {
	v.TraceCode = ""
	v.ProductCode = ""
	v.Report = ""
	v.OrderId = 0
	v.Price = 0
	v.SpuId = 0
	poolReportUploadTopCmd.Put(v)
}
