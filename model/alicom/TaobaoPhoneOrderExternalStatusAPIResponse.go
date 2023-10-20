package alicom

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPhoneOrderExternalStatusAPIResponse 话费外放订单状态接口 API返回值
// taobao.phone.order.external.status
//
// 话费外放订单状态接口
type TaobaoPhoneOrderExternalStatusAPIResponse struct {
	model.CommonResponse
	TaobaoPhoneOrderExternalStatusAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPhoneOrderExternalStatusAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPhoneOrderExternalStatusAPIResponseModel).Reset()
}

// TaobaoPhoneOrderExternalStatusAPIResponseModel is 话费外放订单状态接口 成功返回结果
type TaobaoPhoneOrderExternalStatusAPIResponseModel struct {
	XMLName xml.Name `xml:"phone_order_external_status_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应编码
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 响应描述
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 返回结果
	Model *DistributeTradeOrderInfo `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPhoneOrderExternalStatusAPIResponseModel) Reset() {
	m.RequestId = ""
	m.BizCode = ""
	m.Desc = ""
	m.Model = nil
	m.IsSuccess = false
}

var poolTaobaoPhoneOrderExternalStatusAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPhoneOrderExternalStatusAPIResponse)
	},
}

// GetTaobaoPhoneOrderExternalStatusAPIResponse 从 sync.Pool 获取 TaobaoPhoneOrderExternalStatusAPIResponse
func GetTaobaoPhoneOrderExternalStatusAPIResponse() *TaobaoPhoneOrderExternalStatusAPIResponse {
	return poolTaobaoPhoneOrderExternalStatusAPIResponse.Get().(*TaobaoPhoneOrderExternalStatusAPIResponse)
}

// ReleaseTaobaoPhoneOrderExternalStatusAPIResponse 将 TaobaoPhoneOrderExternalStatusAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPhoneOrderExternalStatusAPIResponse(v *TaobaoPhoneOrderExternalStatusAPIResponse) {
	v.Reset()
	poolTaobaoPhoneOrderExternalStatusAPIResponse.Put(v)
}
