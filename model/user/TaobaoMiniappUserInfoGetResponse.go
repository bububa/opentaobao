package user

import (
    "github.com/bububa/opentaobao/model"
)

/* 
用户开放信息获取 APIResponse
taobao.miniapp.userInfo.get

获取用户的 openId，snsNick（如果用户设置过的话），和加密头像链接
*/
type TaobaoMiniappUserInfoGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoMiniappUserInfoGetResponse `json:"miniapp_userInfo_get_response,omitempty"` 
    TaobaoMiniappUserInfoGetResponse
}

/* model for simplify = false
type TaobaoMiniappUserInfoGetResponse struct {

    // 接口返回model
    
    Result  *struct {
        TaobaoMiniappUserInfoGetResult  *TaobaoMiniappUserInfoGetResult `json:"taobao_miniapp_user_info_get_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoMiniappUserInfoGetResponse struct {

    // 接口返回model
    Result   *TaobaoMiniappUserInfoGetResult `json:"result,omitempty"`

}
