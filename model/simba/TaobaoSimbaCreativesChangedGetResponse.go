package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
分页获取修改过的广告创意ID和修改时间 APIResponse
taobao.simba.creatives.changed.get

分页获取修改过的广告创意ID和修改时间
*/
type TaobaoSimbaCreativesChangedGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSimbaCreativesChangedGetResponse `json:"simba_creatives_changed_get_response,omitempty"` 
    TaobaoSimbaCreativesChangedGetResponse
}

/* model for simplify = false
type TaobaoSimbaCreativesChangedGetResponse struct {

    // 广告创意分页对象
    
    Creatives  *struct {
        CreativePage  *CreativePage `json:"creative_page,omitempty"`
    } `json:"creatives,omitempty"`
    

}
*/

type TaobaoSimbaCreativesChangedGetResponse struct {

    // 广告创意分页对象
    Creatives   *CreativePage `json:"creatives,omitempty"`

}
