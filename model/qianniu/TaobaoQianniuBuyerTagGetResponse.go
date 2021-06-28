package qianniu

import (
    "github.com/bububa/opentaobao/model"
)

/* 
判断买家是否有某些标 APIResponse
taobao.qianniu.buyer.tag.get

判断某个买家是否有某些标
*/
type TaobaoQianniuBuyerTagGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoQianniuBuyerTagGetResponse `json:"qianniu_buyer_tag_get_response,omitempty"` 
    TaobaoQianniuBuyerTagGetResponse
}

/* model for simplify = false
type TaobaoQianniuBuyerTagGetResponse struct {

    // 用户tag信息
    
    UserTagInfo  *struct {
        UserTagQueryResult  *UserTagQueryResult `json:"user_tag_query_result,omitempty"`
    } `json:"user_tag_info,omitempty"`
    

}
*/

type TaobaoQianniuBuyerTagGetResponse struct {

    // 用户tag信息
    UserTagInfo   *UserTagQueryResult `json:"user_tag_info,omitempty"`

}
