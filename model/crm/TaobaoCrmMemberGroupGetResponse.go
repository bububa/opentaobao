package crm

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取买家身上的标签 APIResponse
taobao.crm.member.group.get

获取买家身上的标签，不返回标签的总人数
*/
type TaobaoCrmMemberGroupGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoCrmMemberGroupGetResponse `json:"taobao_crm_member_group_get_response,omitempty"`
}

type TaobaoCrmMemberGroupGetResponse struct {

    // 查询到的当前卖家的当前页的会员
    Groups   []Group `json:"groups,omitempty"`

}
