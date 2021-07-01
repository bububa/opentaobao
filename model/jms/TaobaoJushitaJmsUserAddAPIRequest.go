package jms

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJushitaJmsUserAddAPIRequest
添加ONS消息同步用户 API请求
taobao.jushita.jms.user.add

添加ONS消息同步用户 */
type TaobaoJushitaJmsUserAddAPIRequest struct {
	model.Params
	// topic列表,不填则继承appkey所订阅的topic
	_topicNames []string
}

// NewTaobaoJushitaJmsUserAddRequest 初始化TaobaoJushitaJmsUserAddAPIRequest对象
func NewTaobaoJushitaJmsUserAddRequest() *TaobaoJushitaJmsUserAddAPIRequest {
	return &TaobaoJushitaJmsUserAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJushitaJmsUserAddAPIRequest) GetApiMethodName() string {
	return "taobao.jushita.jms.user.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJushitaJmsUserAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is TopicNames Setter
// topic列表,不填则继承appkey所订阅的topic
func (r *TaobaoJushitaJmsUserAddAPIRequest) SetTopicNames(_topicNames []string) error {
	r._topicNames = _topicNames
	r.Set("topic_names", _topicNames)
	return nil
}

// Get TopicNames Getter
func (r TaobaoJushitaJmsUserAddAPIRequest) GetTopicNames() []string {
	return r._topicNames
}
