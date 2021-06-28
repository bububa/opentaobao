package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
淘鲜达券信息查询接口 APIResponse
alibaba.wdk.bm.coupon.query

淘鲜达品牌营销的券信息查询接口，基于券id查询券相关信息：券id、券名称、分摊信息、面额、创建时间、开始时间、结束时间
*/
type AlibabaWdkBmCouponQueryAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkBmCouponQueryResponse `json:"alibaba_wdk_bm_coupon_query_response,omitempty"` 
    AlibabaWdkBmCouponQueryResponse
}

/* model for simplify = false
type AlibabaWdkBmCouponQueryResponse struct {

    // 结果
    
    Result  *struct {
        BmResult  *BmResult `json:"bm_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkBmCouponQueryResponse struct {

    // 结果
    Result   *BmResult `json:"result,omitempty"`

}
