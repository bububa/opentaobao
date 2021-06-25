package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
营销档期活动创建 APIResponse
alibaba.price.promotion.create

大润发-盒马帮提供新增创建营销活动
*/
type AlibabaPricePromotionCreateAPIResponse struct {
    model.CommonResponse
    Response *AlibabaPricePromotionCreateResponse `json:"alibaba_price_promotion_create_response,omitempty"`
}

type AlibabaPricePromotionCreateResponse struct {

    // 档期活动ID
    Result   int64 `json:"result,omitempty"`

    // 错误描述
    ErrorDesc   string `json:"error_desc,omitempty"`

    // 错误编码，本期不作识别
    SystemCode   string `json:"system_code,omitempty"`

    // 数量，本期不启用
    TotalNum   int64 `json:"total_num,omitempty"`

    // 创建是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
