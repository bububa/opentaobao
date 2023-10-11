package axindata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelFscRouteApiDivisionApplyAPIRequest 线路供应商提交新增城市申请 API请求
// taobao.alitrip.travel.fsc.route.api.division.apply
//
// 线路供应商提交新增城市申请
type TaobaoAlitripTravelFscRouteApiDivisionApplyAPIRequest struct {
	model.Params
	// fscDivisionApplyRequest
	_fscDivisionApplyRequest *FscDivisionApplyRequest
}

// NewTaobaoAlitripTravelFscRouteApiDivisionApplyRequest 初始化TaobaoAlitripTravelFscRouteApiDivisionApplyAPIRequest对象
func NewTaobaoAlitripTravelFscRouteApiDivisionApplyRequest() *TaobaoAlitripTravelFscRouteApiDivisionApplyAPIRequest {
	return &TaobaoAlitripTravelFscRouteApiDivisionApplyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelFscRouteApiDivisionApplyAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.fsc.route.api.division.apply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelFscRouteApiDivisionApplyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripTravelFscRouteApiDivisionApplyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFscDivisionApplyRequest is FscDivisionApplyRequest Setter
// fscDivisionApplyRequest
func (r *TaobaoAlitripTravelFscRouteApiDivisionApplyAPIRequest) SetFscDivisionApplyRequest(_fscDivisionApplyRequest *FscDivisionApplyRequest) error {
	r._fscDivisionApplyRequest = _fscDivisionApplyRequest
	r.Set("fsc_division_apply_request", _fscDivisionApplyRequest)
	return nil
}

// GetFscDivisionApplyRequest FscDivisionApplyRequest Getter
func (r TaobaoAlitripTravelFscRouteApiDivisionApplyAPIRequest) GetFscDivisionApplyRequest() *FscDivisionApplyRequest {
	return r._fscDivisionApplyRequest
}
