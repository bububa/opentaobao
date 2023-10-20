package jms

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaojushitajmsuseraddAPIRequest 添加ONS消息同步用户 API请求
// taobao.jushita.jms.user.add
//
// 添加ONS消息同步用户
type TaobaojushitajmsuseraddAPIRequest struct {
	model.Params
	// topic列表,不填则继承appkey所订阅的topic
	_topicNames []string
}

// NewTaobaojushitajmsuseraddRequest 初始化TaobaojushitajmsuseraddAPIRequest对象
func NewTaobaojushitajmsuseraddRequest() *TaobaojushitajmsuseraddAPIRequest {
	return &TaobaojushitajmsuseraddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaojushitajmsuseraddAPIRequest) GetApiMethodName() string {
	return "taobao.jushita.jms.user.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaojushitajmsuseraddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaojushitajmsuseraddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopicNames is TopicNames Setter
// topic列表,不填则继承appkey所订阅的topic
func (r *TaobaojushitajmsuseraddAPIRequest) SetTopicNames(_topicNames []string) error {
	r._topicNames = _topicNames
	r.Set("topic_names", _topicNames)
	return nil
}

// GetTopicNames TopicNames Getter
func (r TaobaojushitajmsuseraddAPIRequest) GetTopicNames() []string {
	return r._topicNames
}
