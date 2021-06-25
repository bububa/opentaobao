package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
单个产品的报表 APIResponse
alibaba.scbp.effect.product.single.get

单个产品的报表
*/
type AlibabaScbpEffectProductSingleGetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaScbpEffectProductSingleGetResponse `json:"alibaba_scbp_effect_product_single_get_response,omitempty"`
}

type AlibabaScbpEffectProductSingleGetResponse struct {

    // 单个产品的效果数据列表
    SProductEffectList   []SingleProductEffectDto `json:"s_product_effect_list,omitempty"`

    // 总个数
    TotalNum   int64 `json:"total_num,omitempty"`

    // 总页数
    TotalPage   int64 `json:"total_page,omitempty"`

}
