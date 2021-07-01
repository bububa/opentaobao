package alsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmCustomerCreateAPIResponse
创建顾客 API返回值
alibaba.alsc.crm.customer.create

开放本地生活创建顾客功能 */
type AlibabaAlscCrmCustomerCreateAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmCustomerCreateAPIResponseModel
}

// AlibabaAlscCrmCustomerCreateAPIResponseModel is 创建顾客 成功返回结果
type AlibabaAlscCrmCustomerCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_customer_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
