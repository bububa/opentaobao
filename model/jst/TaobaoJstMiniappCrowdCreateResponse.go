package jst

import (
    "github.com/bububa/opentaobao/model"
)

/* 
小程序活动创建 APIResponse
taobao.jst.miniapp.crowd.create

小程序活动创建
*/
type TaobaoJstMiniappCrowdCreateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoJstMiniappCrowdCreateResponse `json:"jst_miniapp_crowd_create_response,omitempty"` 
    TaobaoJstMiniappCrowdCreateResponse
}

/* model for simplify = false
type TaobaoJstMiniappCrowdCreateResponse struct {

    // 活动Code，活动的唯一标识
    
    Result   string `json:"result,omitempty"`
    

    // 状态码
    
    RCode   int64 `json:"r_code,omitempty"`
    

}
*/

type TaobaoJstMiniappCrowdCreateResponse struct {

    // 活动Code，活动的唯一标识
    Result   string `json:"result,omitempty"`

    // 状态码
    RCode   int64 `json:"r_code,omitempty"`

}
