package traveltrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelTradeMemoUpdateAPIResponse 修改一笔交易备注 API返回值
// taobao.alitrip.travel.trade.memo.update
//
// 更新一笔交易备注
type TaobaoAlitripTravelTradeMemoUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelTradeMemoUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelTradeMemoUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripTravelTradeMemoUpdateAPIResponseModel).Reset()
}

// TaobaoAlitripTravelTradeMemoUpdateAPIResponseModel is 修改一笔交易备注 成功返回结果
type TaobaoAlitripTravelTradeMemoUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_trade_memo_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 修改交易备注返回值
	MemoUpdate *MemoUpdate `json:"memo_update,omitempty" xml:"memo_update,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelTradeMemoUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MemoUpdate = nil
}

var poolTaobaoAlitripTravelTradeMemoUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelTradeMemoUpdateAPIResponse)
	},
}

// GetTaobaoAlitripTravelTradeMemoUpdateAPIResponse 从 sync.Pool 获取 TaobaoAlitripTravelTradeMemoUpdateAPIResponse
func GetTaobaoAlitripTravelTradeMemoUpdateAPIResponse() *TaobaoAlitripTravelTradeMemoUpdateAPIResponse {
	return poolTaobaoAlitripTravelTradeMemoUpdateAPIResponse.Get().(*TaobaoAlitripTravelTradeMemoUpdateAPIResponse)
}

// ReleaseTaobaoAlitripTravelTradeMemoUpdateAPIResponse 将 TaobaoAlitripTravelTradeMemoUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripTravelTradeMemoUpdateAPIResponse(v *TaobaoAlitripTravelTradeMemoUpdateAPIResponse) {
	v.Reset()
	poolTaobaoAlitripTravelTradeMemoUpdateAPIResponse.Put(v)
}
