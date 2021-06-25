package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店服务能力授权接口 APIRequest
tmall.mallitemcenter.supplier.ability.update

门店服务能力授权
*/
type TmallMallitemcenterSupplierAbilityUpdateRequest struct {
    model.Params

    // 入参
    param0   *EnableServiceStoreRequestDto 

}

func NewTmallMallitemcenterSupplierAbilityUpdateRequest() *TmallMallitemcenterSupplierAbilityUpdateRequest{
    return &TmallMallitemcenterSupplierAbilityUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TmallMallitemcenterSupplierAbilityUpdateRequest) GetApiMethodName() string {
    return "tmall.mallitemcenter.supplier.ability.update"
}

func (r TmallMallitemcenterSupplierAbilityUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallMallitemcenterSupplierAbilityUpdateRequest) SetParam0(param0 *EnableServiceStoreRequestDto) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r TmallMallitemcenterSupplierAbilityUpdateRequest) GetParam0() *EnableServiceStoreRequestDto {
    return r.param0
}

