package maitix

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMaitixEticketDistributionQueryAPIRequest 分销电子票查询接口 API请求
// alibaba.damai.maitix.eticket.distribution.query
//
// 分销电子票查询接口
type AlibabaDamaiMaitixEticketDistributionQueryAPIRequest struct {
	model.Params
	// 入参param
	_param *EticketQueryParam
}

// NewAlibabaDamaiMaitixEticketDistributionQueryRequest 初始化AlibabaDamaiMaitixEticketDistributionQueryAPIRequest对象
func NewAlibabaDamaiMaitixEticketDistributionQueryRequest() *AlibabaDamaiMaitixEticketDistributionQueryAPIRequest {
	return &AlibabaDamaiMaitixEticketDistributionQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDamaiMaitixEticketDistributionQueryAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMaitixEticketDistributionQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.maitix.eticket.distribution.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMaitixEticketDistributionQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDamaiMaitixEticketDistributionQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参param
func (r *AlibabaDamaiMaitixEticketDistributionQueryAPIRequest) SetParam(_param *EticketQueryParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaDamaiMaitixEticketDistributionQueryAPIRequest) GetParam() *EticketQueryParam {
	return r._param
}

var poolAlibabaDamaiMaitixEticketDistributionQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDamaiMaitixEticketDistributionQueryRequest()
	},
}

// GetAlibabaDamaiMaitixEticketDistributionQueryRequest 从 sync.Pool 获取 AlibabaDamaiMaitixEticketDistributionQueryAPIRequest
func GetAlibabaDamaiMaitixEticketDistributionQueryAPIRequest() *AlibabaDamaiMaitixEticketDistributionQueryAPIRequest {
	return poolAlibabaDamaiMaitixEticketDistributionQueryAPIRequest.Get().(*AlibabaDamaiMaitixEticketDistributionQueryAPIRequest)
}

// ReleaseAlibabaDamaiMaitixEticketDistributionQueryAPIRequest 将 AlibabaDamaiMaitixEticketDistributionQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaDamaiMaitixEticketDistributionQueryAPIRequest(v *AlibabaDamaiMaitixEticketDistributionQueryAPIRequest) {
	v.Reset()
	poolAlibabaDamaiMaitixEticketDistributionQueryAPIRequest.Put(v)
}
