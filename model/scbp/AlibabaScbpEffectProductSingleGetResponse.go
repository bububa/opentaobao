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
    AlibabaScbpEffectProductSingleGetResponse
}

type AlibabaScbpEffectProductSingleGetResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_effect_product_single_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 单个产品的效果数据列表
    
    SProductEffectList   []SingleProductEffectDto `json:"s_product_effect_list,omitempty" xml:"s_product_effect_list>single_product_effect_dto,omitempty"`
    
    
    // 总个数
    
    TotalNum   int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`

    
    // 总页数
    
    TotalPage   int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`

    
}
