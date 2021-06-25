package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改门店信息 APIRequest
tmall.servicecenter.servicestore.update

用于修改门店/网点信息。多个业务共用
*/
type TmallServicecenterServicestoreUpdateRequest struct {
    model.Params

    // 网点/门店
    paramServiceStoreDTO   *ServiceStoreDto 

}

func NewTmallServicecenterServicestoreUpdateRequest() *TmallServicecenterServicestoreUpdateRequest{
    return &TmallServicecenterServicestoreUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServicecenterServicestoreUpdateRequest) GetApiMethodName() string {
    return "tmall.servicecenter.servicestore.update"
}

func (r TmallServicecenterServicestoreUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServicecenterServicestoreUpdateRequest) SetParamServiceStoreDTO(paramServiceStoreDTO *ServiceStoreDto) error {
    r.paramServiceStoreDTO = paramServiceStoreDTO
    r.Set("param_service_store_d_t_o", paramServiceStoreDTO)
    return nil
}

func (r TmallServicecenterServicestoreUpdateRequest) GetParamServiceStoreDTO() *ServiceStoreDto {
    return r.paramServiceStoreDTO
}

