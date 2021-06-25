package mos

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取默认状态下商品列表 APIResponse
alibaba.mos.store.getdefautitems

获取默认状态下商品列表
*/
type AlibabaMosStoreGetdefautitemsAPIResponse struct {
    model.CommonResponse
    Response *AlibabaMosStoreGetdefautitemsResponse `json:"alibaba_mos_store_getdefautitems_response,omitempty"`
}

type AlibabaMosStoreGetdefautitemsResponse struct {

    // result
    Result   *AlibabaMosStoreGetdefautitemsResultDo `json:"result,omitempty"`

}
