package sungari

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSungariDisposeSubmitAPIRequest 商品商家处置提交任务 API请求
// taobao.sungari.dispose.submit
//
// 商品商家处置信息接口，提供政府部门发送处置信息给阿里
type TaobaoSungariDisposeSubmitAPIRequest struct {
	model.Params
	// 平台处置信息入参
	_info *DisposeInfoDo
}

// NewTaobaoSungariDisposeSubmitRequest 初始化TaobaoSungariDisposeSubmitAPIRequest对象
func NewTaobaoSungariDisposeSubmitRequest() *TaobaoSungariDisposeSubmitAPIRequest {
	return &TaobaoSungariDisposeSubmitAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSungariDisposeSubmitAPIRequest) Reset() {
	r._info = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSungariDisposeSubmitAPIRequest) GetApiMethodName() string {
	return "taobao.sungari.dispose.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSungariDisposeSubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSungariDisposeSubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInfo is Info Setter
// 平台处置信息入参
func (r *TaobaoSungariDisposeSubmitAPIRequest) SetInfo(_info *DisposeInfoDo) error {
	r._info = _info
	r.Set("info", _info)
	return nil
}

// GetInfo Info Getter
func (r TaobaoSungariDisposeSubmitAPIRequest) GetInfo() *DisposeInfoDo {
	return r._info
}

var poolTaobaoSungariDisposeSubmitAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSungariDisposeSubmitRequest()
	},
}

// GetTaobaoSungariDisposeSubmitRequest 从 sync.Pool 获取 TaobaoSungariDisposeSubmitAPIRequest
func GetTaobaoSungariDisposeSubmitAPIRequest() *TaobaoSungariDisposeSubmitAPIRequest {
	return poolTaobaoSungariDisposeSubmitAPIRequest.Get().(*TaobaoSungariDisposeSubmitAPIRequest)
}

// ReleaseTaobaoSungariDisposeSubmitAPIRequest 将 TaobaoSungariDisposeSubmitAPIRequest 放入 sync.Pool
func ReleaseTaobaoSungariDisposeSubmitAPIRequest(v *TaobaoSungariDisposeSubmitAPIRequest) {
	v.Reset()
	poolTaobaoSungariDisposeSubmitAPIRequest.Put(v)
}
