package jms

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJushitaJmsTopicsGetAPIRequest 根据用户nick获取开通的消息列表 API请求
// taobao.jushita.jms.topics.get
//
// 根据用户nick获取开通的消息列表
type TaobaoJushitaJmsTopicsGetAPIRequest struct {
	model.Params
	// 卖家nick
	_nick string
}

// NewTaobaoJushitaJmsTopicsGetRequest 初始化TaobaoJushitaJmsTopicsGetAPIRequest对象
func NewTaobaoJushitaJmsTopicsGetRequest() *TaobaoJushitaJmsTopicsGetAPIRequest {
	return &TaobaoJushitaJmsTopicsGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJushitaJmsTopicsGetAPIRequest) GetApiMethodName() string {
	return "taobao.jushita.jms.topics.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJushitaJmsTopicsGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoJushitaJmsTopicsGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 卖家nick
func (r *TaobaoJushitaJmsTopicsGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoJushitaJmsTopicsGetAPIRequest) GetNick() string {
	return r._nick
}
