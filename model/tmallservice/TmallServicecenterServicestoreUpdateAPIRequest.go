package tmallservice

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallServicecenterServicestoreUpdateAPIRequest) Reset() {
	r._paramServiceStoreDTO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterServicestoreUpdateAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.servicestore.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterServicestoreUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterServicestoreUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamServiceStoreDTO is ParamServiceStoreDTO Setter
// 网点/门店
func (r *TmallServicecenterServicestoreUpdateAPIRequest) SetParamServiceStoreDTO(_paramServiceStoreDTO *ServiceStoreDto) error {
	r._paramServiceStoreDTO = _paramServiceStoreDTO
	r.Set("param_service_store_d_t_o", _paramServiceStoreDTO)
	return nil
}

// GetParamServiceStoreDTO ParamServiceStoreDTO Getter
func (r TmallServicecenterServicestoreUpdateAPIRequest) GetParamServiceStoreDTO() *ServiceStoreDto {
	return r._paramServiceStoreDTO
}

var poolTmallServicecenterServicestoreUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallServicecenterServicestoreUpdateRequest()
	},
}

// GetTmallServicecenterServicestoreUpdateRequest 从 sync.Pool 获取 TmallServicecenterServicestoreUpdateAPIRequest
func GetTmallServicecenterServicestoreUpdateAPIRequest() *TmallServicecenterServicestoreUpdateAPIRequest {
	return poolTmallServicecenterServicestoreUpdateAPIRequest.Get().(*TmallServicecenterServicestoreUpdateAPIRequest)
}

// ReleaseTmallServicecenterServicestoreUpdateAPIRequest 将 TmallServicecenterServicestoreUpdateAPIRequest 放入 sync.Pool
func ReleaseTmallServicecenterServicestoreUpdateAPIRequest(v *TmallServicecenterServicestoreUpdateAPIRequest) {
	v.Reset()
	poolTmallServicecenterServicestoreUpdateAPIRequest.Put(v)
}
