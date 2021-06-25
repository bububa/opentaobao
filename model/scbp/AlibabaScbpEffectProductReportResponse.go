package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
所有产品报表 APIResponse
alibaba.scbp.effect.product.report

所有产品报表
*/
type AlibabaScbpEffectProductReportAPIResponse struct {
    model.CommonResponse
    Response *AlibabaScbpEffectProductReportResponse `json:"alibaba_scbp_effect_product_report_response,omitempty"`
}

type AlibabaScbpEffectProductReportResponse struct {

    // 总个数
    TotalNum   int64 `json:"total_num,omitempty"`

    // 总页数
    TotalPage   int64 `json:"total_page,omitempty"`

    // 产品效果数据列表
    ProductEffectList   []ProductEffectDto `json:"product_effect_list,omitempty"`

}
