package crm

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmGroupUpdateAPIRequest 修改一个已经存在的分组 API请求
// taobao.crm.group.update
//
// 修改一个已经存在的分组，接口返回分组的修改是否成功
type TaobaoCrmGroupUpdateAPIRequest struct {
	model.Params
	// 新的分组名，分组名称不能包含|或者：
	_newGroupName string
	// 分组的id
	_groupId int64
}

// NewTaobaoCrmGroupUpdateRequest 初始化TaobaoCrmGroupUpdateAPIRequest对象
func NewTaobaoCrmGroupUpdateRequest() *TaobaoCrmGroupUpdateAPIRequest {
	return &TaobaoCrmGroupUpdateAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoCrmGroupUpdateAPIRequest) Reset() {
	r._newGroupName = ""
	r._groupId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCrmGroupUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.crm.group.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCrmGroupUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoCrmGroupUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTaobaoCrmGroupUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoCrmGroupUpdateRequest()
	},
}

// GetTaobaoCrmGroupUpdateRequest 从 sync.Pool 获取 TaobaoCrmGroupUpdateAPIRequest
func GetTaobaoCrmGroupUpdateAPIRequest() *TaobaoCrmGroupUpdateAPIRequest {
	return poolTaobaoCrmGroupUpdateAPIRequest.Get().(*TaobaoCrmGroupUpdateAPIRequest)
}

// ReleaseTaobaoCrmGroupUpdateAPIRequest 将 TaobaoCrmGroupUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoCrmGroupUpdateAPIRequest(v *TaobaoCrmGroupUpdateAPIRequest) {
	v.Reset()
	poolTaobaoCrmGroupUpdateAPIRequest.Put(v)
}
