package traveltrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelTradeMemoAddAPIResponse 添加一笔交易备注 API返回值
// taobao.alitrip.travel.trade.memo.add
//
// 对一笔交易添加备注
type TaobaoAlitripTravelTradeMemoAddAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelTradeMemoAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelTradeMemoAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripTravelTradeMemoAddAPIResponseModel).Reset()
}

// TaobaoAlitripTravelTradeMemoAddAPIResponseModel is 添加一笔交易备注 成功返回结果
type TaobaoAlitripTravelTradeMemoAddAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_trade_memo_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 交易添加备注返回值
	MemoAdd *MemoCreate `json:"memo_add,omitempty" xml:"memo_add,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelTradeMemoAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MemoAdd = nil
}

var poolTaobaoAlitripTravelTradeMemoAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelTradeMemoAddAPIResponse)
	},
}

// GetTaobaoAlitripTravelTradeMemoAddAPIResponse 从 sync.Pool 获取 TaobaoAlitripTravelTradeMemoAddAPIResponse
func GetTaobaoAlitripTravelTradeMemoAddAPIResponse() *TaobaoAlitripTravelTradeMemoAddAPIResponse {
	return poolTaobaoAlitripTravelTradeMemoAddAPIResponse.Get().(*TaobaoAlitripTravelTradeMemoAddAPIResponse)
}

// ReleaseTaobaoAlitripTravelTradeMemoAddAPIResponse 将 TaobaoAlitripTravelTradeMemoAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripTravelTradeMemoAddAPIResponse(v *TaobaoAlitripTravelTradeMemoAddAPIResponse) {
	v.Reset()
	poolTaobaoAlitripTravelTradeMemoAddAPIResponse.Put(v)
}
