package interact

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMixnickPlaytoweAPIRequest 互动mixNick转微淘 API请求
// taobao.mixnick.playtowe
//
// 微淘应用的混淆nick转为互动类型混淆nick
type TaobaoMixnickPlaytoweAPIRequest struct {
	model.Params
	// 用户的混淆nick
	_mixMix string
}

// NewTaobaoMixnickPlaytoweRequest 初始化TaobaoMixnickPlaytoweAPIRequest对象
func NewTaobaoMixnickPlaytoweRequest() *TaobaoMixnickPlaytoweAPIRequest {
	return &TaobaoMixnickPlaytoweAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoMixnickPlaytoweAPIRequest) Reset() {
	r._mixMix = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMixnickPlaytoweAPIRequest) GetApiMethodName() string {
	return "taobao.mixnick.playtowe"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMixnickPlaytoweAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoMixnickPlaytoweAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMixMix is MixMix Setter
// 用户的混淆nick
func (r *TaobaoMixnickPlaytoweAPIRequest) SetMixMix(_mixMix string) error {
	r._mixMix = _mixMix
	r.Set("mix_mix", _mixMix)
	return nil
}

// GetMixMix MixMix Getter
func (r TaobaoMixnickPlaytoweAPIRequest) GetMixMix() string {
	return r._mixMix
}

var poolTaobaoMixnickPlaytoweAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoMixnickPlaytoweRequest()
	},
}

// GetTaobaoMixnickPlaytoweRequest 从 sync.Pool 获取 TaobaoMixnickPlaytoweAPIRequest
func GetTaobaoMixnickPlaytoweAPIRequest() *TaobaoMixnickPlaytoweAPIRequest {
	return poolTaobaoMixnickPlaytoweAPIRequest.Get().(*TaobaoMixnickPlaytoweAPIRequest)
}

// ReleaseTaobaoMixnickPlaytoweAPIRequest 将 TaobaoMixnickPlaytoweAPIRequest 放入 sync.Pool
func ReleaseTaobaoMixnickPlaytoweAPIRequest(v *TaobaoMixnickPlaytoweAPIRequest) {
	v.Reset()
	poolTaobaoMixnickPlaytoweAPIRequest.Put(v)
}
