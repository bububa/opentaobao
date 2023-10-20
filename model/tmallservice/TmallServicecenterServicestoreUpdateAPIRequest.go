package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenterservicestoreupdateAPIRequest 修改门店信息 API请求
// tmall.servicecenter.servicestore.update
//
// 用于修改门店/网点信息。多个业务共用
type TmallservicecenterservicestoreupdateAPIRequest struct {
	model.Params
	// 网点/门店
	_paramServiceStoreDTO *ServiceStoreDto
}

// NewTmallservicecenterservicestoreupdateRequest 初始化TmallservicecenterservicestoreupdateAPIRequest对象
func NewTmallservicecenterservicestoreupdateRequest() *TmallservicecenterservicestoreupdateAPIRequest {
	return &TmallservicecenterservicestoreupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecenterservicestoreupdateAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.servicestore.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecenterservicestoreupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecenterservicestoreupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamServiceStoreDTO is ParamServiceStoreDTO Setter
// 网点/门店
func (r *TmallservicecenterservicestoreupdateAPIRequest) SetParamServiceStoreDTO(_paramServiceStoreDTO *ServiceStoreDto) error {
	r._paramServiceStoreDTO = _paramServiceStoreDTO
	r.Set("param_service_store_d_t_o", _paramServiceStoreDTO)
	return nil
}

// GetParamServiceStoreDTO ParamServiceStoreDTO Getter
func (r TmallservicecenterservicestoreupdateAPIRequest) GetParamServiceStoreDTO() *ServiceStoreDto {
	return r._paramServiceStoreDTO
}
