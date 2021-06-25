package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
产品库存修改 APIResponse
taobao.fenxiao.product.quantity.update

修改产品库存信息，支持全量修改以及增量修改两种方式
*/
type TaobaoFenxiaoProductQuantityUpdateAPIResponse struct {
    model.CommonResponse
    Response *TaobaoFenxiaoProductQuantityUpdateResponse `json:"taobao_fenxiao_product_quantity_update_response,omitempty"`
}

type TaobaoFenxiaoProductQuantityUpdateResponse struct {

    // 操作结果
    Result   bool `json:"result,omitempty"`

    // 操作时间
    Created   string `json:"created,omitempty"`

}
