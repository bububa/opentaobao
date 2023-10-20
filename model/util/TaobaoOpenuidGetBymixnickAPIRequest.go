package util

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenuidGetBymixnickAPIRequest 通过mixnick转换openuid API请求
// taobao.openuid.get.bymixnick
//
// 通过mixnick转换openuid
type TaobaoOpenuidGetBymixnickAPIRequest struct {
	model.Params
	// 无线类应用获取到的混淆的nick
	_mixNick string
}

// NewTaobaoOpenuidGetBymixnickRequest 初始化TaobaoOpenuidGetBymixnickAPIRequest对象
func NewTaobaoOpenuidGetBymixnickRequest() *TaobaoOpenuidGetBymixnickAPIRequest {
	return &TaobaoOpenuidGetBymixnickAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOpenuidGetBymixnickAPIRequest) Reset() {
	r._mixNick = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenuidGetBymixnickAPIRequest) GetApiMethodName() string {
	return "taobao.openuid.get.bymixnick"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenuidGetBymixnickAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpenuidGetBymixnickAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMixNick is MixNick Setter
// 无线类应用获取到的混淆的nick
func (r *TaobaoOpenuidGetBymixnickAPIRequest) SetMixNick(_mixNick string) error {
	r._mixNick = _mixNick
	r.Set("mix_nick", _mixNick)
	return nil
}

// GetMixNick MixNick Getter
func (r TaobaoOpenuidGetBymixnickAPIRequest) GetMixNick() string {
	return r._mixNick
}

var poolTaobaoOpenuidGetBymixnickAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOpenuidGetBymixnickRequest()
	},
}

// GetTaobaoOpenuidGetBymixnickRequest 从 sync.Pool 获取 TaobaoOpenuidGetBymixnickAPIRequest
func GetTaobaoOpenuidGetBymixnickAPIRequest() *TaobaoOpenuidGetBymixnickAPIRequest {
	return poolTaobaoOpenuidGetBymixnickAPIRequest.Get().(*TaobaoOpenuidGetBymixnickAPIRequest)
}

// ReleaseTaobaoOpenuidGetBymixnickAPIRequest 将 TaobaoOpenuidGetBymixnickAPIRequest 放入 sync.Pool
func ReleaseTaobaoOpenuidGetBymixnickAPIRequest(v *TaobaoOpenuidGetBymixnickAPIRequest) {
	v.Reset()
	poolTaobaoOpenuidGetBymixnickAPIRequest.Put(v)
}
