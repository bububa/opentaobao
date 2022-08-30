package ascp

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangCooperateDistributorQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.cooperate.distributor.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangCooperateDistributorQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
