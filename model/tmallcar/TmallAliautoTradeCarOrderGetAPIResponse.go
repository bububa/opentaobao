package tmallcar

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallAliautoTradeCarOrderGetAPIResponse 整车订单详情查询 API返回值
// tmall.aliauto.trade.car.order.get
//
// 整车订单详情查询接口
type TmallAliautoTradeCarOrderGetAPIResponse struct {
	model.CommonResponse
	TmallAliautoTradeCarOrderGetAPIResponseModel
}

// Reset 清空结构体
func (m *TmallAliautoTradeCarOrderGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallAliautoTradeCarOrderGetAPIResponseModel).Reset()
}

// TmallAliautoTradeCarOrderGetAPIResponseModel is 整车订单详情查询 成功返回结果
type TmallAliautoTradeCarOrderGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_aliauto_trade_car_order_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *AliAutoResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallAliautoTradeCarOrderGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallAliautoTradeCarOrderGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallAliautoTradeCarOrderGetAPIResponse)
	},
}

// GetTmallAliautoTradeCarOrderGetAPIResponse 从 sync.Pool 获取 TmallAliautoTradeCarOrderGetAPIResponse
func GetTmallAliautoTradeCarOrderGetAPIResponse() *TmallAliautoTradeCarOrderGetAPIResponse {
	return poolTmallAliautoTradeCarOrderGetAPIResponse.Get().(*TmallAliautoTradeCarOrderGetAPIResponse)
}

// ReleaseTmallAliautoTradeCarOrderGetAPIResponse 将 TmallAliautoTradeCarOrderGetAPIResponse 保存到 sync.Pool
func ReleaseTmallAliautoTradeCarOrderGetAPIResponse(v *TmallAliautoTradeCarOrderGetAPIResponse) {
	v.Reset()
	poolTmallAliautoTradeCarOrderGetAPIResponse.Put(v)
}
