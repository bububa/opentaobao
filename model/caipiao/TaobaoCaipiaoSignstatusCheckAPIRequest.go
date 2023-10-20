package caipiao

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCaipiaoSignstatusCheckAPIRequest 检查用户是否签署支付宝代购协议 API请求
// taobao.caipiao.signstatus.check
//
// 检查用户是否签署了支付宝代扣协议。如果签署了，返回true; 如果没签署，返回false, 同时返回签署代扣协议的Url。
type TaobaoCaipiaoSignstatusCheckAPIRequest struct {
	model.Params
}

// NewTaobaoCaipiaoSignstatusCheckRequest 初始化TaobaoCaipiaoSignstatusCheckAPIRequest对象
func NewTaobaoCaipiaoSignstatusCheckRequest() *TaobaoCaipiaoSignstatusCheckAPIRequest {
	return &TaobaoCaipiaoSignstatusCheckAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoCaipiaoSignstatusCheckAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCaipiaoSignstatusCheckAPIRequest) GetApiMethodName() string {
	return "taobao.caipiao.signstatus.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCaipiaoSignstatusCheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoCaipiaoSignstatusCheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolTaobaoCaipiaoSignstatusCheckAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoCaipiaoSignstatusCheckRequest()
	},
}

// GetTaobaoCaipiaoSignstatusCheckRequest 从 sync.Pool 获取 TaobaoCaipiaoSignstatusCheckAPIRequest
func GetTaobaoCaipiaoSignstatusCheckAPIRequest() *TaobaoCaipiaoSignstatusCheckAPIRequest {
	return poolTaobaoCaipiaoSignstatusCheckAPIRequest.Get().(*TaobaoCaipiaoSignstatusCheckAPIRequest)
}

// ReleaseTaobaoCaipiaoSignstatusCheckAPIRequest 将 TaobaoCaipiaoSignstatusCheckAPIRequest 放入 sync.Pool
func ReleaseTaobaoCaipiaoSignstatusCheckAPIRequest(v *TaobaoCaipiaoSignstatusCheckAPIRequest) {
	v.Reset()
	poolTaobaoCaipiaoSignstatusCheckAPIRequest.Put(v)
}
