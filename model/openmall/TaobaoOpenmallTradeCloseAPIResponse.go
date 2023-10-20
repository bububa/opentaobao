package openmall

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenmallTradeCloseAPIResponse 关闭订单 API返回值
// taobao.openmall.trade.close
//
// 关闭订单
type TaobaoOpenmallTradeCloseAPIResponse struct {
	model.CommonResponse
	TaobaoOpenmallTradeCloseAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpenmallTradeCloseAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpenmallTradeCloseAPIResponseModel).Reset()
}

// TaobaoOpenmallTradeCloseAPIResponseModel is 关闭订单 成功返回结果
type TaobaoOpenmallTradeCloseAPIResponseModel struct {
	XMLName xml.Name `xml:"openmall_trade_close_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *TopTradeResultVo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpenmallTradeCloseAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoOpenmallTradeCloseAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpenmallTradeCloseAPIResponse)
	},
}

// GetTaobaoOpenmallTradeCloseAPIResponse 从 sync.Pool 获取 TaobaoOpenmallTradeCloseAPIResponse
func GetTaobaoOpenmallTradeCloseAPIResponse() *TaobaoOpenmallTradeCloseAPIResponse {
	return poolTaobaoOpenmallTradeCloseAPIResponse.Get().(*TaobaoOpenmallTradeCloseAPIResponse)
}

// ReleaseTaobaoOpenmallTradeCloseAPIResponse 将 TaobaoOpenmallTradeCloseAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpenmallTradeCloseAPIResponse(v *TaobaoOpenmallTradeCloseAPIResponse) {
	v.Reset()
	poolTaobaoOpenmallTradeCloseAPIResponse.Put(v)
}
