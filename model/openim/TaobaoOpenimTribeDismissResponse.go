package openim

import (
    "github.com/bububa/opentaobao/model"
)

/* 
OPENIM群解散 APIResponse
taobao.openim.tribe.dismiss

OPENIM群解散
*/
type TaobaoOpenimTribeDismissAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoOpenimTribeDismissResponse `json:"openim_tribe_dismiss_response,omitempty"` 
    TaobaoOpenimTribeDismissResponse
}

/* model for simplify = false
type TaobaoOpenimTribeDismissResponse struct {

    // 群服务code
    
    TribeCode   int64 `json:"tribe_code,omitempty"`
    

}
*/

type TaobaoOpenimTribeDismissResponse struct {

    // 群服务code
    TribeCode   int64 `json:"tribe_code,omitempty"`

}
