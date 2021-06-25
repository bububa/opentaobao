package crm

import (
    "github.com/bububa/opentaobao/model"
)

/* 
批量删除分组 APIResponse
taobao.crm.members.groups.batchdelete

批量删除多个会员的公共分组，接口返回删除是否成功，该接口只删除多个会员的公共分组，不是公共分组的，不进行删除。如果入参只输入一个会员，则表示删除该会员的某些分组。
*/
type TaobaoCrmMembersGroupsBatchdeleteAPIResponse struct {
    model.CommonResponse
    Response *TaobaoCrmMembersGroupsBatchdeleteResponse `json:"taobao_crm_members_groups_batchdelete_response,omitempty"`
}

type TaobaoCrmMembersGroupsBatchdeleteResponse struct {

    // 删除是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
