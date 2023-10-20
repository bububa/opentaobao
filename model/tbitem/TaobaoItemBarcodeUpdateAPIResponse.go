package tbitem

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *TaobaoItemBarcodeUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoItemBarcodeUpdateAPIResponseModel).Reset()
}

// TaobaoItemBarcodeUpdateAPIResponseModel is 更新商品条形码信息 成功返回结果
type TaobaoItemBarcodeUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"item_barcode_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品结构里的num_iid，modified
	Item *Item `json:"item,omitempty" xml:"item,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoItemBarcodeUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Item = nil
}

var poolTaobaoItemBarcodeUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoItemBarcodeUpdateAPIResponse)
	},
}

// GetTaobaoItemBarcodeUpdateAPIResponse 从 sync.Pool 获取 TaobaoItemBarcodeUpdateAPIResponse
func GetTaobaoItemBarcodeUpdateAPIResponse() *TaobaoItemBarcodeUpdateAPIResponse {
	return poolTaobaoItemBarcodeUpdateAPIResponse.Get().(*TaobaoItemBarcodeUpdateAPIResponse)
}

// ReleaseTaobaoItemBarcodeUpdateAPIResponse 将 TaobaoItemBarcodeUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoItemBarcodeUpdateAPIResponse(v *TaobaoItemBarcodeUpdateAPIResponse) {
	v.Reset()
	poolTaobaoItemBarcodeUpdateAPIResponse.Put(v)
}
