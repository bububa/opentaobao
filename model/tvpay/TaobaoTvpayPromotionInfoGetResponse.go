package tvpay

import (
    "github.com/bububa/opentaobao/model"
)

/* 
tv支付查询消费抽奖配置 APIResponse
taobao.tvpay.promotion.info.get

查询消费抽奖配置
*/
type TaobaoTvpayPromotionInfoGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTvpayPromotionInfoGetResponse `json:"tvpay_promotion_info_get_response,omitempty"` 
    TaobaoTvpayPromotionInfoGetResponse
}

/* model for simplify = false
type TaobaoTvpayPromotionInfoGetResponse struct {

    // Top返回对象
    
    Result  *struct {
        TopResultDo  *TopResultDo `json:"top_result_do,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoTvpayPromotionInfoGetResponse struct {

    // Top返回对象
    Result   *TopResultDo `json:"result,omitempty"`

}
