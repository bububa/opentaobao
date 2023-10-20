package tmc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotmctopicgroupaddAPIRequest topic分组路由 API请求
// taobao.tmc.topic.group.add
//
// 根据topic名称路由消息到不同的分组。（前提：发送方未指定分组名）
// 如果是需要授权的消息，分组路由先判断用户分组路由(使用taobao.tmc.group.add添加的路由)，用户分组路由不存在时，才会判断topic分组路由
type TaobaotmctopicgroupaddAPIRequest struct {
	model.Params
	// 消息topic名称，多个以逗号(,)分割
	_topics []string
	// 消息分组名，如果不存在，会自动创建
	_groupName string
}

// NewTaobaotmctopicgroupaddRequest 初始化TaobaotmctopicgroupaddAPIRequest对象
func NewTaobaotmctopicgroupaddRequest() *TaobaotmctopicgroupaddAPIRequest {
	return &TaobaotmctopicgroupaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotmctopicgroupaddAPIRequest) GetApiMethodName() string {
	return "taobao.tmc.topic.group.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotmctopicgroupaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotmctopicgroupaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopics is Topics Setter
// 消息topic名称，多个以逗号(,)分割
func (r *TaobaotmctopicgroupaddAPIRequest) SetTopics(_topics []string) error {
	r._topics = _topics
	r.Set("topics", _topics)
	return nil
}

// GetTopics Topics Getter
func (r TaobaotmctopicgroupaddAPIRequest) GetTopics() []string {
	return r._topics
}

// SetGroupName is GroupName Setter
// 消息分组名，如果不存在，会自动创建
func (r *TaobaotmctopicgroupaddAPIRequest) SetGroupName(_groupName string) error {
	r._groupName = _groupName
	r.Set("group_name", _groupName)
	return nil
}

// GetGroupName GroupName Getter
func (r TaobaotmctopicgroupaddAPIRequest) GetGroupName() string {
	return r._groupName
}
