package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
定向推广 按照id操作推广计划的产品，包括新增，删除和更新 API返回值 
alibaba.scbp.target.ad.plan.update.products

定向推广 按照id操作推广计划的产品，包括新增，删除和更新
*/
type AlibabaScbpTargetAdPlanUpdateProductsAPIResponse struct {
    model.CommonResponse
    AlibabaScbpTargetAdPlanUpdateProductsResponse
}

// 定向推广 按照id操作推广计划的产品，包括新增，删除和更新 成功返回结果
type AlibabaScbpTargetAdPlanUpdateProductsResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_target_ad_plan_update_products_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 操作成功的商品ID列表
    ProductIdList   []int64 `json:"product_id_list,omitempty" xml:"product_id_list>int64,omitempty"`
}
