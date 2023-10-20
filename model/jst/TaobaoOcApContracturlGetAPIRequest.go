package jst

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOcApContracturlGetAPIRequest 按用户获取支付宝代扣协议链接地址 API请求
// taobao.oc.ap.contracturl.get
//
// 按用户获取支付宝代扣协议链接地址
type TaobaoOcApContracturlGetAPIRequest struct {
	model.Params
}

// NewTaobaoOcApContracturlGetRequest 初始化TaobaoOcApContracturlGetAPIRequest对象
func NewTaobaoOcApContracturlGetRequest() *TaobaoOcApContracturlGetAPIRequest {
	return &TaobaoOcApContracturlGetAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOcApContracturlGetAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOcApContracturlGetAPIRequest) GetApiMethodName() string {
	return "taobao.oc.ap.contracturl.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOcApContracturlGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOcApContracturlGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolTaobaoOcApContracturlGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOcApContracturlGetRequest()
	},
}

// GetTaobaoOcApContracturlGetRequest 从 sync.Pool 获取 TaobaoOcApContracturlGetAPIRequest
func GetTaobaoOcApContracturlGetAPIRequest() *TaobaoOcApContracturlGetAPIRequest {
	return poolTaobaoOcApContracturlGetAPIRequest.Get().(*TaobaoOcApContracturlGetAPIRequest)
}

// ReleaseTaobaoOcApContracturlGetAPIRequest 将 TaobaoOcApContracturlGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoOcApContracturlGetAPIRequest(v *TaobaoOcApContracturlGetAPIRequest) {
	v.Reset()
	poolTaobaoOcApContracturlGetAPIRequest.Put(v)
}
