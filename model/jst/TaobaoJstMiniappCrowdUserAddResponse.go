package jst

import (
    "github.com/bububa/opentaobao/model"
)

/* 
小程序添加用户到指定的活动 APIResponse
taobao.jst.miniapp.crowd.user.add

小程序添加用户到指定的活动
*/
type TaobaoJstMiniappCrowdUserAddAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoJstMiniappCrowdUserAddResponse `json:"jst_miniapp_crowd_user_add_response,omitempty"` 
    TaobaoJstMiniappCrowdUserAddResponse
}

/* model for simplify = false
type TaobaoJstMiniappCrowdUserAddResponse struct {

    // 添加成功
    
    Result   bool `json:"result,omitempty"`
    

    // 请求成功
    
    RCode   int64 `json:"r_code,omitempty"`
    

}
*/

type TaobaoJstMiniappCrowdUserAddResponse struct {

    // 添加成功
    Result   bool `json:"result,omitempty"`

    // 请求成功
    RCode   int64 `json:"r_code,omitempty"`

}
