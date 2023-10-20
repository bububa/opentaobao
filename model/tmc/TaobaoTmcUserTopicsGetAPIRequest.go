package tmc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTmcUserTopicsGetAPIRequest 获取用户开通的topic列表 API请求
// taobao.tmc.user.topics.get
//
// 获取用户开通的topic列表，授权消息使用
type TaobaoTmcUserTopicsGetAPIRequest struct {
	model.Params
	// 卖家nick
	_nick string
}

// NewTaobaoTmcUserTopicsGetRequest 初始化TaobaoTmcUserTopicsGetAPIRequest对象
func NewTaobaoTmcUserTopicsGetRequest() *TaobaoTmcUserTopicsGetAPIRequest {
	return &TaobaoTmcUserTopicsGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTmcUserTopicsGetAPIRequest) Reset() {
	r._nick = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTmcUserTopicsGetAPIRequest) GetApiMethodName() string {
	return "taobao.tmc.user.topics.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTmcUserTopicsGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTmcUserTopicsGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 卖家nick
func (r *TaobaoTmcUserTopicsGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoTmcUserTopicsGetAPIRequest) GetNick() string {
	return r._nick
}

var poolTaobaoTmcUserTopicsGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTmcUserTopicsGetRequest()
	},
}

// GetTaobaoTmcUserTopicsGetRequest 从 sync.Pool 获取 TaobaoTmcUserTopicsGetAPIRequest
func GetTaobaoTmcUserTopicsGetAPIRequest() *TaobaoTmcUserTopicsGetAPIRequest {
	return poolTaobaoTmcUserTopicsGetAPIRequest.Get().(*TaobaoTmcUserTopicsGetAPIRequest)
}

// ReleaseTaobaoTmcUserTopicsGetAPIRequest 将 TaobaoTmcUserTopicsGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoTmcUserTopicsGetAPIRequest(v *TaobaoTmcUserTopicsGetAPIRequest) {
	v.Reset()
	poolTaobaoTmcUserTopicsGetAPIRequest.Put(v)
}
