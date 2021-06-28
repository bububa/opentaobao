package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
品牌营销导购员券推广统计数据回流 APIResponse
alibaba.txcs.brandmarketing.coupon.statistics.get

请求券统计数据回流
*/
type AlibabaTxcsBrandmarketingCouponStatisticsGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaTxcsBrandmarketingCouponStatisticsGetResponse `json:"alibaba_txcs_brandmarketing_coupon_statistics_get_response,omitempty"` 
    AlibabaTxcsBrandmarketingCouponStatisticsGetResponse
}

/* model for simplify = false
type AlibabaTxcsBrandmarketingCouponStatisticsGetResponse struct {

    // 返回结果
    
    ApiPageResult  *struct {
        ApiPageResult  *ApiPageResult `json:"api_page_result,omitempty"`
    } `json:"api_page_result,omitempty"`
    

}
*/

type AlibabaTxcsBrandmarketingCouponStatisticsGetResponse struct {

    // 返回结果
    ApiPageResult   *ApiPageResult `json:"api_page_result,omitempty"`

}
