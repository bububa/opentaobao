package subuser

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取指定账户子账号的详细信息 APIResponse
taobao.subuser.fullinfo.get

获取指定账户子账号的详细信息，其中包括子账号的账号信息以及员工、部门、职务信息（只能通过主账号登陆并查询属于该主账号下的某个子账号详细信息）
*/
type TaobaoSubuserFullinfoGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSubuserFullinfoGetResponse `json:"subuser_fullinfo_get_response,omitempty"` 
    TaobaoSubuserFullinfoGetResponse
}

/* model for simplify = false
type TaobaoSubuserFullinfoGetResponse struct {

    // 子账号详细信息，其中包括账号基本信息、员工信息和部门职务信息
    
    SubFullinfo  *struct {
        SubUserFullInfo  *SubUserFullInfo `json:"sub_user_full_info,omitempty"`
    } `json:"sub_fullinfo,omitempty"`
    

}
*/

type TaobaoSubuserFullinfoGetResponse struct {

    // 子账号详细信息，其中包括账号基本信息、员工信息和部门职务信息
    SubFullinfo   *SubUserFullInfo `json:"sub_fullinfo,omitempty"`

}
