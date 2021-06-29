package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
淘鲜达外部商户老pos机产生的订单同步进淘鲜达 APIResponse
alibaba.wdk.oldpos.order.create

淘鲜达外部商户老pos机产生的订单同步进淘鲜达
*/
type AlibabaWdkOldposOrderCreateAPIResponse struct {
    model.CommonResponse
    AlibabaWdkOldposOrderCreateResponse
}

type AlibabaWdkOldposOrderCreateResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_oldpos_order_create_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果
    
    Result   *PosOrderCreateResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
