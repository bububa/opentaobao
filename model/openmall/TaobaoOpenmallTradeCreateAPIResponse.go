package openmall

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenmallTradeCreateAPIResponse 创建订单 API返回值
// taobao.openmall.trade.create
//
// 创建Openmall订单
type TaobaoOpenmallTradeCreateAPIResponse struct {
	model.CommonResponse
	TaobaoOpenmallTradeCreateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpenmallTradeCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpenmallTradeCreateAPIResponseModel).Reset()
}

// TaobaoOpenmallTradeCreateAPIResponseModel is 创建订单 成功返回结果
type TaobaoOpenmallTradeCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"openmall_trade_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TopTradeResultVo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpenmallTradeCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoOpenmallTradeCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpenmallTradeCreateAPIResponse)
	},
}

// GetTaobaoOpenmallTradeCreateAPIResponse 从 sync.Pool 获取 TaobaoOpenmallTradeCreateAPIResponse
func GetTaobaoOpenmallTradeCreateAPIResponse() *TaobaoOpenmallTradeCreateAPIResponse {
	return poolTaobaoOpenmallTradeCreateAPIResponse.Get().(*TaobaoOpenmallTradeCreateAPIResponse)
}

// ReleaseTaobaoOpenmallTradeCreateAPIResponse 将 TaobaoOpenmallTradeCreateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpenmallTradeCreateAPIResponse(v *TaobaoOpenmallTradeCreateAPIResponse) {
	v.Reset()
	poolTaobaoOpenmallTradeCreateAPIResponse.Put(v)
}
