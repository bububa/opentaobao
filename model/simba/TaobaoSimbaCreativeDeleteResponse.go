package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除创意 APIResponse
taobao.simba.creative.delete

删除一个创意
*/
type TaobaoSimbaCreativeDeleteAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaCreativeDeleteResponse
}

type TaobaoSimbaCreativeDeleteResponse struct {
    XMLName xml.Name `xml:"simba_creative_delete_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 被删除的创意对象
    
    Creative   *Creative `json:"creative,omitempty" xml:"creative,omitempty"`

    
}
