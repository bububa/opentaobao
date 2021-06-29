package maitix

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦-直接退票 APIResponse
alibaba.damai.maitix.order.directrefund

大麦-退票
*/
type AlibabaDamaiMaitixOrderDirectrefundAPIResponse struct {
    model.CommonResponse
    AlibabaDamaiMaitixOrderDirectrefundResponse
}

type AlibabaDamaiMaitixOrderDirectrefundResponse struct {
    XMLName xml.Name `xml:"alibaba_damai_maitix_order_directrefund_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 退票返回对象
    
    Result   *MxResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
