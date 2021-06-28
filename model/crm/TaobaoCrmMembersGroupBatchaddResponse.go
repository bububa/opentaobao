package crm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
给一批会员添加一个分组 APIResponse
taobao.crm.members.group.batchadd

为一批会员添加分组，接口返回添加是否成功,如至少有一个会员的分组添加成功，接口就返回成功，否则返回失败，如果当前会员已经拥有当前分组，则直接跳过
*/
type TaobaoCrmMembersGroupBatchaddAPIResponse struct {
    model.CommonResponse
    TaobaoCrmMembersGroupBatchaddResponse
}

type TaobaoCrmMembersGroupBatchaddResponse struct {
    XMLName xml.Name `xml:"crm_members_group_batchadd_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 添加操作是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
