package crm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmGroupUpdateAPIRequest 修改一个已经存在的分组 API请求
// taobao.crm.group.update
//
// 修改一个已经存在的分组，接口返回分组的修改是否成功
type TaobaoCrmGroupUpdateAPIRequest struct {
	model.Params
	// 分组的id
	_groupId int64
	// 新的分组名，分组名称不能包含|或者：
	_newGroupName string
}

// NewTaobaoCrmGroupUpdateRequest 初始化TaobaoCrmGroupUpdateAPIRequest对象
func NewTaobaoCrmGroupUpdateRequest() *TaobaoCrmGroupUpdateAPIRequest {
	return &TaobaoCrmGroupUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCrmGroupUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.crm.group.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCrmGroupUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetGroupId is GroupId Setter
// 分组的id
func (r *TaobaoCrmGroupUpdateAPIRequest) SetGroupId(_groupId int64) error {
	r._groupId = _groupId
	r.Set("group_id", _groupId)
	return nil
}

// GetGroupId GroupId Getter
func (r TaobaoCrmGroupUpdateAPIRequest) GetGroupId() int64 {
	return r._groupId
}

// SetNewGroupName is NewGroupName Setter
// 新的分组名，分组名称不能包含|或者：
func (r *TaobaoCrmGroupUpdateAPIRequest) SetNewGroupName(_newGroupName string) error {
	r._newGroupName = _newGroupName
	r.Set("new_group_name", _newGroupName)
	return nil
}

// GetNewGroupName NewGroupName Getter
func (r TaobaoCrmGroupUpdateAPIRequest) GetNewGroupName() string {
	return r._newGroupName
}
