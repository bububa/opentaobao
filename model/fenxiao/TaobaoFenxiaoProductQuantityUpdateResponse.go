package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
产品库存修改 APIResponse
taobao.fenxiao.product.quantity.update

修改产品库存信息，支持全量修改以及增量修改两种方式
*/
type TaobaoFenxiaoProductQuantityUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"fenxiao_product_quantity_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 操作结果
    
    Result   bool `json:"result,omitempty" xml:"