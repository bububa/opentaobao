package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
更新商品条形码信息 APIResponse
taobao.item.barcode.update

通过该接口，将商品以及SKU上得条形码信息补全
*/
type TaobaoItemBarcodeUpdateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoItemBarcodeUpdateResponse `json:"item_barcode_update_response,omitempty"` 
    TaobaoItemBarcodeUpdateResponse
}

/* model for simplify = false
type TaobaoItemBarcodeUpdateResponse struct {

    // 商品结构里的num_iid，modified
    
    Item  *struct {
        Item  *Item `json:"item,omitempty"`
    } `json:"item,omitempty"`
    

}
*/

type TaobaoItemBarcodeUpdateResponse struct {

    // 商品结构里的num_iid，modified
    Item   *Item `json:"item,omitempty"`

}
