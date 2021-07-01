package tmc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTmcUserTopicsGetAPIRequest
获取用户开通的topic列表 API请求
taobao.tmc.user.topics.get

获取用户开通的topic列表，授权消息使用 */
type TaobaoTmcUserTopicsGetAPIRequest struct {
	model.Params
	// 卖家nick
	_nick string
}

// NewTaobaoTmcUserTopicsGetRequest 初始化TaobaoTmcUserTopicsGetAPIRequest对象
func NewTaobaoTmcUserTopicsGetRequest() *TaobaoTmcUserTopicsGetAPIRequest {
	return &TaobaoTmcUserTopicsGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTmcUserTopicsGetAPIRequest) GetApiMethodName() string {
	return "taobao.tmc.user.topics.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTmcUserTopicsGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Nick Setter
// 卖家nick
func (r *TaobaoTmcUserTopicsGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// Get Nick Getter
func (r TaobaoTmcUserTopicsGetAPIRequest) GetNick() string {
	return r._nick
}
