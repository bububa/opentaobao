package util

import (
    "github.com/bububa/opentaobao/model"
)

/* 
TMC授权token APIResponse
taobao.tmc.auth.get

TMC连接授权Token
*/
type TaobaoTmcAuthGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTmcAuthGetResponse `json:"tmc_auth_get_response,omitempty"` 
    TaobaoTmcAuthGetResponse
}

/* model for simplify = false
type TaobaoTmcAuthGetResponse struct {

    // result
    
    Result   string `json:"result,omitempty"`
    

}
*/

type TaobaoTmcAuthGetResponse struct {

    // result
    Result   string `json:"result,omitempty"`

}
