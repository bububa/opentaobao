package util

import (
    "github.com/bububa/opentaobao/model"
)

/* 
通过mixnick转换openuid APIResponse
taobao.openuid.get.bymixnick

通过mixnick转换openuid
*/
type TaobaoOpenuidGetBymixnickAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoOpenuidGetBymixnickResponse `json:"openuid_get_bymixnick_response,omitempty"` 
    TaobaoOpenuidGetBymixnickResponse
}

/* model for simplify = false
type TaobaoOpenuidGetBymixnickResponse struct {

    // OpenUID
    
    OpenUid   string `json:"open_uid,omitempty"`
    

}
*/

type TaobaoOpenuidGetBymixnickResponse struct {

    // OpenUID
    OpenUid   string `json:"open_uid,omitempty"`

}
