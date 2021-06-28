package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
单个产品的报表 APIResponse
alibaba.scbp.effect.product.single.get

单个产品的报表
*/
type AlibabaScbpEffectProductSingleGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_scbp_effect_product_single_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 单个产品的效果数据列表
    
    SProductEffectList   []SingleProductEffectDto `json:"s_product_effect_list,omitempty" xml:"