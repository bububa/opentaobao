package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTmallgenieScpPlanCurrentRawpoGetAPIRequest
二级物料-PO数据同步 API请求
alibaba.tmallgenie.scp.plan.current.rawpo.get

二级物料-PO数据同步（WO-W[TL]) */
type AlibabaTmallgenieScpPlanCurrentRawpoGetAPIRequest struct {
	model.Params
	// 系统自动生成
	_paramRequest *AbstractRequest
}

// NewAlibabaTmallgenieScpPlanCurrentRawpoGetRequest 初始化AlibabaTmallgenieScpPlanCurrentRawpoGetAPIRequest对象
func NewAlibabaTmallgenieScpPlanCurrentRawpoGetRequest() *AlibabaTmallgenieScpPlanCurrentRawpoGetAPIRequest {
	return &AlibabaTmallgenieScpPlanCurrentRawpoGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanCurrentRawpoGetAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.current.rawpo.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanCurrentRawpoGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamRequest Setter
// 系统自动生成
func (r *AlibabaTmallgenieScpPlanCurrentRawpoGetAPIRequest) SetParamRequest(_paramRequest *AbstractRequest) error {
	r._paramRequest = _paramRequest
	r.Set("param_request", _paramRequest)
	return nil
}

// Get ParamRequest Getter
func (r AlibabaTmallgenieScpPlanCurrentRawpoGetAPIRequest) GetParamRequest() *AbstractRequest {
	return r._paramRequest
}
