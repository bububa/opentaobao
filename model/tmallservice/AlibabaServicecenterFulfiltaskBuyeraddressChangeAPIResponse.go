package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaServicecenterFulfiltaskBuyeraddressChangeAPIResponse
修改消费者服务地址 API返回值
alibaba.servicecenter.fulfiltask.buyeraddress.change

当消费者反馈自己的服务地址错误时，可以电话联系服务商修改为正确地址，服务商只能修改派给自己的单子 */
type AlibabaServicecenterFulfiltaskBuyeraddressChangeAPIResponse struct {
	model.CommonResponse
	AlibabaServicecenterFulfiltaskBuyeraddressChangeAPIResponseModel
}

// AlibabaServicecenterFulfiltaskBuyeraddressChangeAPIResponseModel is 修改消费者服务地址 成功返回结果
type AlibabaServicecenterFulfiltaskBuyeraddressChangeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_servicecenter_fulfiltask_buyeraddress_change_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaServicecenterFulfiltaskBuyeraddressChangeResult `json:"result,omitempty" xml:"result,omitempty"`
}
