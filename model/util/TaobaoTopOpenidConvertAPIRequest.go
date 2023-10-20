package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotopopenidconvertAPIRequest 混淆nick转openid API请求
// taobao.top.openid.convert
//
// 混淆nick转openid，生成混淆nick必须与当前请求的isv匹配
type TaobaotopopenidconvertAPIRequest struct {
	model.Params
	// 混淆nick转open_id
	_mixNick string
}

// NewTaobaotopopenidconvertRequest 初始化TaobaotopopenidconvertAPIRequest对象
func NewTaobaotopopenidconvertRequest() *TaobaotopopenidconvertAPIRequest {
	return &TaobaotopopenidconvertAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotopopenidconvertAPIRequest) GetApiMethodName() string {
	return "taobao.top.openid.convert"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotopopenidconvertAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotopopenidconvertAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMixNick is MixNick Setter
// 混淆nick转open_id
func (r *TaobaotopopenidconvertAPIRequest) SetMixNick(_mixNick string) error {
	r._mixNick = _mixNick
	r.Set("mix_nick", _mixNick)
	return nil
}

// GetMixNick MixNick Getter
func (r TaobaotopopenidconvertAPIRequest) GetMixNick() string {
	return r._mixNick
}
