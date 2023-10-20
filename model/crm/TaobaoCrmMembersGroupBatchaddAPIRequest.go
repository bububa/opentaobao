package crm

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmMembersGroupBatchaddAPIRequest 给一批会员添加一个分组 API请求
// taobao.crm.members.group.batchadd
//
// 为一批会员添加分组，接口返回添加是否成功,如至少有一个会员的分组添加成功，接口就返回成功，否则返回失败，如果当前会员已经拥有当前分组，则直接跳过
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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoCrmMembersGroupBatchaddAPIRequest) Reset() {
	r._groupIds = r._groupIds[:0]
	r._buyerNicks = r._buyerNicks[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCrmMembersGroupBatchaddAPIRequest) GetApiMethodName() string {
	return "taobao.crm.members.group.batchadd"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCrmMembersGroupBatchaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoCrmMembersGroupBatchaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGroupIds is GroupIds Setter
// 分组id
func (r *TaobaoCrmMembersGroupBatchaddAPIRequest) SetGroupIds(_groupIds []int64) error {
	r._groupIds = _groupIds
	r.Set("group_ids", _groupIds)
	return nil
}

// GetGroupIds GroupIds Getter
func (r TaobaoCrmMembersGroupBatchaddAPIRequest) GetGroupIds() []int64 {
	return r._groupIds
}

// SetBuyerNicks is BuyerNicks Setter
// 买家昵称列表
func (r *TaobaoCrmMembersGroupBatchaddAPIRequest) SetBuyerNicks(_buyerNicks []string) error {
	r._buyerNicks = _buyerNicks
	r.Set("buyer_nicks", _buyerNicks)
	return nil
}

// GetBuyerNicks BuyerNicks Getter
func (r TaobaoCrmMembersGroupBatchaddAPIRequest) GetBuyerNicks() []string {
	return r._buyerNicks
}

var poolTaobaoCrmMembersGroupBatchaddAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoCrmMembersGroupBatchaddRequest()
	},
}

// GetTaobaoCrmMembersGroupBatchaddRequest 从 sync.Pool 获取 TaobaoCrmMembersGroupBatchaddAPIRequest
func GetTaobaoCrmMembersGroupBatchaddAPIRequest() *TaobaoCrmMembersGroupBatchaddAPIRequest {
	return poolTaobaoCrmMembersGroupBatchaddAPIRequest.Get().(*TaobaoCrmMembersGroupBatchaddAPIRequest)
}

// ReleaseTaobaoCrmMembersGroupBatchaddAPIRequest 将 TaobaoCrmMembersGroupBatchaddAPIRequest 放入 sync.Pool
func ReleaseTaobaoCrmMembersGroupBatchaddAPIRequest(v *TaobaoCrmMembersGroupBatchaddAPIRequest) {
	v.Reset()
	poolTaobaoCrmMembersGroupBatchaddAPIRequest.Put(v)
}
