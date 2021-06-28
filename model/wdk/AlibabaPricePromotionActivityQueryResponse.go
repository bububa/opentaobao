package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询盒马帮档期活动详情 APIResponse
alibaba.price.promotion.activity.query

查询盒马帮档期活动详情
*/
type AlibabaPricePromotionActivityQueryAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaPricePromotionActivityQueryResponse `json:"alibaba_price_promotion_activity_query_response,omitempty"` 
    AlibabaPricePromotionActivityQueryResponse
}

/* model for simplify = false
type AlibabaPricePromotionActivityQueryResponse struct {

    // 接口调用是否成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

    // 大润发促销档期数据
    
    Data   string `json:"data,omitempty"`
    

    // 错误编码
    
    ResultCode   int64 `json:"result_code,omitempty"`
    

    // 大润发档期数据
    
    TotalRecord   int64 `json:"total_record,omitempty"`
    

    // 错误参数
    
    Message   string `json:"message,omitempty"`
    

}
*/

type AlibabaPricePromotionActivityQueryResponse struct {

    // 接口调用是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // 大润发促销档期数据
    Data   string `json:"data,omitempty"`

    // 错误编码
    ResultCode   int64 `json:"result_code,omitempty"`

    // 大润发档期数据
    TotalRecord   int64 `json:"total_record,omitempty"`

    // 错误参数
    Message   string `json:"message,omitempty"`

}
