package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaomixnickplaytoweAPIRequest 互动mixNick转微淘 API请求
// taobao.mixnick.playtowe
//
// 微淘应用的混淆nick转为互动类型混淆nick
type TaobaomixnickplaytoweAPIRequest struct {
	model.Params
	// 用户的混淆nick
	_mixMix string
}

// NewTaobaomixnickplaytoweRequest 初始化TaobaomixnickplaytoweAPIRequest对象
func NewTaobaomixnickplaytoweRequest() *TaobaomixnickplaytoweAPIRequest {
	return &TaobaomixnickplaytoweAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaomixnickplaytoweAPIRequest) GetApiMethodName() string {
	return "taobao.mixnick.playtowe"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaomixnickplaytoweAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaomixnickplaytoweAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMixMix is MixMix Setter
// 用户的混淆nick
func (r *TaobaomixnickplaytoweAPIRequest) SetMixMix(_mixMix string) error {
	r._mixMix = _mixMix
	r.Set("mix_mix", _mixMix)
	return nil
}

// GetMixMix MixMix Getter
func (r TaobaomixnickplaytoweAPIRequest) GetMixMix() string {
	return r._mixMix
}
