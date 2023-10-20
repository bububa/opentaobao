package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkOldposRefundCreateAPIResponse 五道口外部商户老pos机产生的退款单同步进盒马 API返回值
// alibaba.wdk.oldpos.refund.create
//
// 淘鲜达外部商户老pos机产生的退款单同步进淘鲜达
type AlibabaWdkOldposRefundCreateAPIResponse struct {
	model.CommonResponse
	AlibabaWdkOldposRefundCreateAPIResponseModel
}

// AlibabaWdkOldposRefundCreateAPIResponseModel is 五道口外部商户老pos机产生的退款单同步进盒马 成功返回结果
type AlibabaWdkOldposRefundCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_oldpos_refund_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PosRefundCreateResult `json:"result,omitempty" xml:"result,omitempty"`
}
