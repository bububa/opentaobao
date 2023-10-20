package alihealth2

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaotradedrugordergetAPIResponse 查看订单详情 API返回值
// taobao.trade.drug.order.get
//
// 商家查看订单详情
type TaobaotradedrugordergetAPIResponse struct {
	model.CommonResponse
	TaobaotradedrugordergetAPIResponseModel
}

// TaobaotradedrugordergetAPIResponseModel is 查看订单详情 成功返回结果
type TaobaotradedrugordergetAPIResponseModel struct {
	XMLName xml.Name `xml:"trade_drug_order_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TaobaotradedrugordergetResultSet `json:"result,omitempty" xml:"result,omitempty"`
}
