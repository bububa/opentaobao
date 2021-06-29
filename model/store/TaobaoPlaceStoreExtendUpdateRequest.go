package store

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商户门店拓展信息更新接口 API请求
taobao.place.store.extend.update

更新商户门店拓展信息（tags、attribute、bizAtrribute）更新接口
*/
type TaobaoPlaceStoreExtendUpdateRequest struct {
    model.Params
    // 更新数据
    _paramUpdateStoreExtendDTO   *UpdateStoreExtendDTO
}

// 初始化TaobaoPlaceStoreExtendUpdateRequest对象
func NewTaobaoPlaceStoreExtendUpdateRequest() *TaobaoPlaceStoreExtendUpdateRequest{
    return &TaobaoPlaceStoreExtendUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPlaceStoreExtendUpdateRequest) GetApiMethodName() string {
    return "taobao.place.store.extend.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPlaceStoreExtendUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamUpdateStoreExtendDTO Setter
// 更新数据
func (r *TaobaoPlaceStoreExtendUpdateRequest) SetParamUpdateStoreExtendDTO(_paramUpdateStoreExtendDTO *UpdateStoreExtendDTO) error {
    r._paramUpdateStoreExtendDTO = _paramUpdateStoreExtendDTO
    r.Set("param_update_store_extend_d_t_o", _paramUpdateStoreExtendDTO)
    return nil
}

// ParamUpdateStoreExtendDTO Getter
func (r TaobaoPlaceStoreExtendUpdateRequest) GetParamUpdateStoreExtendDTO() *UpdateStoreExtendDTO {
    return r._paramUpdateStoreExtendDTO
}
