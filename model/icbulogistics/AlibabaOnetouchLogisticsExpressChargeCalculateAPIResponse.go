package icbulogistics

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaonetouchlogisticsexpresschargecalculateAPIResponse 计算快递运费&下单参数校验 API返回值
// alibaba.onetouch.logistics.express.charge.calculate
//
// 计算快递运费、下单参数校验
type AlibabaonetouchlogisticsexpresschargecalculateAPIResponse struct {
	model.CommonResponse
	AlibabaonetouchlogisticsexpresschargecalculateAPIResponseModel
}

// AlibabaonetouchlogisticsexpresschargecalculateAPIResponseModel is 计算快递运费&下单参数校验 成功返回结果
type AlibabaonetouchlogisticsexpresschargecalculateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_onetouch_logistics_express_charge_calculate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaonetouchlogisticsexpresschargecalculateResult `json:"result,omitempty" xml:"result,omitempty"`
}
