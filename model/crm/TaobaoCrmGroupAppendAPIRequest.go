package crm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmGroupAppendAPIRequest 将一个分组添加到另外一个分组 API请求
// taobao.crm.group.append
//
// 将某分组下的所有会员添加到另一个分组,注：1.该操作为异步任务，建议先调用taobao.crm.grouptask.check 确保涉及分组上没有任务；2.若分组下某会员分组数超最大限额，则该会员不会被添加到新分组，同时不影响其余会员添加分组，接口调用依然返回成功。
type TaobaoCrmGroupAppendAPIRequest struct {
	model.Params
	// 添加的来源分组
	_fromGroupId int64
	// 添加的目标分组
	_toGroupId int64
}

// NewTaobaoCrmGroupAppendRequest 初始化TaobaoCrmGroupAppendAPIRequest对象
func NewTaobaoCrmGroupAppendRequest() *TaobaoCrmGroupAppendAPIRequest {
	return &TaobaoCrmGroupAppendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCrmGroupAppendAPIRequest) GetApiMethodName() string {
	return "taobao.crm.group.append"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCrmGroupAppendAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetFromGroupId is FromGroupId Setter
// 添加的来源分组
func (r *TaobaoCrmGroupAppendAPIRequest) SetFromGroupId(_fromGroupId int64) error {
	r._fromGroupId = _fromGroupId
	r.Set("from_group_id", _fromGroupId)
	return nil
}

// GetFromGroupId FromGroupId Getter
func (r TaobaoCrmGroupAppendAPIRequest) GetFromGroupId() int64 {
	return r._fromGroupId
}

// SetToGroupId is ToGroupId Setter
// 添加的目标分组
func (r *TaobaoCrmGroupAppendAPIRequest) SetToGroupId(_toGroupId int64) error {
	r._toGroupId = _toGroupId
	r.Set("to_group_id", _toGroupId)
	return nil
}

// GetToGroupId ToGroupId Getter
func (r TaobaoCrmGroupAppendAPIRequest) GetToGroupId() int64 {
	return r._toGroupId
}
