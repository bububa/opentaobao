package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
门店类目获取接口 APIResponse
taobao.qimen.storecategory.get

商家在ERP中调用该接口，获取门店类目
*/
type TaobaoQimenStorecategoryGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoQimenStorecategoryGetResponse `json:"qimen_storecategory_get_response,omitempty"` 
    TaobaoQimenStorecategoryGetResponse
}

/* model for simplify = false
type TaobaoQimenStorecategoryGetResponse struct {

    // 响应信息
    
    Message   string `json:"message,omitempty"`
    

    // 响应标示
    
    Flag   string `json:"flag,omitempty"`
    

    // 响应code
    
    QimenCode   string `json:"qimen_code,omitempty"`
    

    // 类目json字符串
    
    StoreCategory   string `json:"store_category,omitempty"`
    

    // 响应信息
    
    Message   string `json:"message,omitempty"`
    

    // 响应标示
    
    Flag   string `json:"flag,omitempty"`
    

    // 响应code
    
    QimenCode   string `json:"qimen_code,omitempty"`
    

    // 类目json字符串
    
    StoreCategory   string `json:"store_category,omitempty"`
    

}
*/

type TaobaoQimenStorecategoryGetResponse struct {

    // 响应信息
    Message   string `json:"message,omitempty"`

    // 响应标示
    Flag   string `json:"flag,omitempty"`

    // 响应code
    QimenCode   string `json:"qimen_code,omitempty"`

    // 类目json字符串
    StoreCategory   string `json:"store_category,omitempty"`

    // 响应信息
    Message   string `json:"message,omitempty"`

    // 响应标示
    Flag   string `json:"flag,omitempty"`

    // 响应code
    QimenCode   string `json:"qimen_code,omitempty"`

    // 类目json字符串
    StoreCategory   string `json:"store_category,omitempty"`

}
