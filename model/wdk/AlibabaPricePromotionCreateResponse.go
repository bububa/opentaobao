package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
营销档期活动创建 APIResponse
alibaba.price.promotion.create

大润发-盒马帮提供新增创建营销活动
*/
type AlibabaPricePromotionCreateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_price_promotion_create_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 档期活动ID
    
    Result   int64 `json:"result,omitempty" xml:"