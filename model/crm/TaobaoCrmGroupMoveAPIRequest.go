package crm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmGroupMoveAPIRequest 分组移动 API请求
// taobao.crm.group.move
//
// 将一个分组下的所有会员移动到另一个分组，会员从原分组中删除<br/>注：移动属性为异步任务建议先调用taobao.crm.grouptask.check 确保涉及属性上没有任务。
type TaobaoCrmGroupMoveAPIRequest struct {
	model.Params
	// 需要移动的分组
	_fromGroupId int64
	// 目的分组
	_toGroupId int64
}

// NewTaobaoCrmGroupMoveRequest 初始化TaobaoCrmGroupMoveAPIRequest对象
func NewTaobaoCrmGroupMoveRequest() *TaobaoCrmGroupMoveAPIRequest {
	return &TaobaoCrmGroupMoveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCrmGroupMoveAPIRequest) GetApiMethodName() string {
	return "taobao.crm.group.move"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCrmGroupMoveAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetFromGroupId is FromGroupId Setter
// 需要移动的分组
func (r *TaobaoCrmGroupMoveAPIRequest) SetFromGroupId(_fromGroupId int64) error {
	r._fromGroupId = _fromGroupId
	r.Set("from_group_id", _fromGroupId)
	return nil
}

// GetFromGroupId FromGroupId Getter
func (r TaobaoCrmGroupMoveAPIRequest) GetFromGroupId() int64 {
	return r._fromGroupId
}

// SetToGroupId is ToGroupId Setter
// 目的分组
func (r *TaobaoCrmGroupMoveAPIRequest) SetToGroupId(_toGroupId int64) error {
	r._toGroupId = _toGroupId
	r.Set("to_group_id", _toGroupId)
	return nil
}

// GetToGroupId ToGroupId Getter
func (r TaobaoCrmGroupMoveAPIRequest) GetToGroupId() int64 {
	return r._toGroupId
}
