package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTopOpenidConvertAPIRequest 混淆nick转openid API请求
// taobao.top.openid.convert
//
// 混淆nick转openid，生成混淆nick必须与当前请求的isv匹配
type TaobaoTopOpenidConvertAPIRequest struct {
	model.Params
	// 混淆nick转open_id
	_mixNick string
}

// NewTaobaoTopOpenidConvertRequest 初始化TaobaoTopOpenidConvertAPIRequest对象
func NewTaobaoTopOpenidConvertRequest() *TaobaoTopOpenidConvertAPIRequest {
	return &TaobaoTopOpenidConvertAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTopOpenidConvertAPIRequest) GetApiMethodName() string {
	return "taobao.top.openid.convert"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTopOpenidConvertAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetMixNick is MixNick Setter
// 混淆nick转open_id
func (r *TaobaoTopOpenidConvertAPIRequest) SetMixNick(_mixNick string) error {
	r._mixNick = _mixNick
	r.Set("mix_nick", _mixNick)
	return nil
}

// GetMixNick MixNick Getter
func (r TaobaoTopOpenidConvertAPIRequest) GetMixNick() string {
	return r._mixNick
}
