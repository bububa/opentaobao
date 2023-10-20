package cityretail

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallCityretailWmflOrderLogisticsQueryAPIResponse 完美履约订单物流状态查询接口 API返回值
// tmall.cityretail.wmfl.order.logistics.query
//
// 完美履约订单物流状态查询接口，该接口只能查询未完结的履约单以及完结的3天内订单
type TmallCityretailWmflOrderLogisticsQueryAPIResponse struct {
	model.CommonResponse
	TmallCityretailWmflOrderLogisticsQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TmallCityretailWmflOrderLogisticsQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallCityretailWmflOrderLogisticsQueryAPIResponseModel).Reset()
}

// TmallCityretailWmflOrderLogisticsQueryAPIResponseModel is 完美履约订单物流状态查询接口 成功返回结果
type TmallCityretailWmflOrderLogisticsQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_cityretail_wmfl_order_logistics_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *WorkResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallCityretailWmflOrderLogisticsQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallCityretailWmflOrderLogisticsQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallCityretailWmflOrderLogisticsQueryAPIResponse)
	},
}

// GetTmallCityretailWmflOrderLogisticsQueryAPIResponse 从 sync.Pool 获取 TmallCityretailWmflOrderLogisticsQueryAPIResponse
func GetTmallCityretailWmflOrderLogisticsQueryAPIResponse() *TmallCityretailWmflOrderLogisticsQueryAPIResponse {
	return poolTmallCityretailWmflOrderLogisticsQueryAPIResponse.Get().(*TmallCityretailWmflOrderLogisticsQueryAPIResponse)
}

// ReleaseTmallCityretailWmflOrderLogisticsQueryAPIResponse 将 TmallCityretailWmflOrderLogisticsQueryAPIResponse 保存到 sync.Pool
func ReleaseTmallCityretailWmflOrderLogisticsQueryAPIResponse(v *TmallCityretailWmflOrderLogisticsQueryAPIResponse) {
	v.Reset()
	poolTmallCityretailWmflOrderLogisticsQueryAPIResponse.Put(v)
}
