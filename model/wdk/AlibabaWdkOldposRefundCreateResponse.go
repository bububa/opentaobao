package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
五道口外部商户老pos机产生的退款单同步进盒马 APIResponse
alibaba.wdk.oldpos.refund.create

淘鲜达外部商户老pos机产生的退款单同步进淘鲜达
*/
type AlibabaWdkOldposRefundCreateAPIResponse struct {
    model.CommonResponse
    AlibabaWdkOldposRefundCreateResponse
}

type AlibabaWdkOldposRefundCreateResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_oldpos_refund_create_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *PosRefundCreateResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
