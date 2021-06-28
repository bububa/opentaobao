package util

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取授权session信息 APIResponse
taobao.openlink.session.get

帮助第三方isv生成三方session
*/
type TaobaoOpenlinkSessionGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoOpenlinkSessionGetResponse `json:"openlink_session_get_response,omitempty"` 
    TaobaoOpenlinkSessionGetResponse
}

/* model for simplify = false
type TaobaoOpenlinkSessionGetResponse struct {

    // result
    
    Result  *struct {
        TaobaoOpenlinkSessionGetResult  *TaobaoOpenlinkSessionGetResult `json:"taobao_openlink_session_get_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoOpenlinkSessionGetResponse struct {

    // result
    Result   *TaobaoOpenlinkSessionGetResult `json:"result,omitempty"`

}
