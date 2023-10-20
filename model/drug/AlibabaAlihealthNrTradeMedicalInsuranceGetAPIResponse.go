package drug

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthnrtrademedicalinsurancegetAPIResponse 阿里健康医保支付信息获取 API返回值
// alibaba.alihealth.nr.trade.medical.insurance.get
//
// 阿里健康医保支付信息获取
type AlibabaalihealthnrtrademedicalinsurancegetAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthnrtrademedicalinsurancegetAPIResponseModel
}

// AlibabaalihealthnrtrademedicalinsurancegetAPIResponseModel is 阿里健康医保支付信息获取 成功返回结果
type AlibabaalihealthnrtrademedicalinsurancegetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_nr_trade_medical_insurance_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值总
	Result *ResponseResult `json:"result,omitempty" xml:"result,omitempty"`
}
