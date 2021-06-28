package crm

import (
    "github.com/bububa/opentaobao/model"
)

/* 
编辑会员资料 APIResponse
taobao.crm.memberinfo.update

编辑会员的基本资料，接口返回会员信息修改是否成功
*/
type TaobaoCrmMemberinfoUpdateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoCrmMemberinfoUpdateResponse `json:"crm_memberinfo_update_response,omitempty"` 
    TaobaoCrmMemberinfoUpdateResponse
}

/* model for simplify = false
type TaobaoCrmMemberinfoUpdateResponse struct {

    // 会员信息修改是否成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type TaobaoCrmMemberinfoUpdateResponse struct {

    // 会员信息修改是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
