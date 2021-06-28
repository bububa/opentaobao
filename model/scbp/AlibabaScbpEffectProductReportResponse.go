package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
所有产品报表 APIResponse
alibaba.scbp.effect.product.report

所有产品报表
*/
type AlibabaScbpEffectProductReportAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_scbp_effect_product_report_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 总个数
    
    TotalNum   int64 `json:"total_num,omitempty" xml:"