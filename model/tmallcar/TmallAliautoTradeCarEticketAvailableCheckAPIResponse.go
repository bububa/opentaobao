package tmallcar

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallAliautoTradeCarEticketAvailableCheckAPIResponse 天猫汽车电子凭证核销前校验 API返回值
// tmall.aliauto.trade.car.eticket.available.check
//
// 天猫汽车核销码可用性校验
type TmallAliautoTradeCarEticketAvailableCheckAPIResponse struct {
	model.CommonResponse
	TmallAliautoTradeCarEticketAvailableCheckAPIResponseModel
}

// TmallAliautoTradeCarEticketAvailableCheckAPIResponseModel is 天猫汽车电子凭证核销前校验 成功返回结果
type TmallAliautoTradeCarEticketAvailableCheckAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_aliauto_trade_car_eticket_available_check_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *AliAutoResult `json:"result,omitempty" xml:"result,omitempty"`
}
