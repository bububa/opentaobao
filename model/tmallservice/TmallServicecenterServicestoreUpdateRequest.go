package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改门店信息 API请求
tmall.servicecenter.servicestore.update

用于修改门店/网点信息。多个业务共用
*/
type TmallServicecenterServicestoreUpdateRequest struct {
    model.Params
    // 网点/门店
    paramServiceStoreDTO   *ServiceStoreDto
}

// 初始化TmallServicecenterServicestoreUpdateRequest对象
func NewTmallServicecenterServicestoreUpdateRequest() *TmallServicecenterServicestoreUpdateRequest{
    return &TmallServicecenterServicestoreUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterServicestoreUpdateRequest) GetApiMethodName() string {
    return "tmall.servicecenter.servicestore.update"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterServicestoreUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamServiceStoreDTO Setter
// 网点/门店
func (r *TmallServicecenterServicestoreUpdateRequest) SetParamServiceStoreDTO(paramServiceStoreDTO *ServiceStoreDto) error {
    r.paramServiceStoreDTO = paramServiceStoreDTO
    r.Set("param_service_store_d_t_o", paramServiceStoreDTO)
    return nil
}

// ParamServiceStoreDTO Getter
func (r TmallServicecenterServicestoreUpdateRequest) GetParamServiceStoreDTO() *ServiceStoreDto {
    return r.paramServiceStoreDTO
}
