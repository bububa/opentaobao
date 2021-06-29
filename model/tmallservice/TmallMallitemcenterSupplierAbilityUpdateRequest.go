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
type TmallMallitemcenterSupplierAbilityUpdateRequest struct {
    model.Params
    // 入参
    _param0   *EnableServiceStoreRequestDto
}

// 初始化TmallMallitemcenterSupplierAbilityUpdateRequest对象
func NewTmallMallitemcenterSupplierAbilityUpdateRequest() *TmallMallitemcenterSupplierAbilityUpdateRequest{
    return &TmallMallitemcenterSupplierAbilityUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallMallitemcenterSupplierAbilityUpdateRequest) GetApiMethodName() string {
    return "tmall.mallitemcenter.supplier.ability.update"
}

// IRequest interface 方法, 获取API参数
func (r TmallMallitemcenterSupplierAbilityUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 入参
func (r *TmallMallitemcenterSupplierAbilityUpdateRequest) SetParam0(_param0 *EnableServiceStoreRequestDto) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r TmallMallitemcenterSupplierAbilityUpdateRequest) GetParam0() *EnableServiceStoreRequestDto {
    return r._param0
}
