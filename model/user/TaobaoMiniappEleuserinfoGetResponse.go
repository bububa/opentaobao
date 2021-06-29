package user

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取饿了么用户信息 APIResponse
taobao.miniapp.eleuserinfo.get

获取饿了么用户信息
*/
type TaobaoMiniappEleuserinfoGetAPIResponse struct {
    model.CommonResponse
    TaobaoMiniappEleuserinfoGetResponse
}

type TaobaoMiniappEleuserinfoGetResponse struct {
    XMLName xml.Name `xml:"miniapp_eleuserinfo_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // traceId
    
    TraceId   string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`

    
    // 1
    
    Result   *EleUicInfo `json:"result,omitempty" xml:"result,omitempty"`

    
}
