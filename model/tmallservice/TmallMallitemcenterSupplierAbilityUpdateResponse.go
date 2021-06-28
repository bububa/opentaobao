package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
门店服务能力授权接口 APIResponse
tmall.mallitemcenter.supplier.ability.update

门店服务能力授权
*/
type TmallMallitemcenterSupplierAbilityUpdateAPIResponse struct {
    model.CommonResponse
    // Response *TmallMallitemcenterSupplierAbilityUpdateResponse `json:"tmall_mallitemcenter_supplier_ability_update_response,omitempty"` 
    TmallMallitemcenterSupplierAbilityUpdateResponse
}

/* model for simplify = false
type TmallMallitemcenterSupplierAbilityUpdateResponse struct {

    // 接口返回model
    
    Result  *struct {
        TmallMallitemcenterSupplierAbilityUpdateResult  *TmallMallitemcenterSupplierAbilityUpdateResult `json:"tmall_mallitemcenter_supplier_ability_update_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TmallMallitemcenterSupplierAbilityUpdateResponse struct {

    // 接口返回model
    Result   *TmallMallitemcenterSupplierAbilityUpdateResult `json:"result,omitempty"`

}
