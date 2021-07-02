package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItemBarcodeUpdateAPIResponse 更新商品条形码信息 API返回值
// taobao.item.barcode.update
//
// 通过该接口，将商品以及SKU上得条形码信息补全
type TaobaoItemBarcodeUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoItemBarcodeUpdateAPIResponseModel
}

// TaobaoItemBarcodeUpdateAPIResponseModel is 更新商品条形码信息 成功返回结果
type TaobaoItemBarcodeUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"item_barcode_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品结构里的num_iid，modified
	Item *Item `json:"item,omitempty" xml:"item,omitempty"`
}
