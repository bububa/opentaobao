package openmall

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenmallTradeGetAPIResponse 查询订单详情 API返回值
// taobao.openmall.trade.get
//
// 查询订单详情
type TaobaoOpenmallTradeGetAPIResponse struct {
	model.CommonResponse
	TaobaoOpenmallTradeGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpenmallTradeGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpenmallTradeGetAPIResponseModel).Reset()
}

// TaobaoOpenmallTradeGetAPIResponseModel is 查询订单详情 成功返回结果
type TaobaoOpenmallTradeGetAPIResponseModel struct {
	XMLName xml.Name `xml:"openmall_trade_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TopTradeDetailVo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpenmallTradeGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoOpenmallTradeGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpenmallTradeGetAPIResponse)
	},
}

// GetTaobaoOpenmallTradeGetAPIResponse 从 sync.Pool 获取 TaobaoOpenmallTradeGetAPIResponse
func GetTaobaoOpenmallTradeGetAPIResponse() *TaobaoOpenmallTradeGetAPIResponse {
	return poolTaobaoOpenmallTradeGetAPIResponse.Get().(*TaobaoOpenmallTradeGetAPIResponse)
}

// ReleaseTaobaoOpenmallTradeGetAPIResponse 将 TaobaoOpenmallTradeGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpenmallTradeGetAPIResponse(v *TaobaoOpenmallTradeGetAPIResponse) {
	v.Reset()
	poolTaobaoOpenmallTradeGetAPIResponse.Put(v)
}
