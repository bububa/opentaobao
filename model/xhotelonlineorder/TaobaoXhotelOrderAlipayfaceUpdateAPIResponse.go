package xhotelonlineorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelOrderAlipayfaceUpdateAPIResponse 酒店信用住订单状态更新 API返回值
// taobao.xhotel.order.alipayface.update
//
// 完成对信用住或者面付订单的状态的更新。包含订单状态的确认，入离店状态的更新等等。（不适用于预付订单）
type TaobaoXhotelOrderAlipayfaceUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelOrderAlipayfaceUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelOrderAlipayfaceUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelOrderAlipayfaceUpdateAPIResponseModel).Reset()
}

// TaobaoXhotelOrderAlipayfaceUpdateAPIResponseModel is 酒店信用住订单状态更新 成功返回结果
type TaobaoXhotelOrderAlipayfaceUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_order_alipayface_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回提示信息
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelOrderAlipayfaceUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolTaobaoXhotelOrderAlipayfaceUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelOrderAlipayfaceUpdateAPIResponse)
	},
}

// GetTaobaoXhotelOrderAlipayfaceUpdateAPIResponse 从 sync.Pool 获取 TaobaoXhotelOrderAlipayfaceUpdateAPIResponse
func GetTaobaoXhotelOrderAlipayfaceUpdateAPIResponse() *TaobaoXhotelOrderAlipayfaceUpdateAPIResponse {
	return poolTaobaoXhotelOrderAlipayfaceUpdateAPIResponse.Get().(*TaobaoXhotelOrderAlipayfaceUpdateAPIResponse)
}

// ReleaseTaobaoXhotelOrderAlipayfaceUpdateAPIResponse 将 TaobaoXhotelOrderAlipayfaceUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelOrderAlipayfaceUpdateAPIResponse(v *TaobaoXhotelOrderAlipayfaceUpdateAPIResponse) {
	v.Reset()
	poolTaobaoXhotelOrderAlipayfaceUpdateAPIResponse.Put(v)
}
