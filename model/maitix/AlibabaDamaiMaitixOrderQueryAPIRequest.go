package maitix

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadamaimaitixorderqueryAPIRequest 大麦-查询分销单 API请求
// alibaba.damai.maitix.order.query
//
// 查询分销单
type AlibabadamaimaitixorderqueryAPIRequest struct {
	model.Params
	// 分销单查询入参
	_param *MoaOrderQueryParam
}

// NewAlibabadamaimaitixorderqueryRequest 初始化AlibabadamaimaitixorderqueryAPIRequest对象
func NewAlibabadamaimaitixorderqueryRequest() *AlibabadamaimaitixorderqueryAPIRequest {
	return &AlibabadamaimaitixorderqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadamaimaitixorderqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.maitix.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadamaimaitixorderqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadamaimaitixorderqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 分销单查询入参
func (r *AlibabadamaimaitixorderqueryAPIRequest) SetParam(_param *MoaOrderQueryParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabadamaimaitixorderqueryAPIRequest) GetParam() *MoaOrderQueryParam {
	return r._param
}
