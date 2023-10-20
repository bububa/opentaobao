package crm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaocrmgroupupdateAPIRequest 修改一个已经存在的分组 API请求
// taobao.crm.group.update
//
// 修改一个已经存在的分组，接口返回分组的修改是否成功
type TaobaocrmgroupupdateAPIRequest struct {
	model.Params
	// 新的分组名，分组名称不能包含|或者：
	_newGroupName string
	// 分组的id
	_groupId int64
}

// NewTaobaocrmgroupupdateRequest 初始化TaobaocrmgroupupdateAPIRequest对象
func NewTaobaocrmgroupupdateRequest() *TaobaocrmgroupupdateAPIRequest {
	return &TaobaocrmgroupupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaocrmgroupupdateAPIRequest) GetApiMethodName() string {
	return "taobao.crm.group.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaocrmgroupupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaocrmgroupupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNewGroupName is NewGroupName Setter
// 新的分组名，分组名称不能包含|或者：
func (r *TaobaocrmgroupupdateAPIRequest) SetNewGroupName(_newGroupName string) error {
	r._newGroupName = _newGroupName
	r.Set("new_group_name", _newGroupName)
	return nil
}

// GetNewGroupName NewGroupName Getter
func (r TaobaocrmgroupupdateAPIRequest) GetNewGroupName() string {
	return r._newGroupName
}

// SetGroupId is GroupId Setter
// 分组的id
func (r *TaobaocrmgroupupdateAPIRequest) SetGroupId(_groupId int64) error {
	r._groupId = _groupId
	r.Set("group_id", _groupId)
	return nil
}

// GetGroupId GroupId Getter
func (r TaobaocrmgroupupdateAPIRequest) GetGroupId() int64 {
	return r._groupId
}
