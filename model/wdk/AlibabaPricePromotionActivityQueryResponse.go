package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询盒马帮档期活动详情 APIResponse
alibaba.price.promotion.activity.query

查询盒马帮档期活动详情
*/
type AlibabaPricePromotionActivityQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_price_promotion_activity_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口调用是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"