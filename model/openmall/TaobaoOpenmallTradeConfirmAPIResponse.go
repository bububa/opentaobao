package openmall

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenmallTradeConfirmAPIResponse 确认收货 API返回值
// taobao.openmall.trade.confirm
//
// 确认订单收货
type TaobaoOpenmallTradeConfirmAPIResponse struct {
	model.CommonResponse
	TaobaoOpenmallTradeConfirmAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpenmallTradeConfirmAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpenmallTradeConfirmAPIResponseModel).Reset()
}

// TaobaoOpenmallTradeConfirmAPIResponseModel is 确认收货 成功返回结果
type TaobaoOpenmallTradeConfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"openmall_trade_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TopTradeResultVo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpenmallTradeConfirmAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoOpenmallTradeConfirmAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpenmallTradeConfirmAPIResponse)
	},
}

// GetTaobaoOpenmallTradeConfirmAPIResponse 从 sync.Pool 获取 TaobaoOpenmallTradeConfirmAPIResponse
func GetTaobaoOpenmallTradeConfirmAPIResponse() *TaobaoOpenmallTradeConfirmAPIResponse {
	return poolTaobaoOpenmallTradeConfirmAPIResponse.Get().(*TaobaoOpenmallTradeConfirmAPIResponse)
}

// ReleaseTaobaoOpenmallTradeConfirmAPIResponse 将 TaobaoOpenmallTradeConfirmAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpenmallTradeConfirmAPIResponse(v *TaobaoOpenmallTradeConfirmAPIResponse) {
	v.Reset()
	poolTaobaoOpenmallTradeConfirmAPIResponse.Put(v)
}
