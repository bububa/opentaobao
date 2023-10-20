package tmallcar

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallAliautoFulfillmentDeliverySynAPIResponse 交付状态及物流信息同步 API返回值
// tmall.aliauto.fulfillment.delivery.syn
//
// 交付状态及物流信息同步
type TmallAliautoFulfillmentDeliverySynAPIResponse struct {
	model.CommonResponse
	TmallAliautoFulfillmentDeliverySynAPIResponseModel
}

// Reset 清空结构体
func (m *TmallAliautoFulfillmentDeliverySynAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallAliautoFulfillmentDeliverySynAPIResponseModel).Reset()
}

// TmallAliautoFulfillmentDeliverySynAPIResponseModel is 交付状态及物流信息同步 成功返回结果
type TmallAliautoFulfillmentDeliverySynAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_aliauto_fulfillment_delivery_syn_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *AliAutoResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallAliautoFulfillmentDeliverySynAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallAliautoFulfillmentDeliverySynAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallAliautoFulfillmentDeliverySynAPIResponse)
	},
}

// GetTmallAliautoFulfillmentDeliverySynAPIResponse 从 sync.Pool 获取 TmallAliautoFulfillmentDeliverySynAPIResponse
func GetTmallAliautoFulfillmentDeliverySynAPIResponse() *TmallAliautoFulfillmentDeliverySynAPIResponse {
	return poolTmallAliautoFulfillmentDeliverySynAPIResponse.Get().(*TmallAliautoFulfillmentDeliverySynAPIResponse)
}

// ReleaseTmallAliautoFulfillmentDeliverySynAPIResponse 将 TmallAliautoFulfillmentDeliverySynAPIResponse 保存到 sync.Pool
func ReleaseTmallAliautoFulfillmentDeliverySynAPIResponse(v *TmallAliautoFulfillmentDeliverySynAPIResponse) {
	v.Reset()
	poolTmallAliautoFulfillmentDeliverySynAPIResponse.Put(v)
}
