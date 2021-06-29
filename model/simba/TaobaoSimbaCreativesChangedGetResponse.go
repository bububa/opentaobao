package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
分页获取修改过的广告创意ID和修改时间 APIResponse
taobao.simba.creatives.changed.get

分页获取修改过的广告创意ID和修改时间
*/
type TaobaoSimbaCreativesChangedGetAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaCreativesChangedGetResponse
}

type TaobaoSimbaCreativesChangedGetResponse struct {
    XMLName xml.Name `xml:"simba_creatives_changed_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 广告创意分页对象
    
    Creatives   *CreativePage `json:"creatives,omitempty" xml:"creatives,omitempty"`

    
}
