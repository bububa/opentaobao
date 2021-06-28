package lifeservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
门店关系新增 APIResponse
taobao.place.store.relation.add

新增授权用户的门店关系信息
*/
type TaobaoPlaceStoreRelationAddAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoPlaceStoreRelationAddResponse `json:"place_store_relation_add_response,omitempty"` 
    TaobaoPlaceStoreRelationAddResponse
}

/* model for simplify = false
type TaobaoPlaceStoreRelationAddResponse struct {

    // 是否失败
    
    Failure   bool `json:"failure,omitempty"`
    

    // 返回结果：true成功；false失败
    
    Result   int64 `json:"result,omitempty"`
    

    // 个数
    
    TotalNum   int64 `json:"total_num,omitempty"`
    

}
*/

type TaobaoPlaceStoreRelationAddResponse struct {

    // 是否失败
    Failure   bool `json:"failure,omitempty"`

    // 返回结果：true成功；false失败
    Result   int64 `json:"result,omitempty"`

    // 个数
    TotalNum   int64 `json:"total_num,omitempty"`

}
