package crm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoCrmMembersGroupsBatchdeleteAPIRequest
批量删除分组 API请求
taobao.crm.members.groups.batchdelete

批量删除多个会员的公共分组，接口返回删除是否成功，该接口只删除多个会员的公共分组，不是公共分组的，不进行删除。如果入参只输入一个会员，则表示删除该会员的某些分组。 */
type TaobaoCrmMembersGroupsBatchdeleteAPIRequest struct {
	model.Params
	// 买家昵称列表
	_buyerNicks []string
	// 会员需要删除的分组
	_groupIds []int64
}

// New
