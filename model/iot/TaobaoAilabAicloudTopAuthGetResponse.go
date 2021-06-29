package iot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
登陆 APIResponse
taobao.ailab.aicloud.top.auth.get

登陆
*/
type TaobaoAilabAicloudTopAuthGetAPIResponse struct {
    model.CommonResponse
    TaobaoAilabAicloudTopAuthGetResponse
}

type TaobaoAilabAicloudTopAuthGetResponse struct {
    XMLName xml.Name `xml:"ailab_aicloud_top_auth_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AiCloudResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
