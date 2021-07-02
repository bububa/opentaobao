package maitix

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMaitixOrderQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.maitix.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMaitixOrderQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
