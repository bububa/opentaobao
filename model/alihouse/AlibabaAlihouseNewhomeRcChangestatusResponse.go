package alihouse

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
图文草稿状态更新 APIResponse
alibaba.alihouse.newhome.rc.changestatus

图文草稿状态更新
*/
type AlibabaAlihouseNewhomeRcChangestatusAPIResponse struct {
    model.CommonResponse
    AlibabaAlihouseNewhomeRcChangestatusResponse
}

type AlibabaAlihouseNewhomeRcChangestatusResponse struct {
    XMLName xml.Name `xml:"alibaba_alihouse_newhome_rc_changestatus_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaAlihouseNewhomeRcChangestatusResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
