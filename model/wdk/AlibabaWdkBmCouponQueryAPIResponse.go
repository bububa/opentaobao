package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
淘鲜达券信息查询接口 API返回值 
alibaba.wdk.bm.coupon.query

淘鲜达品牌营销的券信息查询接口，基于券id查询券相关信息：券id、券名称、分摊信息、面额、创建时间、开始时间、结束时间
*/
type AlibabaWdkBmCouponQueryAPIResponse struct {
    model.CommonResponse
    AlibabaWdkBmCouponQueryAPIResponseModel
}

// 淘鲜达券信息查询接口 成功返回结果
type AlibabaWdkBmCouponQueryAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_wdk_bm_coupon_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果
    Result   *BmResult `json:"result,omitempty" xml:"result,omitempty"`
}
