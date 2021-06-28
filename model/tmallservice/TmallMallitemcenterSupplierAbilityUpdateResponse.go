package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
门店服务能力授权接口 APIResponse
tmall.mallitemcenter.supplier.ability.update

门店服务能力授权
*/
type TmallMallitemcenterSupplierAbilityUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_mallitemcenter_supplier_ability_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口返回model
    
    Result   *TmallMallitemcenterSupplierAbilityUpdateResult `json:"result,omitempty" xml:"