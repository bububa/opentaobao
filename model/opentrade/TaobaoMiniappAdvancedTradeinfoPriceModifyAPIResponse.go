package opentrade

import (
	"encoding/xml"

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

// TaobaoMiniappAdvancedTradeinfoPriceModifyAPIResponseModel is 高级定制商家传入改价信息 成功返回结果
type TaobaoMiniappAdvancedTradeinfoPriceModifyAPIResponseModel struct {
	XMLName xml.Name `xml:"miniapp_advanced_tradeinfo_price_modify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AbilityResponse `json:"result,omitempty" xml:"result,omitempty"`
}
