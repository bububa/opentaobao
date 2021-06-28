package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
仓内加工单创建接口 APIResponse
taobao.qimen.storeprocess.create

ERP调用奇门的接口,创建仓内加工单
*/
type TaobaoQimenStoreprocessCreateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoQimenStoreprocessCreateResponse `json:"qimen_storeprocess_create_response,omitempty"` 
    TaobaoQimenStoreprocessCreateResponse
}

/* model for simplify = false
type TaobaoQimenStoreprocessCreateResponse struct {

    // 
    
    Response  *struct {
        StoreProcessCreateResponse  *StoreProcessCreateResponse `json:"store_process_create_response,omitempty"`
    } `json:"response,omitempty"`
    

}
*/

type TaobaoQimenStoreprocessCreateResponse struct {

    // 
    Response   *StoreProcessCreateResponse `json:"response,omitempty"`

}
