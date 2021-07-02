package tmc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTmcTopicGroupAddAPIRequest topic分组路由 API请求
// taobao.tmc.topic.group.add
//
// 根据topic名称路由消息到不同的分组。（前提：发送方未指定分组名）
// 如果是需要授权的消息，分组路由先判断用户分组路由(使用taobao.tmc.group.add添加的路由)，用户分组路由不存在时，才会判断topic分组路由
type TaobaoTmcTopicGroupAddAPIRequest struct {
	model.Params
	// 消息分组名，如果不存在，会自动创建
	_groupName string
	// 消息topic名称，多个以逗号(,)分割
	_topics []string
}

// NewTaobaoTmcTopicGroupAddRequest 初始化TaobaoTmcTopicGroupAddAPIRequest对象
func NewTaobaoTmcTopicGroupAddRequest() *TaobaoTmcTopicGroupAddAPIRequest {
	return &TaobaoTmcTopicGroupAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTmcTopicGroupAddAPIRequest) GetApiMethodName() string {
	return "taobao.tmc.topic.group.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTmcTopicGroupAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is GroupName Setter
// 消息分组名，如果不存在，会自动创建
func (r *TaobaoTmcTopicGroupAddAPIRequest) SetGroupName(_groupName string) error {
	r._groupName = _groupName
	r.Set("group_name", _groupName)
	return nil
}

// Get GroupName Getter
func (r TaobaoTmcTopicGroupAddAPIRequest) GetGroupName() string {
	return r._groupName
}

// Set is Topics Setter
// 消息topic名称，多个以逗号(,)分割
func (r *TaobaoTmcTopicGroupAddAPIRequest) SetTopics(_topics []string) error {
	r._topics = _topics
	r.Set("topics", _topics)
	return nil
}

// Get Topics Getter
func (r TaobaoTmcTopicGroupAddAPIRequest) GetTopics() []string {
	return r._topics
}
