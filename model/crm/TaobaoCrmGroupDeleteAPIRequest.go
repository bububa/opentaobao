package crm

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCrmGroupDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.crm.group.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCrmGroupDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
