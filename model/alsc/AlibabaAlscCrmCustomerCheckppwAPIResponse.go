package alsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmCustomerCheckppwAPIResponse
校验支付密码 API返回值
alibaba.alsc.crm.customer.checkppw

校验支付密码 */
type AlibabaAlscCrmCustomerCheckppwAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmCustomerCheckppwAPIResponseModel
}

// AlibabaAlscCrmCustomerCheckppwAPIResponseModel is 校验支付密码 成功返回结果
type AlibabaAlscCrmCustomerCheckppwAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_customer_checkppw_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
