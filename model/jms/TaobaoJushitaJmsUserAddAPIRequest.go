package jms

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJushitaJmsUserAddAPIRequest 添加ONS消息同步用户 API请求
// taobao.jushita.jms.user.add
//
// 添加ONS消息同步用户
type TaobaoJushitaJmsUserAddAPIRequest struct {
	model.Params
	// topic列表,不填则继承appkey所订阅的topic
	_topicNames []string
}

// NewTaobaoJushitaJmsUserAddRequest 初始化TaobaoJushitaJmsUserAddAPIRequest对象
func NewTaobaoJushitaJmsUserAddRequest() *TaobaoJushitaJmsUserAddAPIRequest {
	return &TaobaoJushitaJmsUserAddAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoJushitaJmsUserAddAPIRequest) Reset() {
	r._topicNames = r._topicNames[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJushitaJmsUserAddAPIRequest) GetApiMethodName() string {
	return "taobao.jushita.jms.user.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJushitaJmsUserAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoJushitaJmsUserAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopicNames is TopicNames Setter
// topic列表,不填则继承appkey所订阅的topic
func (r *TaobaoJushitaJmsUserAddAPIRequest) SetTopicNames(_topicNames []string) error {
	r._topicNames = _topicNames
	r.Set("topic_names", _topicNames)
	return nil
}

// GetTopicNames TopicNames Getter
func (r TaobaoJushitaJmsUserAddAPIRequest) GetTopicNames() []string {
	return r._topicNames
}

var poolTaobaoJushitaJmsUserAddAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoJushitaJmsUserAddRequest()
	},
}

// GetTaobaoJushitaJmsUserAddRequest 从 sync.Pool 获取 TaobaoJushitaJmsUserAddAPIRequest
func GetTaobaoJushitaJmsUserAddAPIRequest() *TaobaoJushitaJmsUserAddAPIRequest {
	return poolTaobaoJushitaJmsUserAddAPIRequest.Get().(*TaobaoJushitaJmsUserAddAPIRequest)
}

// ReleaseTaobaoJushitaJmsUserAddAPIRequest 将 TaobaoJushitaJmsUserAddAPIRequest 放入 sync.Pool
func ReleaseTaobaoJushitaJmsUserAddAPIRequest(v *TaobaoJushitaJmsUserAddAPIRequest) {
	v.Reset()
	poolTaobaoJushitaJmsUserAddAPIRequest.Put(v)
}
