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

// NewTaobaoCrmMembersGroupBatchaddRequest 初始化TaobaoCrmMembersGroupBatchaddAPIRequest对象
func NewTaobaoCrmMembersGroupBatchaddRequest() *TaobaoCrmMembersGroupBatchaddAPIRequest {
	return &TaobaoCrmMembersGroupBatchaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCrmMembersGroupBatchaddAPIRequest) GetApiMethodName() string {
	return "taobao.crm.members.group.batchadd"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCrmMembersGroupBatchaddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is GroupIds Setter
// 分组id
func (r *TaobaoCrmMembersGroupBatchaddAPIRequest) SetGroupIds(_groupIds []int64) error {
	r._groupIds = _groupIds
	r.Set("group_ids", _groupIds)
	return nil
}

// Get GroupIds Getter
func (r TaobaoCrmMembersGroupBatchaddAPIRequest) GetGroupIds() []int64 {
	return r._groupIds
}

// Set is BuyerNicks Setter
// 买家昵称列表
func (r *TaobaoCrmMembersGroupBatchaddAPIRequest) SetBuyerNicks(_buyerNicks []string) error {
	r._buyerNicks = _buyerNicks
	r.Set("buyer_nicks", _buyerNicks)
	return nil
}

// Get BuyerNicks Getter
func (r TaobaoCrmMembersGroupBatchaddAPIRequest) GetBuyerNicks() []string {
	return r._buyerNicks
}
