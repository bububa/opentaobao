package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
所有产品报表 API返回值 
alibaba.scbp.effect.product.report

所有产品报表
*/
type AlibabaScbpEffectProductReportAPIResponse struct {
    model.CommonResponse
    AlibabaScbpEffectProductReportResponse
}

// 所有产品报表 成功返回结果
type AlibabaScbpEffectProductReportResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_effect_product_report_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 总个数
    TotalNum   int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
    // 总页数
    TotalPage   int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
    // 产品效果数据列表
    ProductEffectList   []ProductEffectDto `json:"product_effect_list,omitempty" xml:"product_effect_list>product_effect_dto,omitempty"`
}
