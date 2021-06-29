package tmallnr

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV查询活动 APIResponse
alibaba.lsy.crm.activity.getbaseinfo

ISV查询活动
*/
type AlibabaLsyCrmActivityGetbaseinfoAPIResponse struct {
    model.CommonResponse
    AlibabaLsyCrmActivityGetbaseinfoResponse
}

type AlibabaLsyCrmActivityGetbaseinfoResponse struct {
    XMLName xml.Name `xml:"alibaba_lsy_crm_activity_getbaseinfo_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *AlibabaLsyCrmActivityGetbaseinfoResultDo `json:"result,omitempty" xml:"result,omitempty"`

    
}
