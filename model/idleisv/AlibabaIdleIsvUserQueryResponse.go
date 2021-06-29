package idleisv

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商ISV闲鱼用户信息查询 APIResponse
alibaba.idle.isv.user.query

服务商ISV闲鱼用户信息查询
*/
type AlibabaIdleIsvUserQueryAPIResponse struct {
    model.CommonResponse
    AlibabaIdleIsvUserQueryResponse
}

type AlibabaIdleIsvUserQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_idle_isv_user_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 系统自动生成
    
    Result   *TopResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
