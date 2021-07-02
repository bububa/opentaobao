package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterServicestoreUpdateAPIRequest 修改门店信息 API请求
// tmall.servicecenter.servicestore.update
//
// 用于修改门店/网点信息。多个业务共用
type TmallServicecenterServicestoreUpdateAPIRequest struct {
	model.Params
	// 网点/门店
	_paramServiceStoreDTO *ServiceStoreDto
}

// NewTmallServicecenterServicestoreUpdateRequest 初始化TmallServicecenterServicestoreUpdateAPIRequest对象
func NewTmallServicecenterServicestoreUpdateRequest() *TmallServicecenterServicestoreUpdateAPIRequest {
	return &TmallServicecenterServicestoreUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterServicestoreUpdateAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.servicestore.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterServicestoreUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamServiceStoreDTO Setter
// 网点/门店
func (r *TmallServicecenterServicestoreUpdateAPIRequest) SetParamServiceStoreDTO(_paramServiceStoreDTO *ServiceStoreDto) error {
	r._paramServiceStoreDTO = _paramServiceStoreDTO
	r.Set("param_service_store_d_t_o", _paramServiceStoreDTO)
	return nil
}

// Get ParamServiceStoreDTO Getter
func (r TmallServicecenterServicestoreUpdateAPIRequest) GetParamServiceStoreDTO() *ServiceStoreDto {
	return r._paramServiceStoreDTO
}
