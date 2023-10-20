package openmall

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenmallTradeRenderAPIResponse 渲染订单价格 API返回值
// taobao.openmall.trade.render
//
// 请求渲染订单价格
type TaobaoOpenmallTradeRenderAPIResponse struct {
	model.CommonResponse
	TaobaoOpenmallTradeRenderAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpenmallTradeRenderAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpenmallTradeRenderAPIResponseModel).Reset()
}

// TaobaoOpenmallTradeRenderAPIResponseModel is 渲染订单价格 成功返回结果
type TaobaoOpenmallTradeRenderAPIResponseModel struct {
	XMLName xml.Name `xml:"openmall_trade_render_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TopTradeResultVo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpenmallTradeRenderAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoOpenmallTradeRenderAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpenmallTradeRenderAPIResponse)
	},
}

// GetTaobaoOpenmallTradeRenderAPIResponse 从 sync.Pool 获取 TaobaoOpenmallTradeRenderAPIResponse
func GetTaobaoOpenmallTradeRenderAPIResponse() *TaobaoOpenmallTradeRenderAPIResponse {
	return poolTaobaoOpenmallTradeRenderAPIResponse.Get().(*TaobaoOpenmallTradeRenderAPIResponse)
}

// ReleaseTaobaoOpenmallTradeRenderAPIResponse 将 TaobaoOpenmallTradeRenderAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpenmallTradeRenderAPIResponse(v *TaobaoOpenmallTradeRenderAPIResponse) {
	v.Reset()
	poolTaobaoOpenmallTradeRenderAPIResponse.Put(v)
}
