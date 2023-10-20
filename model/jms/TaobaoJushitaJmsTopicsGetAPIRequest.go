package jms

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaojushitajmstopicsgetAPIRequest 根据用户nick获取开通的消息列表 API请求
// taobao.jushita.jms.topics.get
//
// 根据用户nick获取开通的消息列表
type TaobaojushitajmstopicsgetAPIRequest struct {
	model.Params
	// 卖家nick
	_nick string
}

// NewTaobaojushitajmstopicsgetRequest 初始化TaobaojushitajmstopicsgetAPIRequest对象
func NewTaobaojushitajmstopicsgetRequest() *TaobaojushitajmstopicsgetAPIRequest {
	return &TaobaojushitajmstopicsgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaojushitajmstopicsgetAPIRequest) GetApiMethodName() string {
	return "taobao.jushita.jms.topics.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaojushitajmstopicsgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaojushitajmstopicsgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 卖家nick
func (r *TaobaojushitajmstopicsgetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaojushitajmstopicsgetAPIRequest) GetNick() string {
	return r._nick
}
