package trade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoservindustryfinancegeexorderupdateAPIResponse 即科订单结果更新回调 API返回值
// taobao.servindustry.finance.geex.order.update
//
// 即科订单结果更新回调本地接口
type TaobaoservindustryfinancegeexorderupdateAPIResponse struct {
	model.CommonResponse
	TaobaoservindustryfinancegeexorderupdateAPIResponseModel
}

// TaobaoservindustryfinancegeexorderupdateAPIResponseModel is 即科订单结果更新回调 成功返回结果
type TaobaoservindustryfinancegeexorderupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"servindustry_finance_geex_order_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *RetryResult `json:"result,omitempty" xml:"result,omitempty"`
}
