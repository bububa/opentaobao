package wlbimports

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取运单信息 APIResponse
taobao.wlb.imports.waybill.get

一般进口商家，获取订单的电子面单链接地址
*/
type TaobaoWlbImportsWaybillGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoWlbImportsWaybillGetResponse `json:"wlb_imports_waybill_get_response,omitempty"` 
    TaobaoWlbImportsWaybillGetResponse
}

/* model for simplify = false
type TaobaoWlbImportsWaybillGetResponse struct {

    // 电子面单链接地址
    
    Waybillurl   string `json:"waybillurl,omitempty"`
    

}
*/

type TaobaoWlbImportsWaybillGetResponse struct {

    // 电子面单链接地址
    Waybillurl   string `json:"waybillurl,omitempty"`

}
