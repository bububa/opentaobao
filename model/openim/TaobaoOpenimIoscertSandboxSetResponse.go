package openim

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
设置开发环境证书 APIResponse
taobao.openim.ioscert.sandbox.set

设置开发环境证书
*/
type TaobaoOpenimIoscertSandboxSetAPIResponse struct {
    model.CommonResponse
    TaobaoOpenimIoscertSandboxSetResponse
}

type TaobaoOpenimIoscertSandboxSetResponse struct {
    XMLName xml.Name `xml:"openim_ioscert_sandbox_set_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 操作成功
    
    Code   string `json:"code,omitempty" xml:"code,omitempty"`

    
}
