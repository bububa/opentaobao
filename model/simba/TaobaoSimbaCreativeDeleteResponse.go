package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
删除创意 APIResponse
taobao.simba.creative.delete

删除一个创意
*/
type TaobaoSimbaCreativeDeleteAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSimbaCreativeDeleteResponse `json:"simba_creative_delete_response,omitempty"` 
    TaobaoSimbaCreativeDeleteResponse
}

/* model for simplify = false
type TaobaoSimbaCreativeDeleteResponse struct {

    // 被删除的创意对象
    
    Creative  *struct {
        Creative  *Creative `json:"creative,omitempty"`
    } `json:"creative,omitempty"`
    

}
*/

type TaobaoSimbaCreativeDeleteResponse struct {

    // 被删除的创意对象
    Creative   *Creative `json:"creative,omitempty"`

}
