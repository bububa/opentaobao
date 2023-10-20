package util

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTopOpenidConvertAPIRequest) Reset() {
	r._mixNick = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTopOpenidConvertAPIRequest) GetApiMethodName() string {
	return "taobao.top.openid.convert"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTopOpenidConvertAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTopOpenidConvertAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTaobaoTopOpenidConvertAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTopOpenidConvertRequest()
	},
}

// GetTaobaoTopOpenidConvertRequest 从 sync.Pool 获取 TaobaoTopOpenidConvertAPIRequest
func GetTaobaoTopOpenidConvertAPIRequest() *TaobaoTopOpenidConvertAPIRequest {
	return poolTaobaoTopOpenidConvertAPIRequest.Get().(*TaobaoTopOpenidConvertAPIRequest)
}

// ReleaseTaobaoTopOpenidConvertAPIRequest 将 TaobaoTopOpenidConvertAPIRequest 放入 sync.Pool
func ReleaseTaobaoTopOpenidConvertAPIRequest(v *TaobaoTopOpenidConvertAPIRequest) {
	v.Reset()
	poolTaobaoTopOpenidConvertAPIRequest.Put(v)
}
