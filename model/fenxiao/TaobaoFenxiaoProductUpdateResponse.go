package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
更新产品 APIResponse
taobao.fenxiao.product.update

更新分销平台产品数据，不传更新数据返回失败<br><br/>1. 对sku进行增、删操作时，原有的sku_ids字段会被忽略，请使用sku_properties和sku_properties_del。<br>
*/
type TaobaoFenxiaoProductUpdateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoFenxiaoProductUpdateResponse `json:"fenxiao_product_update_response,omitempty"` 
    TaobaoFenxiaoProductUpdateResponse
}

/* model for simplify = false
type TaobaoFenxiaoProductUpdateResponse struct {

    // 产品ID
    
    Pid   int64 `json:"pid,omitempty"`
    

    // 更新时间，时间格式：yyyy-MM-dd HH:mm:ss
    
    Modified   string `json:"modified,omitempty"`
    

}
*/

type TaobaoFenxiaoProductUpdateResponse struct {

    // 产品ID
    Pid   int64 `json:"pid,omitempty"`

    // 更新时间，时间格式：yyyy-MM-dd HH:mm:ss
    Modified   string `json:"modified,omitempty"`

}
