package maitix

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMaitixOrderQueryAPIRequest 大麦-查询分销单 API请求
// alibaba.damai.maitix.order.query
//
// 查询分销单
type AlibabaDamaiMaitixOrderQueryAPIRequest struct {
	model.Params
	// 分销单查询入参
	_param *MoaOrderQueryParam
}

// NewAlibabaDamaiMaitixOrderQueryRequest 初始化AlibabaDamaiMaitixOrderQueryAPIRequest对象
func NewAlibabaDamaiMaitixOrderQueryRequest() *AlibabaDamaiMaitixOrderQueryAPIRequest {
	return &AlibabaDamaiMaitixOrderQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDamaiMaitixOrderQueryAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMaitixOrderQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.maitix.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMaitixOrderQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDamaiMaitixOrderQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 分销单查询入参
func (r *AlibabaDamaiMaitixOrderQueryAPIRequest) SetParam(_param *MoaOrderQueryParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaDamaiMaitixOrderQueryAPIRequest) GetParam() *MoaOrderQueryParam {
	return r._param
}

var poolAlibabaDamaiMaitixOrderQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDamaiMaitixOrderQueryRequest()
	},
}

// GetAlibabaDamaiMaitixOrderQueryRequest 从 sync.Pool 获取 AlibabaDamaiMaitixOrderQueryAPIRequest
func GetAlibabaDamaiMaitixOrderQueryAPIRequest() *AlibabaDamaiMaitixOrderQueryAPIRequest {
	return poolAlibabaDamaiMaitixOrderQueryAPIRequest.Get().(*AlibabaDamaiMaitixOrderQueryAPIRequest)
}

// ReleaseAlibabaDamaiMaitixOrderQueryAPIRequest 将 AlibabaDamaiMaitixOrderQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaDamaiMaitixOrderQueryAPIRequest(v *AlibabaDamaiMaitixOrderQueryAPIRequest) {
	v.Reset()
	poolAlibabaDamaiMaitixOrderQueryAPIRequest.Put(v)
}
