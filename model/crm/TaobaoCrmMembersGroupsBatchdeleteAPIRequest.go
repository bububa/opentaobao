package crm

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmMembersGroupsBatchdeleteAPIRequest 批量删除分组 API请求
// taobao.crm.members.groups.batchdelete
//
// 批量删除多个会员的公共分组，接口返回删除是否成功，该接口只删除多个会员的公共分组，不是公共分组的，不进行删除。如果入参只输入一个会员，则表示删除该会员的某些分组。
type TaobaoCrmMembersGroupsBatchdeleteAPIRequest struct {
	model.Params
	// 买家昵称列表
	_buyerNicks []string
	// 会员需要删除的分组
	_groupIds []int64
}

// NewTaobaoCrmMembersGroupsBatchdeleteRequest 初始化TaobaoCrmMembersGroupsBatchdeleteAPIRequest对象
func NewTaobaoCrmMembersGroupsBatchdeleteRequest() *TaobaoCrmMembersGroupsBatchdeleteAPIRequest {
	return &TaobaoCrmMembersGroupsBatchdeleteAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoCrmMembersGroupsBatchdeleteAPIRequest) Reset() {
	r._buyerNicks = r._buyerNicks[:0]
	r._groupIds = r._groupIds[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCrmMembersGroupsBatchdeleteAPIRequest) GetApiMethodName() string {
	return "taobao.crm.members.groups.batchdelete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCrmMembersGroupsBatchdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoCrmMembersGroupsBatchdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBuyerNicks is BuyerNicks Setter
// 买家昵称列表
func (r *TaobaoCrmMembersGroupsBatchdeleteAPIRequest) SetBuyerNicks(_buyerNicks []string) error {
	r._buyerNicks = _buyerNicks
	r.Set("buyer_nicks", _buyerNicks)
	return nil
}

// GetBuyerNicks BuyerNicks Getter
func (r TaobaoCrmMembersGroupsBatchdeleteAPIRequest) GetBuyerNicks() []string {
	return r._buyerNicks
}

// SetGroupIds is GroupIds Setter
// 会员需要删除的分组
func (r *TaobaoCrmMembersGroupsBatchdeleteAPIRequest) SetGroupIds(_groupIds []int64) error {
	r._groupIds = _groupIds
	r.Set("group_ids", _groupIds)
	return nil
}

// GetGroupIds GroupIds Getter
func (r TaobaoCrmMembersGroupsBatchdeleteAPIRequest) GetGroupIds() []int64 {
	return r._groupIds
}

var poolTaobaoCrmMembersGroupsBatchdeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoCrmMembersGroupsBatchdeleteRequest()
	},
}

// GetTaobaoCrmMembersGroupsBatchdeleteRequest 从 sync.Pool 获取 TaobaoCrmMembersGroupsBatchdeleteAPIRequest
func GetTaobaoCrmMembersGroupsBatchdeleteAPIRequest() *TaobaoCrmMembersGroupsBatchdeleteAPIRequest {
	return poolTaobaoCrmMembersGroupsBatchdeleteAPIRequest.Get().(*TaobaoCrmMembersGroupsBatchdeleteAPIRequest)
}

// ReleaseTaobaoCrmMembersGroupsBatchdeleteAPIRequest 将 TaobaoCrmMembersGroupsBatchdeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoCrmMembersGroupsBatchdeleteAPIRequest(v *TaobaoCrmMembersGroupsBatchdeleteAPIRequest) {
	v.Reset()
	poolTaobaoCrmMembersGroupsBatchdeleteAPIRequest.Put(v)
}
