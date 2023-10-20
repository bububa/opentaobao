package aliexpress

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressLogisticsAbnormalorderQueryAPIResponse 异常订单查询 API返回值
// aliexpress.logistics.abnormalorder.query
//
// 异常订单查询
type AliexpressLogisticsAbnormalorderQueryAPIResponse struct {
	model.CommonResponse
	AliexpressLogisticsAbnormalorderQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressLogisticsAbnormalorderQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressLogisticsAbnormalorderQueryAPIResponseModel).Reset()
}

// AliexpressLogisticsAbnormalorderQueryAPIResponseModel is 异常订单查询 成功返回结果
type AliexpressLogisticsAbnormalorderQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_logistics_abnormalorder_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 订单详情
	Result []AeopWarehouseAbnormalOrderResult `json:"result,omitempty" xml:"result>aeop_warehouse_abnormal_order_result,omitempty"`
	// 错误描述
	ErrorDesc string `json:"error_desc,omitempty" xml:"error_desc,omitempty"`
	// 是否成功
	IsSuccess string `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 总页数
	TotalPage int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
	// 总条数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 当前页
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressLogisticsAbnormalorderQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = m.Result[:0]
	m.ErrorDesc = ""
	m.IsSuccess = ""
	m.TotalPage = 0
	m.TotalCount = 0
	m.CurrentPage = 0
}

var poolAliexpressLogisticsAbnormalorderQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressLogisticsAbnormalorderQueryAPIResponse)
	},
}

// GetAliexpressLogisticsAbnormalorderQueryAPIResponse 从 sync.Pool 获取 AliexpressLogisticsAbnormalorderQueryAPIResponse
func GetAliexpressLogisticsAbnormalorderQueryAPIResponse() *AliexpressLogisticsAbnormalorderQueryAPIResponse {
	return poolAliexpressLogisticsAbnormalorderQueryAPIResponse.Get().(*AliexpressLogisticsAbnormalorderQueryAPIResponse)
}

// ReleaseAliexpressLogisticsAbnormalorderQueryAPIResponse 将 AliexpressLogisticsAbnormalorderQueryAPIResponse 保存到 sync.Pool
func ReleaseAliexpressLogisticsAbnormalorderQueryAPIResponse(v *AliexpressLogisticsAbnormalorderQueryAPIResponse) {
	v.Reset()
	poolAliexpressLogisticsAbnormalorderQueryAPIResponse.Put(v)
}
