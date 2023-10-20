package tmc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTmcTopicGroupDeleteAPIRequest 删除消息topic分组路由 API请求
// taobao.tmc.topic.group.delete
//
// 删除根据topic名称路由消息到不同的分组关系
type TaobaoTmcTopicGroupDeleteAPIRequest struct {
	model.Params
	// 消息topic名称，多个以逗号(,)分割
	_topics []string
	// 消息分组名
	_groupName string
	// 消息分组Id，一般不用填写，如果分组已经被删除，则根据问题排查工具返回的ID删除路由关系
	_groupId int64
}

// NewTaobaoTmcTopicGroupDeleteRequest 初始化TaobaoTmcTopicGroupDeleteAPIRequest对象
func NewTaobaoTmcTopicGroupDeleteRequest() *TaobaoTmcTopicGroupDeleteAPIRequest {
	return &TaobaoTmcTopicGroupDeleteAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTmcTopicGroupDeleteAPIRequest) Reset() {
	r._topics = r._topics[:0]
	r._groupName = ""
	r._groupId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTmcTopicGroupDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.tmc.topic.group.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTmcTopicGroupDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTmcTopicGroupDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopics is Topics Setter
// 消息topic名称，多个以逗号(,)分割
func (r *TaobaoTmcTopicGroupDeleteAPIRequest) SetTopics(_topics []string) error {
	r._topics = _topics
	r.Set("topics", _topics)
	return nil
}

// GetTopics Topics Getter
func (r TaobaoTmcTopicGroupDeleteAPIRequest) GetTopics() []string {
	return r._topics
}

// SetGroupName is GroupName Setter
// 消息分组名
func (r *TaobaoTmcTopicGroupDeleteAPIRequest) SetGroupName(_groupName string) error {
	r._groupName = _groupName
	r.Set("group_name", _groupName)
	return nil
}

// GetGroupName GroupName Getter
func (r TaobaoTmcTopicGroupDeleteAPIRequest) GetGroupName() string {
	return r._groupName
}

// SetGroupId is GroupId Setter
// 消息分组Id，一般不用填写，如果分组已经被删除，则根据问题排查工具返回的ID删除路由关系
func (r *TaobaoTmcTopicGroupDeleteAPIRequest) SetGroupId(_groupId int64) error {
	r._groupId = _groupId
	r.Set("group_id", _groupId)
	return nil
}

// GetGroupId GroupId Getter
func (r TaobaoTmcTopicGroupDeleteAPIRequest) GetGroupId() int64 {
	return r._groupId
}

var poolTaobaoTmcTopicGroupDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTmcTopicGroupDeleteRequest()
	},
}

// GetTaobaoTmcTopicGroupDeleteRequest 从 sync.Pool 获取 TaobaoTmcTopicGroupDeleteAPIRequest
func GetTaobaoTmcTopicGroupDeleteAPIRequest() *TaobaoTmcTopicGroupDeleteAPIRequest {
	return poolTaobaoTmcTopicGroupDeleteAPIRequest.Get().(*TaobaoTmcTopicGroupDeleteAPIRequest)
}

// ReleaseTaobaoTmcTopicGroupDeleteAPIRequest 将 TaobaoTmcTopicGroupDeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoTmcTopicGroupDeleteAPIRequest(v *TaobaoTmcTopicGroupDeleteAPIRequest) {
	v.Reset()
	poolTaobaoTmcTopicGroupDeleteAPIRequest.Put(v)
}
