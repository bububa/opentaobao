package ascp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangCooperateDistributorQueryAPIRequest 商家关系查询分销商 API请求
// alibaba.dchain.aoxiang.cooperate.distributor.query
//
// 商家关系查询分销商
type AlibabaDchainAoxiangCooperateDistributorQueryAPIRequest struct {
	model.Params
	// 请求入参
	_queryDistributorRequest *QueryDistributorRequest
}

// NewAlibabaDchainAoxiangCooperateDistributorQueryRequest 初始化AlibabaDchainAoxiangCooperateDistributorQueryAPIRequest对象
func NewAlibabaDchainAoxiangCooperateDistributorQueryRequest() *AlibabaDchainAoxiangCooperateDistributorQueryAPIRequest {
	return &AlibabaDchainAoxiangCooperateDistributorQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDchainAoxiangCooperateDistributorQueryAPIRequest) Reset() {
	r._queryDistributorRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangCooperateDistributorQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.cooperate.distributor.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangCooperateDistributorQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDchainAoxiangCooperateDistributorQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryDistributorRequest is QueryDistributorRequest Setter
// 请求入参
func (r *AlibabaDchainAoxiangCooperateDistributorQueryAPIRequest) SetQueryDistributorRequest(_queryDistributorRequest *QueryDistributorRequest) error {
	r._queryDistributorRequest = _queryDistributorRequest
	r.Set("query_distributor_request", _queryDistributorRequest)
	return nil
}

// GetQueryDistributorRequest QueryDistributorRequest Getter
func (r AlibabaDchainAoxiangCooperateDistributorQueryAPIRequest) GetQueryDistributorRequest() *QueryDistributorRequest {
	return r._queryDistributorRequest
}

var poolAlibabaDchainAoxiangCooperateDistributorQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDchainAoxiangCooperateDistributorQueryRequest()
	},
}

// GetAlibabaDchainAoxiangCooperateDistributorQueryRequest 从 sync.Pool 获取 AlibabaDchainAoxiangCooperateDistributorQueryAPIRequest
func GetAlibabaDchainAoxiangCooperateDistributorQueryAPIRequest() *AlibabaDchainAoxiangCooperateDistributorQueryAPIRequest {
	return poolAlibabaDchainAoxiangCooperateDistributorQueryAPIRequest.Get().(*AlibabaDchainAoxiangCooperateDistributorQueryAPIRequest)
}

// ReleaseAlibabaDchainAoxiangCooperateDistributorQueryAPIRequest 将 AlibabaDchainAoxiangCooperateDistributorQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaDchainAoxiangCooperateDistributorQueryAPIRequest(v *AlibabaDchainAoxiangCooperateDistributorQueryAPIRequest) {
	v.Reset()
	poolAlibabaDchainAoxiangCooperateDistributorQueryAPIRequest.Put(v)
}
