package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
SKU查询接口 APIResponse
taobao.fenxiao.product.skus.get

产品sku查询
*/
type TaobaoFenxiaoProductSkusGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoFenxiaoProductSkusGetResponse `json:"fenxiao_product_skus_get_response,omitempty"` 
    TaobaoFenxiaoProductSkusGetResponse
}

/* model for simplify = false
type TaobaoFenxiaoProductSkusGetResponse struct {

    // sku信息
    
    Skus  struct {
        FenxiaoSku  []FenxiaoSku `json:"fenxiao_sku,omitempty"`
    } `json:"skus,omitempty"`
    

    // 记录数量
    
    TotalResults   int64 `json:"total_results,omitempty"`
    

}
*/

type TaobaoFenxiaoProductSkusGetResponse struct {

    // sku信息
    Skus   []FenxiaoSku `json:"skus,omitempty"`

    // 记录数量
    TotalResults   int64 `json:"total_results,omitempty"`

}
