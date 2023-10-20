package maitix

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMaitixOrderConfirmAPIRequest 大麦-出票 API请求
// alibaba.damai.maitix.order.confirm
//
// 出票
type AlibabaDamaiMaitixOrderConfirmAPIRequest struct {
	model.Params
	// 出票入参
	_param *MoaConfirmOrderParam
}

// NewAlibabaDamaiMaitixOrderConfirmRequest 初始化AlibabaDamaiMaitixOrderConfirmAPIRequest对象
func NewAlibabaDamaiMaitixOrderConfirmRequest() *AlibabaDamaiMaitixOrderConfirmAPIRequest {
	return &AlibabaDamaiMaitixOrderConfirmAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDamaiMaitixOrderConfirmAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMaitixOrderConfirmAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.maitix.order.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMaitixOrderConfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDamaiMaitixOrderConfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 出票入参
func (r *AlibabaDamaiMaitixOrderConfirmAPIRequest) SetParam(_param *MoaConfirmOrderParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaDamaiMaitixOrderConfirmAPIRequest) GetParam() *MoaConfirmOrderParam {
	return r._param
}

var poolAlibabaDamaiMaitixOrderConfirmAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDamaiMaitixOrderConfirmRequest()
	},
}

// GetAlibabaDamaiMaitixOrderConfirmRequest 从 sync.Pool 获取 AlibabaDamaiMaitixOrderConfirmAPIRequest
func GetAlibabaDamaiMaitixOrderConfirmAPIRequest() *AlibabaDamaiMaitixOrderConfirmAPIRequest {
	return poolAlibabaDamaiMaitixOrderConfirmAPIRequest.Get().(*AlibabaDamaiMaitixOrderConfirmAPIRequest)
}

// ReleaseAlibabaDamaiMaitixOrderConfirmAPIRequest 将 AlibabaDamaiMaitixOrderConfirmAPIRequest 放入 sync.Pool
func ReleaseAlibabaDamaiMaitixOrderConfirmAPIRequest(v *AlibabaDamaiMaitixOrderConfirmAPIRequest) {
	v.Reset()
	poolAlibabaDamaiMaitixOrderConfirmAPIRequest.Put(v)
}
