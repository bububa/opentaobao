package store

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商户门店拓展信息更新接口 APIRequest
taobao.place.store.extend.update

更新商户门店拓展信息（tags、attribute、bizAtrribute）更新接口
*/
type TaobaoPlaceStoreExtendUpdateRequest struct {
    model.Params

    // 更新数据
    paramUpdateStoreExtendDTO   *UpdateStoreExtendDto 

}

func NewTaobaoPlaceStoreExtendUpdateRequest() *TaobaoPlaceStoreExtendUpdateRequest{
    return &TaobaoPlaceStoreExtendUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPlaceStoreExtendUpdateRequest) GetApiMethodName() string {
    return "taobao.place.store.extend.update"
}

func (r TaobaoPlaceStoreExtendUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPlaceStoreExtendUpdateRequest) SetParamUpdateStoreExtendDTO(paramUpdateStoreExtendDTO *UpdateStoreExtendDto) error {
    r.paramUpdateStoreExtendDTO = paramUpdateStoreExtendDTO
    r.Set("param_update_store_extend_d_t_o", paramUpdateStoreExtendDTO)
    return nil
}

func (r TaobaoPlaceStoreExtendUpdateRequest) GetParamUpdateStoreExtendDTO() *UpdateStoreExtendDto {
    return r.paramUpdateStoreExtendDTO
}

