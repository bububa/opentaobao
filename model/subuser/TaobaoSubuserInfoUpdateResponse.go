package subuser

import (
    "github.com/bububa/opentaobao/model"
)

/* 
修改指定账户子账号的基本信息 APIResponse
taobao.subuser.info.update

修改指定账户子账号的基本信息（通过主账号登陆只能修改属于该主账号的子账号基本信息）
*/
type TaobaoSubuserInfoUpdateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSubuserInfoUpdateResponse `json:"subuser_info_update_response,omitempty"` 
    TaobaoSubuserInfoUpdateResponse
}

/* model for simplify = false
type TaobaoSubuserInfoUpdateResponse struct {

    // 操作是否成功 true:操作成功; false:操作失败
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type TaobaoSubuserInfoUpdateResponse struct {

    // 操作是否成功 true:操作成功; false:操作失败
    IsSuccess   bool `json:"is_success,omitempty"`

}
