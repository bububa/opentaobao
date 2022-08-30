package crm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmMembersGroupsBatchdeletePrivyAPIRequest 批量删除分组（隐私号版） API请求
// taobao.crm.members.groups.batchdelete.privy
//
// 批量删除多个会员的公共分组，接口返回删除是否成功，该接口只删除多个会员的公共分组，不是公共分组的，不进行删除。如果入参只输入一个会员，则表示删除该会员的某些分组。
type TaobaoCrmMembersGroupsBatchdeletePrivyAPIRequest struct {
	model.Params
	// 会员需要删除的分组
	_groupIds []int64
	// ouid列表
	_ouids []string
}

// NewTaobaoCrmMembersGroupsBatchdeletePrivyRequest 初始化TaobaoCrmMembersGroupsBatchdeletePrivyAPIRequest对象
func NewTaobaoCrmMembersGroupsBatchdeletePrivyRequest() *TaobaoCrmMembersGroupsBatchdeletePrivyAPIRequest {
	return &TaobaoCrmMembersGroupsBatchdeletePrivyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCrmMembersGroupsBatchdeletePrivyAPIRequest) GetApiMethodName() string {
	return "taobao.crm.members.groups.batchdelete.privy"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCrmMembersGroupsBatchdeletePrivyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetGroupIds is GroupIds Setter
// 会员需要删除的分组
func (r *TaobaoCrmMembersGroupsBatchdeletePrivyAPIRequest) SetGroupIds(_groupIds []int64) error {
	r._groupIds = _groupIds
	r.Set("group_ids", _groupIds)
	return nil
}

// GetGroupIds GroupIds Getter
func (r TaobaoCrmMembersGroupsBatchdeletePrivyAPIRequest) GetGroupIds() []int64 {
	return r._groupIds
}

// SetOuids is Ouids Setter
// ouid列表
func (r *TaobaoCrmMembersGroupsBatchdeletePrivyAPIRequest) SetOuids(_ouids []string) error {
	r._ouids = _ouids
	r.Set("ouids", _ouids)
	return nil
}

// GetOuids Ouids Getter
func (r TaobaoCrmMembersGroupsBatchdeletePrivyAPIRequest) GetOuids() []string {
	return r._ouids
}
