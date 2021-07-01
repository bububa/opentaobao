package maitix

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMaitixEticketDistributionQueryAPIRequest
分销电子票查询接口 API请求
alibaba.damai.maitix.eticket.distribution.query

分销电子票查询接口 */
type AlibabaDamaiMaitixEticketDistributionQueryAPIRequest struct {
	model.Params
	// 入参param
	_param *EticketQueryParam
}

// NewAlibabaDamaiMaitixEticketDistributionQueryRequest 初始化AlibabaDamaiMaitixEticketDistributionQueryAPIRequest对象
func NewAlibabaDamaiMaitixEticketDistributionQueryRequest() *AlibabaDamaiMaitixEticketDistributionQueryAPIRequest {
	return &AlibabaDamaiMaitixEticketDistributionQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMaitixEticketDistributionQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.maitix.eticket.distribution.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMaitixEticketDistributionQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param Setter
// 入参param
func (r *AlibabaDamaiMaitixEticketDistributionQueryAPIRequest) SetParam(_param *EticketQueryParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// Get Param Getter
func (r AlibabaDamaiMaitixEticketDistributionQueryAPIRequest) GetParam() *EticketQueryParam {
	return r._param
}
