package crm

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmGroupDeleteAPIRequest 删除分组 API请求
// taobao.crm.group.delete
//
// 将该分组下的所有会员移除出该组，同时删除该分组。注：删除分组为异步任务，必须先调用taobao.crm.grouptask.check 确保涉及属性上没有任务。
type TaobaoCrmGroupDeleteAPIRequest struct {
	model.Params
	// 要删除的分组id
	_groupId int64
}

// NewTaobaoCrmGroupDeleteRequest 初始化TaobaoCrmGroupDeleteAPIRequest对象
func NewTaobaoCrmGroupDeleteRequest() *TaobaoCrmGroupDeleteAPIRequest {
	return &TaobaoCrmGroupDeleteAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoCrmGroupDeleteAPIRequest) Reset() {
	r._groupId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCrmGroupDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.crm.group.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCrmGroupDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoCrmGroupDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGroupId is GroupId Setter
// 要删除的分组id
func (r *TaobaoCrmGroupDeleteAPIRequest) SetGroupId(_groupId int64) error {
	r._groupId = _groupId
	r.Set("group_id", _groupId)
	return nil
}

// GetGroupId GroupId Getter
func (r TaobaoCrmGroupDeleteAPIRequest) GetGroupId() int64 {
	return r._groupId
}

var poolTaobaoCrmGroupDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoCrmGroupDeleteRequest()
	},
}

// GetTaobaoCrmGroupDeleteRequest 从 sync.Pool 获取 TaobaoCrmGroupDeleteAPIRequest
func GetTaobaoCrmGroupDeleteAPIRequest() *TaobaoCrmGroupDeleteAPIRequest {
	return poolTaobaoCrmGroupDeleteAPIRequest.Get().(*TaobaoCrmGroupDeleteAPIRequest)
}

// ReleaseTaobaoCrmGroupDeleteAPIRequest 将 TaobaoCrmGroupDeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoCrmGroupDeleteAPIRequest(v *TaobaoCrmGroupDeleteAPIRequest) {
	v.Reset()
	poolTaobaoCrmGroupDeleteAPIRequest.Put(v)
}
