package tmc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotmcusertopicsgetAPIRequest 获取用户开通的topic列表 API请求
// taobao.tmc.user.topics.get
//
// 获取用户开通的topic列表，授权消息使用
type TaobaotmcusertopicsgetAPIRequest struct {
	model.Params
	// 卖家nick
	_nick string
}

// NewTaobaotmcusertopicsgetRequest 初始化TaobaotmcusertopicsgetAPIRequest对象
func NewTaobaotmcusertopicsgetRequest() *TaobaotmcusertopicsgetAPIRequest {
	return &TaobaotmcusertopicsgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotmcusertopicsgetAPIRequest) GetApiMethodName() string {
	return "taobao.tmc.user.topics.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotmcusertopicsgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotmcusertopicsgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 卖家nick
func (r *TaobaotmcusertopicsgetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaotmcusertopicsgetAPIRequest) GetNick() string {
	return r._nick
}
