package xhoteloffline

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelOrderAlipayfaceExtendAPIResponse 信用住订单延住接口 API返回值
// taobao.xhotel.order.alipayface.extend
//
// 信用住订单延住接口，用于将已有的信用住支付订单checkOut和订单金额等更新
type TaobaoXhotelOrderAlipayfaceExtendAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelOrderAlipayfaceExtendAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelOrderAlipayfaceExtendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelOrderAlipayfaceExtendAPIResponseModel).Reset()
}

// TaobaoXhotelOrderAlipayfaceExtendAPIResponseModel is 信用住订单延住接口 成功返回结果
type TaobaoXhotelOrderAlipayfaceExtendAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_order_alipayface_extend_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参成功
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelOrderAlipayfaceExtendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolTaobaoXhotelOrderAlipayfaceExtendAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelOrderAlipayfaceExtendAPIResponse)
	},
}

// GetTaobaoXhotelOrderAlipayfaceExtendAPIResponse 从 sync.Pool 获取 TaobaoXhotelOrderAlipayfaceExtendAPIResponse
func GetTaobaoXhotelOrderAlipayfaceExtendAPIResponse() *TaobaoXhotelOrderAlipayfaceExtendAPIResponse {
	return poolTaobaoXhotelOrderAlipayfaceExtendAPIResponse.Get().(*TaobaoXhotelOrderAlipayfaceExtendAPIResponse)
}

// ReleaseTaobaoXhotelOrderAlipayfaceExtendAPIResponse 将 TaobaoXhotelOrderAlipayfaceExtendAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelOrderAlipayfaceExtendAPIResponse(v *TaobaoXhotelOrderAlipayfaceExtendAPIResponse) {
	v.Reset()
	poolTaobaoXhotelOrderAlipayfaceExtendAPIResponse.Put(v)
}
