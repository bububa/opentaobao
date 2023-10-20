package crm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaocrmgroupmoveAPIRequest 分组移动 API请求
// taobao.crm.group.move
//
// 将一个分组下的所有会员移动到另一个分组，会员从原分组中删除&lt;br/&gt;注：移动属性为异步任务建议先调用taobao.crm.grouptask.check 确保涉及属性上没有任务。
type TaobaocrmgroupmoveAPIRequest struct {
	model.Params
	// 需要移动的分组
	_fromGroupId int64
	// 目的分组
	_toGroupId int64
}

// NewTaobaocrmgroupmoveRequest 初始化TaobaocrmgroupmoveAPIRequest对象
func NewTaobaocrmgroupmoveRequest() *TaobaocrmgroupmoveAPIRequest {
	return &TaobaocrmgroupmoveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaocrmgroupmoveAPIRequest) GetApiMethodName() string {
	return "taobao.crm.group.move"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaocrmgroupmoveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaocrmgroupmoveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFromGroupId is FromGroupId Setter
// 需要移动的分组
func (r *TaobaocrmgroupmoveAPIRequest) SetFromGroupId(_fromGroupId int64) error {
	r._fromGroupId = _fromGroupId
	r.Set("from_group_id", _fromGroupId)
	return nil
}

// GetFromGroupId FromGroupId Getter
func (r TaobaocrmgroupmoveAPIRequest) GetFromGroupId() int64 {
	return r._fromGroupId
}

// SetToGroupId is ToGroupId Setter
// 目的分组
func (r *TaobaocrmgroupmoveAPIRequest) SetToGroupId(_toGroupId int64) error {
	r._toGroupId = _toGroupId
	r.Set("to_group_id", _toGroupId)
	return nil
}

// GetToGroupId ToGroupId Getter
func (r TaobaocrmgroupmoveAPIRequest) GetToGroupId() int64 {
	return r._toGroupId
}
