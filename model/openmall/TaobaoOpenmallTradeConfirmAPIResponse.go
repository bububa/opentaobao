package openmall

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenmalltradeconfirmAPIResponse 确认收货 API返回值
// taobao.openmall.trade.confirm
//
// 确认订单收货
type TaobaoopenmalltradeconfirmAPIResponse struct {
	model.CommonResponse
	TaobaoopenmalltradeconfirmAPIResponseModel
}

// TaobaoopenmalltradeconfirmAPIResponseModel is 确认收货 成功返回结果
type TaobaoopenmalltradeconfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"openmall_trade_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TopTradeResultVo `json:"result,omitempty" xml:"result,omitempty"`
}
