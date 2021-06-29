package fundplatform

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
制卡商通知制卡完成 APIResponse
alibaba.fundplatform.cardorders.status.make.finish

当制卡完成后，制卡商需要调用该接口，通知我们制卡已完成。
*/
type AlibabaFundplatformCardordersStatusMakeFinishAPIResponse struct {
    model.CommonResponse
    AlibabaFundplatformCardordersStatusMakeFinishResponse
}

type AlibabaFundplatformCardordersStatusMakeFinishResponse struct {
    XMLName xml.Name `xml:"alibaba_fundplatform_cardorders_status_make_finish_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *CardMakingInformResponse `json:"result,omitempty" xml:"result,omitempty"`

    
}
