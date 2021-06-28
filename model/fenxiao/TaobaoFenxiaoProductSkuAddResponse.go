package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
产品sku添加接口 APIResponse
taobao.fenxiao.product.sku.add

添加产品SKU信息
*/
type TaobaoFenxiaoProductSkuAddAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoFenxiaoProductSkuAddResponse `json:"fenxiao_product_sku_add_response,omitempty"` 
    TaobaoFenxiaoProductSkuAddResponse
}

/* model for simplify = false
type TaobaoFenxiaoProductSkuAddResponse struct {

    // 操作结果
    
    Result   bool `json:"result,omitempty"`
    

    // 操作时间
    
    Created   string `json:"created,omitempty"`
    

}
*/

type TaobaoFenxiaoProductSkuAddResponse struct {

    // 操作结果
    Result   bool `json:"result,omitempty"`

    // 操作时间
    Created   string `json:"created,omitempty"`

}
