package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkorderSharestockInsuranceRefundgetAPIResponse 共享库存投保业务售后逆向订单数据获取 API返回值
// alibaba.wdkorder.sharestock.insurance.refundget
//
// 共享库存投保业务售后逆向订单数据获取
type AlibabaWdkorderSharestockInsuranceRefundgetAPIResponse struct {
	model.CommonResponse
	AlibabaWdkorderSharestockInsuranceRefundgetAPIResponseModel
}

// AlibabaWdkorderSharestockInsuranceRefundgetAPIResponseModel is 共享库存投保业务售后逆向订单数据获取 成功返回结果
type AlibabaWdkorderSharestockInsuranceRefundgetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdkorder_sharestock_insurance_refundget_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TopBaseResult `json:"result,omitempty" xml:"result,omitempty"`
}
