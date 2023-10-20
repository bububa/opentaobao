package drugtrace

import (
	"sync"
)

// ResponseBody 结构体
type ResponseBody struct {
	// 流向对象
	FlowList []FlowEntity `json:"flow_list,omitempty" xml:"flow_list>flow_entity,omitempty"`
	// 首次查询时间
	FirstQueryTime string `json:"first_query_time,omitempty" xml:"first_query_time,omitempty"`
	// 最后操作时间
	LastBizDate string `json:"last_biz_date,omitempty" xml:"last_biz_date,omitempty"`
	// 药品状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 查询次数
	QueryCount string `json:"query_count,omitempty" xml:"query_count,omitempty"`
	// 商品信息、生产信息
	ProductInfoList *ProductInfoList `json:"product_info_list,omitempty" xml:"product_info_list,omitempty"`
}

var poolResponseBody = sync.Pool{
	New: func() any {
		return new(ResponseBody)
	},
}

// GetResponseBody() 从对象池中获取ResponseBody
func GetResponseBody() *ResponseBody {
	return poolResponseBody.Get().(*ResponseBody)
}

// ReleaseResponseBody 释放ResponseBody
func ReleaseResponseBody(v *ResponseBody) {
	v.FlowList = v.FlowList[:0]
	v.FirstQueryTime = ""
	v.LastBizDate = ""
	v.Status = ""
	v.QueryCount = ""
	v.ProductInfoList = nil
	poolResponseBody.Put(v)
}
