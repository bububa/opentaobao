package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainaoxiangcooperatedistributorqueryAPIRequest 商家关系查询分销商 API请求
// alibaba.dchain.aoxiang.cooperate.distributor.query
//
// 商家关系查询分销商
type AlibabadchainaoxiangcooperatedistributorqueryAPIRequest struct {
	model.Params
	// 请求入参
	_queryDistributorRequest *QueryDistributorRequest
}

// NewAlibabadchainaoxiangcooperatedistributorqueryRequest 初始化AlibabadchainaoxiangcooperatedistributorqueryAPIRequest对象
func NewAlibabadchainaoxiangcooperatedistributorqueryRequest() *AlibabadchainaoxiangcooperatedistributorqueryAPIRequest {
	return &AlibabadchainaoxiangcooperatedistributorqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadchainaoxiangcooperatedistributorqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.cooperate.distributor.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadchainaoxiangcooperatedistributorqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadchainaoxiangcooperatedistributorqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryDistributorRequest is QueryDistributorRequest Setter
// 请求入参
func (r *AlibabadchainaoxiangcooperatedistributorqueryAPIRequest) SetQueryDistributorRequest(_queryDistributorRequest *QueryDistributorRequest) error {
	r._queryDistributorRequest = _queryDistributorRequest
	r.Set("query_distributor_request", _queryDistributorRequest)
	return nil
}

// GetQueryDistributorRequest QueryDistributorRequest Getter
func (r AlibabadchainaoxiangcooperatedistributorqueryAPIRequest) GetQueryDistributorRequest() *QueryDistributorRequest {
	return r._queryDistributorRequest
}
