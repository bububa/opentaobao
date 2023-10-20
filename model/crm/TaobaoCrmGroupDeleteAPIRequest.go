package crm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaocrmgroupdeleteAPIRequest 删除分组 API请求
// taobao.crm.group.delete
//
// 将该分组下的所有会员移除出该组，同时删除该分组。注：删除分组为异步任务，必须先调用taobao.crm.grouptask.check 确保涉及属性上没有任务。
type TaobaocrmgroupdeleteAPIRequest struct {
	model.Params
	// 要删除的分组id
	_groupId int64
}

// NewTaobaocrmgroupdeleteRequest 初始化TaobaocrmgroupdeleteAPIRequest对象
func NewTaobaocrmgroupdeleteRequest() *TaobaocrmgroupdeleteAPIRequest {
	return &TaobaocrmgroupdeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaocrmgroupdeleteAPIRequest) GetApiMethodName() string {
	return "taobao.crm.group.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaocrmgroupdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaocrmgroupdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGroupId is GroupId Setter
// 要删除的分组id
func (r *TaobaocrmgroupdeleteAPIRequest) SetGroupId(_groupId int64) error {
	r._groupId = _groupId
	r.Set("group_id", _groupId)
	return nil
}

// GetGroupId GroupId Getter
func (r TaobaocrmgroupdeleteAPIRequest) GetGroupId() int64 {
	return r._groupId
}
