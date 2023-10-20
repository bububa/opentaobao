package opentrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappAdvancedTradeinfoPriceModifyAPIResponse 高级定制商家传入改价信息 API返回值
// taobao.miniapp.advanced.tradeinfo.price.modify
//
// 高级定制商家传入改价信息
type TaobaoMiniappAdvancedTradeinfoPriceModifyAPIResponse struct {
	model.CommonResponse
	TaobaoMiniappAdvancedTradeinfoPriceModifyAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoMiniappAdvancedTradeinfoPriceModifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoMiniappAdvancedTradeinfoPriceModifyAPIResponseModel).Reset()
}

// TaobaoMiniappAdvancedTradeinfoPriceModifyAPIResponseModel is 高级定制商家传入改价信息 成功返回结果
type TaobaoMiniappAdvancedTradeinfoPriceModifyAPIResponseModel struct {
	XMLName xml.Name `xml:"miniapp_advanced_tradeinfo_price_modify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AbilityResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoMiniappAdvancedTradeinfoPriceModifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoMiniappAdvancedTradeinfoPriceModifyAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoMiniappAdvancedTradeinfoPriceModifyAPIResponse)
	},
}

// GetTaobaoMiniappAdvancedTradeinfoPriceModifyAPIResponse 从 sync.Pool 获取 TaobaoMiniappAdvancedTradeinfoPriceModifyAPIResponse
func GetTaobaoMiniappAdvancedTradeinfoPriceModifyAPIResponse() *TaobaoMiniappAdvancedTradeinfoPriceModifyAPIResponse {
	return poolTaobaoMiniappAdvancedTradeinfoPriceModifyAPIResponse.Get().(*TaobaoMiniappAdvancedTradeinfoPriceModifyAPIResponse)
}

// ReleaseTaobaoMiniappAdvancedTradeinfoPriceModifyAPIResponse 将 TaobaoMiniappAdvancedTradeinfoPriceModifyAPIResponse 保存到 sync.Pool
func ReleaseTaobaoMiniappAdvancedTradeinfoPriceModifyAPIResponse(v *TaobaoMiniappAdvancedTradeinfoPriceModifyAPIResponse) {
	v.Reset()
	poolTaobaoMiniappAdvancedTradeinfoPriceModifyAPIResponse.Put(v)
}
