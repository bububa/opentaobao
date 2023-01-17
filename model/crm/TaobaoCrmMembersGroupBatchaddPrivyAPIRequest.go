package crm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmMembersGroupBatchaddPrivyAPIRequest 一批会员添加分组(隐私号版） API请求
// taobao.crm.members.group.batchadd.privy
//
// 为一批会员添加分组，接口返回添加是否成功,如至少有一个会员的分组添加成功，接口就返回成功，否则返回失败，如果当前会员已经拥有当前分组，则直接跳过
type TaobaoCrmMembersGroupBatchaddPrivyAPIRequest struct {
	model.Params
	// 分组id
	_groupIds []int64
	// ouid列表
	_ouids []string
}

// NewTaobaoCrmMembersGroupBatchaddPrivyRequest 初始化TaobaoCrmMembersGroupBatchaddPrivyAPIRequest对象
func NewTaobaoCrmMembersGroupBatchaddPrivyRequest() *TaobaoCrmMembersGroupBatchaddPrivyAPIRequest {
	return &TaobaoCrmMembersGroupBatchaddPrivyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCrmMembersGroupBatchaddPrivyAPIRequest) GetApiMethodName() string {
	return "taobao.crm.members.group.batchadd.privy"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCrmMembersGroupBatchaddPrivyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoCrmMembersGroupBatchaddPrivyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGroupIds is GroupIds Setter
// 分组id
func (r *TaobaoCrmMembersGroupBatchaddPrivyAPIRequest) SetGroupIds(_groupIds []int64) error {
	r._groupIds = _groupIds
	r.Set("group_ids", _groupIds)
	return nil
}

// GetGroupIds GroupIds Getter
func (r TaobaoCrmMembersGroupBatchaddPrivyAPIRequest) GetGroupIds() []int64 {
	return r._groupIds
}

// SetOuids is Ouids Setter
// ouid列表
func (r *TaobaoCrmMembersGroupBatchaddPrivyAPIRequest) SetOuids(_ouids []string) error {
	r._ouids = _ouids
	r.Set("ouids", _ouids)
	return nil
}

// GetOuids Ouids Getter
func (r TaobaoCrmMembersGroupBatchaddPrivyAPIRequest) GetOuids() []string {
	return r._ouids
}
