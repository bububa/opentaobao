package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更新商品条形码信息 APIResponse
taobao.item.barcode.update

通过该接口，将商品以及SKU上得条形码信息补全
*/
type TaobaoItemBarcodeUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"item_barcode_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 商品结构里的num_iid，modified
    
    Item   *Item `json:"item,omitempty" xml:"