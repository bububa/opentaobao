package axindata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelFscRouteApiProjectCloseAPIRequest 关闭团期 API请求
// taobao.alitrip.travel.fsc.route.api.project.close
//
// 关闭团期
type TaobaoAlitripTravelFscRouteApiProjectCloseAPIRequest struct {
	model.Params
	// fscProjectCloseRequest
	_fscProjectCloseRequest *FscProjectCloseRequest
}

// NewTaobaoAlitripTravelFscRouteApiProjectCloseRequest 初始化TaobaoAlitripTravelFscRouteApiProjectCloseAPIRequest对象
func NewTaobaoAlitripTravelFscRouteApiProjectCloseRequest() *TaobaoAlitripTravelFscRouteApiProjectCloseAPIRequest {
	return &TaobaoAlitripTravelFscRouteApiProjectCloseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelFscRouteApiProjectCloseAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.fsc.route.api.project.close"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelFscRouteApiProjectCloseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripTravelFscRouteApiProjectCloseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFscProjectCloseRequest is FscProjectCloseRequest Setter
// fscProjectCloseRequest
func (r *TaobaoAlitripTravelFscRouteApiProjectCloseAPIRequest) SetFscProjectCloseRequest(_fscProjectCloseRequest *FscProjectCloseRequest) error {
	r._fscProjectCloseRequest = _fscProjectCloseRequest
	r.Set("fsc_project_close_request", _fscProjectCloseRequest)
	return nil
}

// GetFscProjectCloseRequest FscProjectCloseRequest Getter
func (r TaobaoAlitripTravelFscRouteApiProjectCloseAPIRequest) GetFscProjectCloseRequest() *FscProjectCloseRequest {
	return r._fscProjectCloseRequest
}
