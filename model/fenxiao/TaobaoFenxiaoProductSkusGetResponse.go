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
    Response *TaobaoFenxiaoProductSkusGetResponse `json:"taobao_fenxiao_product_skus_get_response,omitempty"`
}

type TaobaoFenxiaoProductSkusGetResponse struct {

    // sku信息
    Skus   []FenxiaoSku `json:"skus,omitempty"`

    // 记录数量
    TotalResults   int64 `json:"total_results,omitempty"`

}
