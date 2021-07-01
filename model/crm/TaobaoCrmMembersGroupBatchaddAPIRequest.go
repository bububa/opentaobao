package crm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoCrmMembersGroupBatchaddAPIRequest
给一批会员添加一个分组 API请求
taobao.crm.members.group.batchadd

为一批会员添加分组，接口返回添加是否成功,如至少有一个会员的分组添加成功，接口就返回成功，否则返回失败，如果当前会员已经拥有当前分组，则直接跳过 */
type TaobaoCrmMembersGroupBatchaddAPIRequest struct {
	model.Params
	// 分组id
	_groupIds []int64
	// 买家昵称列表
	_buyerNicks []string
}

// New
