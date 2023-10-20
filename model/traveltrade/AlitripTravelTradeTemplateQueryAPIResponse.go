package traveltrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripTravelTradeTemplateQueryAPIResponse 订单服务详情模版查询 API返回值
// alitrip.travel.trade.template.query
//
// 通过订单ID获取标注模版信息，商家可以根据模版来填充行业信息
type AlitripTravelTradeTemplateQueryAPIResponse struct {
	model.CommonResponse
	AlitripTravelTradeTemplateQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripTravelTradeTemplateQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripTravelTradeTemplateQueryAPIResponseModel).Reset()
}

// AlitripTravelTradeTemplateQueryAPIResponseModel is 订单服务详情模版查询 成功返回结果
type AlitripTravelTradeTemplateQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_trade_template_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 订单服务标注模版获取结果
	Result *AlitripTravelTradeTemplateQueryResultSet `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripTravelTradeTemplateQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripTravelTradeTemplateQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripTravelTradeTemplateQueryAPIResponse)
	},
}

// GetAlitripTravelTradeTemplateQueryAPIResponse 从 sync.Pool 获取 AlitripTravelTradeTemplateQueryAPIResponse
func GetAlitripTravelTradeTemplateQueryAPIResponse() *AlitripTravelTradeTemplateQueryAPIResponse {
	return poolAlitripTravelTradeTemplateQueryAPIResponse.Get().(*AlitripTravelTradeTemplateQueryAPIResponse)
}

// ReleaseAlitripTravelTradeTemplateQueryAPIResponse 将 AlitripTravelTradeTemplateQueryAPIResponse 保存到 sync.Pool
func ReleaseAlitripTravelTradeTemplateQueryAPIResponse(v *AlitripTravelTradeTemplateQueryAPIResponse) {
	v.Reset()
	poolAlitripTravelTradeTemplateQueryAPIResponse.Put(v)
}
