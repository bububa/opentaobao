package axindata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelfscrouteapidivisionapplyAPIRequest 线路供应商提交新增城市申请 API请求
// taobao.alitrip.travel.fsc.route.api.division.apply
//
// 线路供应商提交新增城市申请
type TaobaoalitriptravelfscrouteapidivisionapplyAPIRequest struct {
	model.Params
	// fscDivisionApplyRequest
	_fscDivisionApplyRequest *FscDivisionApplyRequest
}

// NewTaobaoalitriptravelfscrouteapidivisionapplyRequest 初始化TaobaoalitriptravelfscrouteapidivisionapplyAPIRequest对象
func NewTaobaoalitriptravelfscrouteapidivisionapplyRequest() *TaobaoalitriptravelfscrouteapidivisionapplyAPIRequest {
	return &TaobaoalitriptravelfscrouteapidivisionapplyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitriptravelfscrouteapidivisionapplyAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.fsc.route.api.division.apply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitriptravelfscrouteapidivisionapplyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitriptravelfscrouteapidivisionapplyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFscDivisionApplyRequest is FscDivisionApplyRequest Setter
// fscDivisionApplyRequest
func (r *TaobaoalitriptravelfscrouteapidivisionapplyAPIRequest) SetFscDivisionApplyRequest(_fscDivisionApplyRequest *FscDivisionApplyRequest) error {
	r._fscDivisionApplyRequest = _fscDivisionApplyRequest
	r.Set("fsc_division_apply_request", _fscDivisionApplyRequest)
	return nil
}

// GetFscDivisionApplyRequest FscDivisionApplyRequest Getter
func (r TaobaoalitriptravelfscrouteapidivisionapplyAPIRequest) GetFscDivisionApplyRequest() *FscDivisionApplyRequest {
	return r._fscDivisionApplyRequest
}
