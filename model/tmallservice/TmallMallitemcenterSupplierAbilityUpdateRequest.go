package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店服务能力授权接口 API请求
tmall.mallitemcenter.supplier.ability.update

门店服务能力授权
*/
type TmallMallitemcenterSupplierAbilityUpdateAPIRequest struct {
    model.Params
    // 入参
    _param0   *EnableServiceStoreRequestDTO
}

// 初始化TmallMallitemcenterSupplierAbilityUpdateAPIRequest对象
func NewTmallMallitemcenterSupplierAbilityUpdateRequest() *TmallMallitemcenterSupplierAbilityUpdateAPIRequest{
    return &TmallMallitemcenterSupplierAbilityUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallMallitemcenterSupplierAbilityUpdateAPIRequest) GetApiMethodName() string {
    return "tmall.mallitemcenter.supplier.ability.update"
}

// IRequest interface 方法, 获取API参数
func (r TmallMallitemcenterSupplierAbilityUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 入参
func (r *TmallMallitemcenterSupplierAbilityUpdateAPIRequest) SetParam0(_param0 *EnableServiceStoreRequestDTO) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r TmallMallitemcenterSupplierAbilityUpdateAPIRequest) GetParam0() *EnableServiceStoreRequestDTO {
    return r._param0
}
