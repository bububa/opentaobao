package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
门店删除接口 APIResponse
taobao.qimen.store.delete

商家在ERP等系统中调用该接口，删除线下门店
*/
type TaobaoQimenStoreDeleteAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoQimenStoreDeleteResponse `json:"qimen_store_delete_response,omitempty"` 
    TaobaoQimenStoreDeleteResponse
}

/* model for simplify = false
type TaobaoQimenStoreDeleteResponse struct {

    // 响应信息
    
    Message   string `json:"message,omitempty"`
    

    // 响应标示
    
    Flag   string `json:"flag,omitempty"`
    

    // 响应code
    
    QimenCode   string `json:"qimen_code,omitempty"`
    

}
*/

type TaobaoQimenStoreDeleteResponse struct {

    // 响应信息
    Message   string `json:"message,omitempty"`

    // 响应标示
    Flag   string `json:"flag,omitempty"`

    // 响应code
    QimenCode   string `json:"qimen_code,omitempty"`

}
