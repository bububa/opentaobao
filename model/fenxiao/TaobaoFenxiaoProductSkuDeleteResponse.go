package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
产品SKU删除接口 APIResponse
taobao.fenxiao.product.sku.delete

根据sku properties删除sku数据
*/
type TaobaoFenxiaoProductSkuDeleteAPIResponse struct {
    model.CommonResponse
    Response *TaobaoFenxiaoProductSkuDeleteResponse `json:"taobao_fenxiao_product_sku_delete_response,omitempty"`
}

type TaobaoFenxiaoProductSkuDeleteResponse struct {

    // 操作结果
    Result   bool `json:"result,omitempty"`

    // 操作时间
    Created   string `json:"created,omitempty"`

}
