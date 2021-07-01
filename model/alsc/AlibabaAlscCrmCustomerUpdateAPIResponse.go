package alsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmCustomerUpdateAPIResponse
更新顾客信息 API返回值
alibaba.alsc.crm.customer.update

更新顾客信息 */
type AlibabaAlscCrmCustomerUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmCustomerUpdateAPIResponseModel
}

// AlibabaAlscCrmCustomerUpdateAPIResponseModel is 更新顾客信息 成功返回结果
type AlibabaAlscCrmCustomerUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_customer_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
