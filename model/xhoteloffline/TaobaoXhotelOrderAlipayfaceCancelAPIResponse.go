package xhoteloffline

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelOrderAlipayfaceCancelAPIResponse 线下信用住订单取消 API返回值
// taobao.xhotel.order.alipayface.cancel
//
// 提供给卖家进行线下信用住的订单取消。此接口仅仅支持线下信用住订单的取消
type TaobaoXhotelOrderAlipayfaceCancelAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelOrderAlipayfaceCancelAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelOrderAlipayfaceCancelAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelOrderAlipayfaceCancelAPIResponseModel).Reset()
}

// TaobaoXhotelOrderAlipayfaceCancelAPIResponseModel is 线下信用住订单取消 成功返回结果
type TaobaoXhotelOrderAlipayfaceCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_order_alipayface_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回描述
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelOrderAlipayfaceCancelAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolTaobaoXhotelOrderAlipayfaceCancelAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelOrderAlipayfaceCancelAPIResponse)
	},
}

// GetTaobaoXhotelOrderAlipayfaceCancelAPIResponse 从 sync.Pool 获取 TaobaoXhotelOrderAlipayfaceCancelAPIResponse
func GetTaobaoXhotelOrderAlipayfaceCancelAPIResponse() *TaobaoXhotelOrderAlipayfaceCancelAPIResponse {
	return poolTaobaoXhotelOrderAlipayfaceCancelAPIResponse.Get().(*TaobaoXhotelOrderAlipayfaceCancelAPIResponse)
}

// ReleaseTaobaoXhotelOrderAlipayfaceCancelAPIResponse 将 TaobaoXhotelOrderAlipayfaceCancelAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelOrderAlipayfaceCancelAPIResponse(v *TaobaoXhotelOrderAlipayfaceCancelAPIResponse) {
	v.Reset()
	poolTaobaoXhotelOrderAlipayfaceCancelAPIResponse.Put(v)
}
