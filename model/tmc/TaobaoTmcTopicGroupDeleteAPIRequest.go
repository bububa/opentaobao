package tmc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotmctopicgroupdeleteAPIRequest 删除消息topic分组路由 API请求
// taobao.tmc.topic.group.delete
//
// 删除根据topic名称路由消息到不同的分组关系
type TaobaotmctopicgroupdeleteAPIRequest struct {
	model.Params
	// 消息topic名称，多个以逗号(,)分割
	_topics []string
	// 消息分组名
	_groupName string
	// 消息分组Id，一般不用填写，如果分组已经被删除，则根据问题排查工具返回的ID删除路由关系
	_groupId int64
}

// NewTaobaotmctopicgroupdeleteRequest 初始化TaobaotmctopicgroupdeleteAPIRequest对象
func NewTaobaotmctopicgroupdeleteRequest() *TaobaotmctopicgroupdeleteAPIRequest {
	return &TaobaotmctopicgroupdeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotmctopicgroupdeleteAPIRequest) GetApiMethodName() string {
	return "taobao.tmc.topic.group.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotmctopicgroupdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotmctopicgroupdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopics is Topics Setter
// 消息topic名称，多个以逗号(,)分割
func (r *TaobaotmctopicgroupdeleteAPIRequest) SetTopics(_topics []string) error {
	r._topics = _topics
	r.Set("topics", _topics)
	return nil
}

// GetTopics Topics Getter
func (r TaobaotmctopicgroupdeleteAPIRequest) GetTopics() []string {
	return r._topics
}

// SetGroupName is GroupName Setter
// 消息分组名
func (r *TaobaotmctopicgroupdeleteAPIRequest) SetGroupName(_groupName string) error {
	r._groupName = _groupName
	r.Set("group_name", _groupName)
	return nil
}

// GetGroupName GroupName Getter
func (r TaobaotmctopicgroupdeleteAPIRequest) GetGroupName() string {
	return r._groupName
}

// SetGroupId is GroupId Setter
// 消息分组Id，一般不用填写，如果分组已经被删除，则根据问题排查工具返回的ID删除路由关系
func (r *TaobaotmctopicgroupdeleteAPIRequest) SetGroupId(_groupId int64) error {
	r._groupId = _groupId
	r.Set("group_id", _groupId)
	return nil
}

// GetGroupId GroupId Getter
func (r TaobaotmctopicgroupdeleteAPIRequest) GetGroupId() int64 {
	return r._groupId
}
