package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
逆向提交 APIResponse
alibaba.wdk.reverse.creatrefund

逆向申请提交
*/
type AlibabaWdkReverseCreatrefundAPIResponse struct {
    model.CommonResponse
    AlibabaWdkReverseCreatrefundResponse
}

type AlibabaWdkReverseCreatrefundResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_reverse_creatrefund_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // ReverseResult
    
    Result   *ReverseResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
