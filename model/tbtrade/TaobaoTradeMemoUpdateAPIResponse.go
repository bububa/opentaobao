package tbtrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTradeMemoUpdateAPIResponse 修改交易备注 API返回值
// taobao.trade.memo.update
//
// 需要商家或以上权限才可调用此接口，可重复调用本接口更新交易备注，本接口同时具有添加备注的功能
type TaobaoTradeMemoUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoTradeMemoUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTradeMemoUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTradeMemoUpdateAPIResponseModel).Reset()
}

// TaobaoTradeMemoUpdateAPIResponseModel is 修改交易备注 成功返回结果
type TaobaoTradeMemoUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"trade_memo_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 更新交易的备注信息后返回的Trade，其中可用字段为tid和modified
	Trade *Trade `json:"trade,omitempty" xml:"trade,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTradeMemoUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Trade = nil
}

var poolTaobaoTradeMemoUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTradeMemoUpdateAPIResponse)
	},
}

// GetTaobaoTradeMemoUpdateAPIResponse 从 sync.Pool 获取 TaobaoTradeMemoUpdateAPIResponse
func GetTaobaoTradeMemoUpdateAPIResponse() *TaobaoTradeMemoUpdateAPIResponse {
	return poolTaobaoTradeMemoUpdateAPIResponse.Get().(*TaobaoTradeMemoUpdateAPIResponse)
}

// ReleaseTaobaoTradeMemoUpdateAPIResponse 将 TaobaoTradeMemoUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTradeMemoUpdateAPIResponse(v *TaobaoTradeMemoUpdateAPIResponse) {
	v.Reset()
	poolTaobaoTradeMemoUpdateAPIResponse.Put(v)
}
