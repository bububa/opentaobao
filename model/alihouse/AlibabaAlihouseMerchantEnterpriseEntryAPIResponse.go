package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousemerchantenterpriseentryAPIResponse 机构入驻 API返回值
// alibaba.alihouse.merchant.enterprise.entry
//
// 机构入驻
type AlibabaalihousemerchantenterpriseentryAPIResponse struct {
	model.CommonResponse
	AlibabaalihousemerchantenterpriseentryAPIResponseModel
}

// AlibabaalihousemerchantenterpriseentryAPIResponseModel is 机构入驻 成功返回结果
type AlibabaalihousemerchantenterpriseentryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_merchant_enterprise_entry_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaalihousemerchantenterpriseentryResult `json:"result,omitempty" xml:"result,omitempty"`
}
