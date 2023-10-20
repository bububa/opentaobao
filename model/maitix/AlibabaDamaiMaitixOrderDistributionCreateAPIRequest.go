package maitix

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMaitixOrderDistributionCreateAPIRequest 大麦-新分销下单 API请求
// alibaba.damai.maitix.order.distribution.create
//
// createDistributionOrder
type AlibabaDamaiMaitixOrderDistributionCreateAPIRequest struct {
	model.Params
	// 下单参数param
	_param *MoaOrderParam
}

// NewAlibabaDamaiMaitixOrderDistributionCreateRequest 初始化AlibabaDamaiMaitixOrderDistributionCreateAPIRequest对象
func NewAlibabaDamaiMaitixOrderDistributionCreateRequest() *AlibabaDamaiMaitixOrderDistributionCreateAPIRequest {
	return &AlibabaDamaiMaitixOrderDistributionCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDamaiMaitixOrderDistributionCreateAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMaitixOrderDistributionCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.maitix.order.distribution.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMaitixOrderDistributionCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDamaiMaitixOrderDistributionCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 下单参数param
func (r *AlibabaDamaiMaitixOrderDistributionCreateAPIRequest) SetParam(_param *MoaOrderParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaDamaiMaitixOrderDistributionCreateAPIRequest) GetParam() *MoaOrderParam {
	return r._param
}

var poolAlibabaDamaiMaitixOrderDistributionCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDamaiMaitixOrderDistributionCreateRequest()
	},
}

// GetAlibabaDamaiMaitixOrderDistributionCreateRequest 从 sync.Pool 获取 AlibabaDamaiMaitixOrderDistributionCreateAPIRequest
func GetAlibabaDamaiMaitixOrderDistributionCreateAPIRequest() *AlibabaDamaiMaitixOrderDistributionCreateAPIRequest {
	return poolAlibabaDamaiMaitixOrderDistributionCreateAPIRequest.Get().(*AlibabaDamaiMaitixOrderDistributionCreateAPIRequest)
}

// ReleaseAlibabaDamaiMaitixOrderDistributionCreateAPIRequest 将 AlibabaDamaiMaitixOrderDistributionCreateAPIRequest 放入 sync.Pool
func ReleaseAlibabaDamaiMaitixOrderDistributionCreateAPIRequest(v *AlibabaDamaiMaitixOrderDistributionCreateAPIRequest) {
	v.Reset()
	poolAlibabaDamaiMaitixOrderDistributionCreateAPIRequest.Put(v)
}
