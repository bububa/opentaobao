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
    Response *TaobaoQimenSupplierSynchronizeResponse `json:"taobao_qimen_supplier_synchronize_response,omitempty"`
}

type TaobaoQimenSupplierSynchronizeResponse struct {

    // 
    Response   *Response `json:"response,omitempty"`

}
