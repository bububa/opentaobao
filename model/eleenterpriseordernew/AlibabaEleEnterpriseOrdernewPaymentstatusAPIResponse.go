package eleenterpriseordernew

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleEnterpriseOrdernewPaymentstatusAPIResponse 设置订单支付 API返回值
// alibaba.ele.enterprise.ordernew.paymentstatus
//
// 设置订单支付成功
type AlibabaEleEnterpriseOrdernewPaymentstatusAPIResponse struct {
	model.CommonResponse
	AlibabaEleEnterpriseOrdernewPaymentstatusAPIResponseModel
}

// AlibabaEleEnterpriseOrdernewPaymentstatusAPIResponseModel is 设置订单支付 成功返回结果
type AlibabaEleEnterpriseOrdernewPaymentstatusAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ele_enterprise_ordernew_paymentstatus_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应code
	EnterpriseCode string `json:"enterprise_code,omitempty" xml:"enterprise_code,omitempty"`
	// 响应信息
	EnterpriseMsg string `json:"enterprise_msg,omitempty" xml:"enterprise_msg,omitempty"`
	// 请求id
	EnterpriseRequestid string `json:"enterprise_requestid,omitempty" xml:"enterprise_requestid,omitempty"`
}
