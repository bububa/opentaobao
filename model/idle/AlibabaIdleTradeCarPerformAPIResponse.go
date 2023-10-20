package idle

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidletradecarperformAPIResponse 二手车寄卖履约接口 API返回值
// alibaba.idle.trade.car.perform
//
// 二手车寄卖履约接口
type AlibabaidletradecarperformAPIResponse struct {
	model.CommonResponse
	AlibabaidletradecarperformAPIResponseModel
}

// AlibabaidletradecarperformAPIResponseModel is 二手车寄卖履约接口 成功返回结果
type AlibabaidletradecarperformAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_trade_car_perform_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *CarConsignmentResult `json:"result,omitempty" xml:"result,omitempty"`
}
