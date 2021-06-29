package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
增加创意 APIResponse
taobao.simba.creative.add

创建一个创意
*/
type TaobaoSimbaCreativeAddAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaCreativeAddResponse
}

type TaobaoSimbaCreativeAddResponse struct {
    XMLName xml.Name `xml:"simba_creative_add_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 新增加的创意对象
    
    Creative   *Creative `json:"creative,omitempty" xml:"creative,omitempty"`

    
}
