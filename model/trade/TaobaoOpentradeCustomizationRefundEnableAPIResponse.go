package trade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpentradeCustomizationRefundEnableAPIResponse 定制订单设置允许仅退款 API返回值
// taobao.opentrade.customization.refund.enable
//
// 定制订单设置允许仅退款
type TaobaoOpentradeCustomizationRefundEnableAPIResponse struct {
	model.CommonResponse
	TaobaoOpentradeCustomizationRefundEnableAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpentradeCustomizationRefundEnableAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpentradeCustomizationRefundEnableAPIResponseModel).Reset()
}

// TaobaoOpentradeCustomizationRefundEnableAPIResponseModel is 定制订单设置允许仅退款 成功返回结果
type TaobaoOpentradeCustomizationRefundEnableAPIResponseModel struct {
	XMLName xml.Name `xml:"opentrade_customization_refund_enable_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否设置成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpentradeCustomizationRefundEnableAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = false
}

var poolTaobaoOpentradeCustomizationRefundEnableAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpentradeCustomizationRefundEnableAPIResponse)
	},
}

// GetTaobaoOpentradeCustomizationRefundEnableAPIResponse 从 sync.Pool 获取 TaobaoOpentradeCustomizationRefundEnableAPIResponse
func GetTaobaoOpentradeCustomizationRefundEnableAPIResponse() *TaobaoOpentradeCustomizationRefundEnableAPIResponse {
	return poolTaobaoOpentradeCustomizationRefundEnableAPIResponse.Get().(*TaobaoOpentradeCustomizationRefundEnableAPIResponse)
}

// ReleaseTaobaoOpentradeCustomizationRefundEnableAPIResponse 将 TaobaoOpentradeCustomizationRefundEnableAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpentradeCustomizationRefundEnableAPIResponse(v *TaobaoOpentradeCustomizationRefundEnableAPIResponse) {
	v.Reset()
	poolTaobaoOpentradeCustomizationRefundEnableAPIResponse.Put(v)
}
