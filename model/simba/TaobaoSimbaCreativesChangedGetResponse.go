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
    Response *TaobaoSimbaCreativesChangedGetResponse `json:"taobao_simba_creatives_changed_get_response,omitempty"`
}

type TaobaoSimbaCreativesChangedGetResponse struct {

    // 广告创意分页对象
    Creatives   *CreativePage `json:"creatives,omitempty"`

}
