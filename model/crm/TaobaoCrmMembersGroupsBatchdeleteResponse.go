package crm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量删除分组 API返回值 
taobao.crm.members.groups.batchdelete

批量删除多个会员的公共分组，接口返回删除是否成功，该接口只删除多个会员的公共分组，不是公共分组的，不进行删除。如果入参只输入一个会员，则表示删除该会员的某些分组。
*/
type TaobaoCrmMembersGroupsBatchdeleteAPIResponse struct {
    model.CommonResponse
    TaobaoCrmMembersGroupsBatchdeleteResponse
}

// 批量删除分组 成功返回结果
type TaobaoCrmMembersGroupsBatchdeleteResponse struct {
    XMLName xml.Name `xml:"crm_members_groups_batchdelete_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 删除是否成功
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
