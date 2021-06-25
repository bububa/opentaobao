package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
批发市场产品搜索 APIResponse
alibaba.wholesale.goods.search

批发市场产品搜索
*/
type AlibabaWholesaleGoodsSearchAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWholesaleGoodsSearchResponse `json:"alibaba_wholesale_goods_search_response,omitempty"`
}

type AlibabaWholesaleGoodsSearchResponse struct {

    // 在线批发商品搜索结果
    WholesaleGoodsSearchResult   *WholesaleSearchOpenResult `json:"wholesale_goods_search_result,omitempty"`

}
