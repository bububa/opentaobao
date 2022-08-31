package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeActivityCustomerSaveAPIResponse 销售活动批量保存定向用户 API返回值
// alibaba.alihouse.newhome.activity.customer.save
//
// 销售活动批量保存定向用户
type AlibabaAlihouseNewhomeActivityCustomerSaveAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeActivityCustomerSaveAPIResponseModel
}

// AlibabaAlihouseNewhomeActivityCustomerSaveAPIResponseModel is 销售活动批量保存定向用户 成功返回结果
type AlibabaAlihouseNewhomeActivityCustomerSaveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_activity_customer_save_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回报文属性
	Result *AlibabaAlihouseNewhomeActivityCustomerSaveResult `json:"result,omitempty" xml:"result,omitempty"`
}
