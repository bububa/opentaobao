package tbtrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTradeMemoAddAPIResponse 对一笔交易添加备注 API返回值
// taobao.trade.memo.add
//
// 根据登录用户的身份（买家或卖家），自动添加相应的交易备注,不能重复调用些接口添加备注，需要更新备注请用taobao.trade.memo.update
type TaobaoTradeMemoAddAPIResponse struct {
	model.CommonResponse
	TaobaoTradeMemoAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTradeMemoAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTradeMemoAddAPIResponseModel).Reset()
}

// TaobaoTradeMemoAddAPIResponseModel is 对一笔交易添加备注 成功返回结果
type TaobaoTradeMemoAddAPIResponseModel struct {
	XMLName xml.Name `xml:"trade_memo_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 对一笔交易添加备注后返回其对应的Trade，Trade中可用的返回字段有tid和created
	Trade *Trade `json:"trade,omitempty" xml:"trade,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTradeMemoAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Trade = nil
}

var poolTaobaoTradeMemoAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTradeMemoAddAPIResponse)
	},
}

// GetTaobaoTradeMemoAddAPIResponse 从 sync.Pool 获取 TaobaoTradeMemoAddAPIResponse
func GetTaobaoTradeMemoAddAPIResponse() *TaobaoTradeMemoAddAPIResponse {
	return poolTaobaoTradeMemoAddAPIResponse.Get().(*TaobaoTradeMemoAddAPIResponse)
}

// ReleaseTaobaoTradeMemoAddAPIResponse 将 TaobaoTradeMemoAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTradeMemoAddAPIResponse(v *TaobaoTradeMemoAddAPIResponse) {
	v.Reset()
	poolTaobaoTradeMemoAddAPIResponse.Put(v)
}
