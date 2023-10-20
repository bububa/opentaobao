package store

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPlaceStoreExtendUpdateAPIRequest 商户门店拓展信息更新接口 API请求
// taobao.place.store.extend.update
//
// 更新商户门店拓展信息（tags、attribute、bizAtrribute）更新接口
type TaobaoPlaceStoreExtendUpdateAPIRequest struct {
	model.Params
	// 更新数据
	_paramUpdateStoreExtendDTO *UpdateStoreExtendDto
}

// NewTaobaoPlaceStoreExtendUpdateRequest 初始化TaobaoPlaceStoreExtendUpdateAPIRequest对象
func NewTaobaoPlaceStoreExtendUpdateRequest() *TaobaoPlaceStoreExtendUpdateAPIRequest {
	return &TaobaoPlaceStoreExtendUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPlaceStoreExtendUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.place.store.extend.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPlaceStoreExtendUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPlaceStoreExtendUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamUpdateStoreExtendDTO is ParamUpdateStoreExtendDTO Setter
// 更新数据
func (r *TaobaoPlaceStoreExtendUpdateAPIRequest) SetParamUpdateStoreExtendDTO(_paramUpdateStoreExtendDTO *UpdateStoreExtendDto) error {
	r._paramUpdateStoreExtendDTO = _paramUpdateStoreExtendDTO
	r.Set("param_update_store_extend_d_t_o", _paramUpdateStoreExtendDTO)
	return nil
}

// GetParamUpdateStoreExtendDTO ParamUpdateStoreExtendDTO Getter
func (r TaobaoPlaceStoreExtendUpdateAPIRequest) GetParamUpdateStoreExtendDTO() *UpdateStoreExtendDto {
	return r._paramUpdateStoreExtendDTO
}
