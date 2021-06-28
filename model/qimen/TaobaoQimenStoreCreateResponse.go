package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
门店新增接口 APIResponse
taobao.qimen.store.create

isv调用接口来讲线下门店同步到线上
*/
type TaobaoQimenStoreCreateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoQimenStoreCreateResponse `json:"qimen_store_create_response,omitempty"` 
    TaobaoQimenStoreCreateResponse
}

/* model for simplify = false
type TaobaoQimenStoreCreateResponse struct {

    // 响应信息
    
    Message   string `json:"message,omitempty"`
    

    // 响应标示
    
    Flag   string `json:"flag,omitempty"`
    

    // 响应code
    
    QimenCode   string `json:"qimen_code,omitempty"`
    

    // 返回的门店id
    
    StoreId   int64 `json:"store_id,omitempty"`
    

}
*/

type TaobaoQimenStoreCreateResponse struct {

    // 响应信息
    Message   string `json:"message,omitempty"`

    // 响应标示
    Flag   string `json:"flag,omitempty"`

    // 响应code
    QimenCode   string `json:"qimen_code,omitempty"`

    // 返回的门店id
    StoreId   int64 `json:"store_id,omitempty"`

}
