package util

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTimeGetAPIRequest 获取淘宝系统当前时间 API请求
// taobao.time.get
//
// 获取淘宝系统当前时间
type TaobaoTimeGetAPIRequest struct {
	model.Params
}

// NewTaobaoTimeGetRequest 初始化TaobaoTimeGetAPIRequest对象
func NewTaobaoTimeGetRequest() *TaobaoTimeGetAPIRequest {
	return &TaobaoTimeGetAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTimeGetAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTimeGetAPIRequest) GetApiMethodName() string {
	return "taobao.time.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTimeGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTimeGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolTaobaoTimeGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTimeGetRequest()
	},
}

// GetTaobaoTimeGetRequest 从 sync.Pool 获取 TaobaoTimeGetAPIRequest
func GetTaobaoTimeGetAPIRequest() *TaobaoTimeGetAPIRequest {
	return poolTaobaoTimeGetAPIRequest.Get().(*TaobaoTimeGetAPIRequest)
}

// ReleaseTaobaoTimeGetAPIRequest 将 TaobaoTimeGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoTimeGetAPIRequest(v *TaobaoTimeGetAPIRequest) {
	v.Reset()
	poolTaobaoTimeGetAPIRequest.Put(v)
}
