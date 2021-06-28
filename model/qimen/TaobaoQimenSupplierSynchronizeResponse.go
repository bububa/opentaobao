package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
供应商同步接口 APIResponse
taobao.qimen.supplier.synchronize

这个接口用来同步供应商信息
*/
type TaobaoQimenSupplierSynchronizeAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoQimenSupplierSynchronizeResponse `json:"qimen_supplier_synchronize_response,omitempty"` 
    TaobaoQimenSupplierSynchronizeResponse
}

/* model for simplify = false
type TaobaoQimenSupplierSynchronizeResponse struct {

    // 
    
    Response  *struct {
        Response  *Response `json:"response,omitempty"`
    } `json:"response,omitempty"`
    

}
*/

type TaobaoQimenSupplierSynchronizeResponse struct {

    // 
    Response   *Response `json:"response,omitempty"`

}
