package icbu

import (
    "github.com/bububa/opentaobao/model"
)

/* 
ICBU商品发布草稿接口 APIResponse
alibaba.icbu.product.add.draft

发布商品草稿,支持sourcing/一口价商品，支持英文和多种语言原发商品
*/
type AlibabaIcbuProductAddDraftAPIResponse struct {
    model.CommonResponse
    Response *AlibabaIcbuProductAddDraftResponse `json:"alibaba_icbu_product_add_draft_response,omitempty"`
}

type AlibabaIcbuProductAddDraftResponse struct {

    // 混淆后的产品ID
    ProductId   string `json:"product_id,omitempty"`

}
