package xhotelonlineorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelOrderAlipayfaceSettleAPIResponse 信用住订单结账接口 API返回值
// taobao.xhotel.order.alipayface.settle
//
// 用于离店付订单在客人离店后，发起结账以及扣款等后续动作
type TaobaoXhotelOrderAlipayfaceSettleAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelOrderAlipayfaceSettleAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelOrderAlipayfaceSettleAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelOrderAlipayfaceSettleAPIResponseModel).Reset()
}

// TaobaoXhotelOrderAlipayfaceSettleAPIResponseModel is 信用住订单结账接口 成功返回结果
type TaobaoXhotelOrderAlipayfaceSettleAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_order_alipayface_settle_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelOrderAlipayfaceSettleAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolTaobaoXhotelOrderAlipayfaceSettleAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelOrderAlipayfaceSettleAPIResponse)
	},
}

// GetTaobaoXhotelOrderAlipayfaceSettleAPIResponse 从 sync.Pool 获取 TaobaoXhotelOrderAlipayfaceSettleAPIResponse
func GetTaobaoXhotelOrderAlipayfaceSettleAPIResponse() *TaobaoXhotelOrderAlipayfaceSettleAPIResponse {
	return poolTaobaoXhotelOrderAlipayfaceSettleAPIResponse.Get().(*TaobaoXhotelOrderAlipayfaceSettleAPIResponse)
}

// ReleaseTaobaoXhotelOrderAlipayfaceSettleAPIResponse 将 TaobaoXhotelOrderAlipayfaceSettleAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelOrderAlipayfaceSettleAPIResponse(v *TaobaoXhotelOrderAlipayfaceSettleAPIResponse) {
	v.Reset()
	poolTaobaoXhotelOrderAlipayfaceSettleAPIResponse.Put(v)
}
