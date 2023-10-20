package alicom

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWtTradeOrderResultcallbackAPIResponse 商家回调接口 API返回值
// taobao.wt.trade.order.resultcallback
//
// 阿里通信定制服务，商家发货后进行调用该接口，用于自动发货并确认收货
type TaobaoWtTradeOrderResultcallbackAPIResponse struct {
	model.CommonResponse
	TaobaoWtTradeOrderResultcallbackAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWtTradeOrderResultcallbackAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWtTradeOrderResultcallbackAPIResponseModel).Reset()
}

// TaobaoWtTradeOrderResultcallbackAPIResponseModel is 商家回调接口 成功返回结果
type TaobaoWtTradeOrderResultcallbackAPIResponseModel struct {
	XMLName xml.Name `xml:"wt_trade_order_resultcallback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *CommonRtnDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWtTradeOrderResultcallbackAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoWtTradeOrderResultcallbackAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWtTradeOrderResultcallbackAPIResponse)
	},
}

// GetTaobaoWtTradeOrderResultcallbackAPIResponse 从 sync.Pool 获取 TaobaoWtTradeOrderResultcallbackAPIResponse
func GetTaobaoWtTradeOrderResultcallbackAPIResponse() *TaobaoWtTradeOrderResultcallbackAPIResponse {
	return poolTaobaoWtTradeOrderResultcallbackAPIResponse.Get().(*TaobaoWtTradeOrderResultcallbackAPIResponse)
}

// ReleaseTaobaoWtTradeOrderResultcallbackAPIResponse 将 TaobaoWtTradeOrderResultcallbackAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWtTradeOrderResultcallbackAPIResponse(v *TaobaoWtTradeOrderResultcallbackAPIResponse) {
	v.Reset()
	poolTaobaoWtTradeOrderResultcallbackAPIResponse.Put(v)
}
