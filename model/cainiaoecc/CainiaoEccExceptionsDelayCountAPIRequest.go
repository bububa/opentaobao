package cainiaoecc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoEccExceptionsDelayCountAPIRequest 菜鸟控制塔包裹滞留异常统计信息获取 API请求
// cainiao.ecc.exceptions.delay.count
//
// 菜鸟控制塔包裹滞留异常统计信息获取
type CainiaoEccExceptionsDelayCountAPIRequest struct {
	model.Params
}

// NewCainiaoEccExceptionsDelayCountRequest 初始化CainiaoEccExceptionsDelayCountAPIRequest对象
func NewCainiaoEccExceptionsDelayCountRequest() *CainiaoEccExceptionsDelayCountAPIRequest {
	return &CainiaoEccExceptionsDelayCountAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoEccExceptionsDelayCountAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoEccExceptionsDelayCountAPIRequest) GetApiMethodName() string {
	return "cainiao.ecc.exceptions.delay.count"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoEccExceptionsDelayCountAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoEccExceptionsDelayCountAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolCainiaoEccExceptionsDelayCountAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoEccExceptionsDelayCountRequest()
	},
}

// GetCainiaoEccExceptionsDelayCountRequest 从 sync.Pool 获取 CainiaoEccExceptionsDelayCountAPIRequest
func GetCainiaoEccExceptionsDelayCountAPIRequest() *CainiaoEccExceptionsDelayCountAPIRequest {
	return poolCainiaoEccExceptionsDelayCountAPIRequest.Get().(*CainiaoEccExceptionsDelayCountAPIRequest)
}

// ReleaseCainiaoEccExceptionsDelayCountAPIRequest 将 CainiaoEccExceptionsDelayCountAPIRequest 放入 sync.Pool
func ReleaseCainiaoEccExceptionsDelayCountAPIRequest(v *CainiaoEccExceptionsDelayCountAPIRequest) {
	v.Reset()
	poolCainiaoEccExceptionsDelayCountAPIRequest.Put(v)
}
