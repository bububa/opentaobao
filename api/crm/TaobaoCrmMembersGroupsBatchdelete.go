package crm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/crm"
)

// TaobaoCrmMembersGroupsBatchdelete 批量删除分组
// taobao.crm.members.groups.batchdelete
//
// 批量删除多个会员的公共分组，接口返回删除是否成功，该接口只删除多个会员的公共分组，不是公共分组的，不进行删除。如果入参只输入一个会员，则表示删除该会员的某些分组。
func TaobaoCrmMembersGroupsBatchdelete(clt *core.SDKClient, req *crm.TaobaoCrmMembersGroupsBatchdeleteAPIRequest, resp *crm.TaobaoCrmMembersGroupsBatchdeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
