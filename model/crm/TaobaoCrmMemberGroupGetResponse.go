package crm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取买家身上的标签 APIResponse
taobao.crm.member.group.get

获取买家身上的标签，不返回标签的总人数
*/
type TaobaoCrmMemberGroupGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"crm_member_group_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 查询到的当前卖家的当前页的会员
    
    Groups   []Group `json:"groups,omitempty" xml:"