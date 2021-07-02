package hotelhstdf

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripHotelHstdfPoilocationGetAPIRequest 根据平台城市id分页查询poi location API请求
// alitrip.hotel.hstdf.poilocation.get
//
// 根据平台城市id分页查询poi location
type AlitripHotelHstdfPoilocationGetAPIRequest struct {
	model.Params
	// 参数封装
	_paramGetByTrdiDivisionIdParam *GetByTrdiDivisionIdParam
}

// NewAlitripHotelHstdfPoilocationGetRequest 初始化AlitripHotelHstdfPoilocationGetAPIRequest对象
func NewAlitripHotelHstdfPoilocationGetRequest() *AlitripHotelHstdfPoilocationGetAPIRequest {
	return &AlitripHotelHstdfPoilocationGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripHotelHstdfPoilocationGetAPIRequest) GetApiMethodName() string {
	return "alitrip.hotel.hstdf.poilocation.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripHotelHstdfPoilocationGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParamGetByTrdiDivisionIdParam is ParamGetByTrdiDivisionIdParam Setter
// 参数封装
func (r *AlitripHotelHstdfPoilocationGetAPIRequest) SetParamGetByTrdiDivisionIdParam(_paramGetByTrdiDivisionIdParam *GetByTrdiDivisionIdParam) error {
	r._paramGetByTrdiDivisionIdParam = _paramGetByTrdiDivisionIdParam
	r.Set("param_get_by_trdi_division_id_param", _paramGetByTrdiDivisionIdParam)
	return nil
}

// GetParamGetByTrdiDivisionIdParam ParamGetByTrdiDivisionIdParam Getter
func (r AlitripHotelHstdfPoilocationGetAPIRequest) GetParamGetByTrdiDivisionIdParam() *GetByTrdiDivisionIdParam {
	return r._paramGetByTrdiDivisionIdParam
}
