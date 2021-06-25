package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
产品sku编辑接口 APIResponse
taobao.fenxiao.product.sku.update

产品SKU信息更新
*/
type TaobaoFenxiaoProductSkuUpdateAPIResponse struct {
    model.CommonResponse
    Response *TaobaoFenxiaoProductSkuUpdateResponse `json:"taobao_fenxiao_product_sku_update_response,omitempty"`
}

type TaobaoFenxiaoProductSkuUpdateResponse struct {

    // 操作结果
    Result   bool `json:"result,omitempty"`

    // 操作时间
    Created   string `json:"created,omitempty"`

}
