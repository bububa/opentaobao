package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询活动详情列表 APIResponse
taobao.ump.details.get

分页查询优惠详情列表
*/
type TaobaoUmpDetailsGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoUmpDetailsGetResponse `json:"ump_details_get_response,omitempty"` 
    TaobaoUmpDetailsGetResponse
}

/* model for simplify = false
type TaobaoUmpDetailsGetResponse struct {

    // 活动详情的信息
    
    Contents  struct {
        String  []string `json:"string,omitempty"`
    } `json:"contents,omitempty"`
    

    // 记录总数
    
    TotalCount   int64 `json:"total_count,omitempty"`
    

}
*/

type TaobaoUmpDetailsGetResponse struct {

    // 活动详情的信息
    Contents   []string `json:"contents,omitempty"`

    // 记录总数
    TotalCount   int64 `json:"total_count,omitempty"`

}
