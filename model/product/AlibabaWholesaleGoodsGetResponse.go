package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询阿里巴巴批发市场商品详情 APIResponse
alibaba.wholesale.goods.get

查询阿里巴巴批发市场商品详情
*/
type AlibabaWholesaleGoodsGetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWholesaleGoodsGetResponse `json:"alibaba_wholesale_goods_get_response,omitempty"`
}

type AlibabaWholesaleGoodsGetResponse struct {

    // wholesale goods detail result
    WholesaleGoodsResult   *WholesaleGoodsOpenResult `json:"wholesale_goods_result,omitempty"`

}