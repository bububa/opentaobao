package tmallcar

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallaliautotradecareticketconsumeAPIResponse 天猫汽车电子凭证核销 API返回值
// tmall.aliauto.trade.car.eticket.consume
//
// 为商家提供电子凭证核销接口，支持分账
type TmallaliautotradecareticketconsumeAPIResponse struct {
	model.CommonResponse
	TmallaliautotradecareticketconsumeAPIResponseModel
}

// TmallaliautotradecareticketconsumeAPIResponseModel is 天猫汽车电子凭证核销 成功返回结果
type TmallaliautotradecareticketconsumeAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_aliauto_trade_car_eticket_consume_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *AliAutoResult `json:"result,omitempty" xml:"result,omitempty"`
}
