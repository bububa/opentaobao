package alihealthoutflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthoutflowdrugsupervisionqueryAPIRequest 监管平台药品查询 API请求
// alibaba.alihealth.outflow.drug.supervision.query
//
// 获取监管平台药品数据
type AlibabaalihealthoutflowdrugsupervisionqueryAPIRequest struct {
	model.Params
	// 请求
	_request1 *OuterDrugVo
}

// NewAlibabaalihealthoutflowdrugsupervisionqueryRequest 初始化AlibabaalihealthoutflowdrugsupervisionqueryAPIRequest对象
func NewAlibabaalihealthoutflowdrugsupervisionqueryRequest() *AlibabaalihealthoutflowdrugsupervisionqueryAPIRequest {
	return &AlibabaalihealthoutflowdrugsupervisionqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthoutflowdrugsupervisionqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.outflow.drug.supervision.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthoutflowdrugsupervisionqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthoutflowdrugsupervisionqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest1 is Request1 Setter
// 请求
func (r *AlibabaalihealthoutflowdrugsupervisionqueryAPIRequest) SetRequest1(_request1 *OuterDrugVo) error {
	r._request1 = _request1
	r.Set("request1", _request1)
	return nil
}

// GetRequest1 Request1 Getter
func (r AlibabaalihealthoutflowdrugsupervisionqueryAPIRequest) GetRequest1() *OuterDrugVo {
	return r._request1
}
