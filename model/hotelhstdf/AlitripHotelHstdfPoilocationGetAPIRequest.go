package hotelhstdf

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitriphotelhstdfpoilocationgetAPIRequest 根据平台城市id分页查询poi location API请求
// alitrip.hotel.hstdf.poilocation.get
//
// 根据平台城市id分页查询poi location
type AlitriphotelhstdfpoilocationgetAPIRequest struct {
	model.Params
	// 参数封装
	_paramGetByTrdiDivisionIdParam *GetByTrdiDivisionIdParam
}

// NewAlitriphotelhstdfpoilocationgetRequest 初始化AlitriphotelhstdfpoilocationgetAPIRequest对象
func NewAlitriphotelhstdfpoilocationgetRequest() *AlitriphotelhstdfpoilocationgetAPIRequest {
	return &AlitriphotelhstdfpoilocationgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitriphotelhstdfpoilocationgetAPIRequest) GetApiMethodName() string {
	return "alitrip.hotel.hstdf.poilocation.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitriphotelhstdfpoilocationgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitriphotelhstdfpoilocationgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamGetByTrdiDivisionIdParam is ParamGetByTrdiDivisionIdParam Setter
// 参数封装
func (r *AlitriphotelhstdfpoilocationgetAPIRequest) SetParamGetByTrdiDivisionIdParam(_paramGetByTrdiDivisionIdParam *GetByTrdiDivisionIdParam) error {
	r._paramGetByTrdiDivisionIdParam = _paramGetByTrdiDivisionIdParam
	r.Set("param_get_by_trdi_division_id_param", _paramGetByTrdiDivisionIdParam)
	return nil
}

// GetParamGetByTrdiDivisionIdParam ParamGetByTrdiDivisionIdParam Getter
func (r AlitriphotelhstdfpoilocationgetAPIRequest) GetParamGetByTrdiDivisionIdParam() *GetByTrdiDivisionIdParam {
	return r._paramGetByTrdiDivisionIdParam
}
