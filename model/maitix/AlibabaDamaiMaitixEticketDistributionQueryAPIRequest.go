package maitix

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadamaimaitixeticketdistributionqueryAPIRequest 分销电子票查询接口 API请求
// alibaba.damai.maitix.eticket.distribution.query
//
// 分销电子票查询接口
type AlibabadamaimaitixeticketdistributionqueryAPIRequest struct {
	model.Params
	// 入参param
	_param *EticketQueryParam
}

// NewAlibabadamaimaitixeticketdistributionqueryRequest 初始化AlibabadamaimaitixeticketdistributionqueryAPIRequest对象
func NewAlibabadamaimaitixeticketdistributionqueryRequest() *AlibabadamaimaitixeticketdistributionqueryAPIRequest {
	return &AlibabadamaimaitixeticketdistributionqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadamaimaitixeticketdistributionqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.maitix.eticket.distribution.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadamaimaitixeticketdistributionqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadamaimaitixeticketdistributionqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参param
func (r *AlibabadamaimaitixeticketdistributionqueryAPIRequest) SetParam(_param *EticketQueryParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabadamaimaitixeticketdistributionqueryAPIRequest) GetParam() *EticketQueryParam {
	return r._param
}
