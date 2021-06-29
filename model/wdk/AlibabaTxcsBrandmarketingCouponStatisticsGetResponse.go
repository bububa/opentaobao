package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
品牌营销导购员券推广统计数据回流 APIResponse
alibaba.txcs.brandmarketing.coupon.statistics.get

请求券统计数据回流
*/
type AlibabaTxcsBrandmarketingCouponStatisticsGetAPIResponse struct {
    model.CommonResponse
    AlibabaTxcsBrandmarketingCouponStatisticsGetResponse
}

type AlibabaTxcsBrandmarketingCouponStatisticsGetResponse struct {
    XMLName xml.Name `xml:"alibaba_txcs_brandmarketing_coupon_statistics_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    ApiPageResult   *ApiPageResult `json:"api_page_result,omitempty" xml:"api_page_result,omitempty"`

    
}
