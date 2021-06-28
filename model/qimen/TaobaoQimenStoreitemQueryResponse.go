package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
门店关联商品查询接口 APIResponse
taobao.qimen.storeitem.query

商家在ERP等系统中调用该接口，查询某门店所关联的线上商品列表
*/
type TaobaoQimenStoreitemQueryAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoQimenStoreitemQueryResponse `json:"qimen_storeitem_query_response,omitempty"` 
    TaobaoQimenStoreitemQueryResponse
}

/* model for simplify = false
type TaobaoQimenStoreitemQueryResponse struct {

    // 响应消息
    
    Message   string `json:"message,omitempty"`
    

    // 商品列表
    
    ItemIds  struct {
        Json  []string `json:"string,omitempty"`
    } `json:"item_ids,omitempty"`
    

    // 响应标签
    
    Flag   string `json:"flag,omitempty"`
    

    // 商品总数
    
    TotalLines   int64 `json:"total_lines,omitempty"`
    

    // 响应code
    
    QimenCode   string `json:"qimen_code,omitempty"`
    

}
*/

type TaobaoQimenStoreitemQueryResponse struct {

    // 响应消息
    Message   string `json:"message,omitempty"`

    // 商品列表
    ItemIds   []string `json:"item_ids,omitempty"`

    // 响应标签
    Flag   string `json:"flag,omitempty"`

    // 商品总数
    TotalLines   int64 `json:"total_lines,omitempty"`

    // 响应code
    QimenCode   string `json:"qimen_code,omitempty"`

}
