package openmall

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenmallItemSaleareaGetAPIResponse 查询商品可售区域 API返回值
// taobao.openmall.item.salearea.get
//
// 获取openmall商品的可售区域
type TaobaoOpenmallItemSaleareaGetAPIResponse struct {
	model.CommonResponse
	TaobaoOpenmallItemSaleareaGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpenmallItemSaleareaGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpenmallItemSaleareaGetAPIResponseModel).Reset()
}

// TaobaoOpenmallItemSaleareaGetAPIResponseModel is 查询商品可售区域 成功返回结果
type TaobaoOpenmallItemSaleareaGetAPIResponseModel struct {
	XMLName xml.Name `xml:"openmall_item_salearea_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TaobaoOpenmallItemSaleareaGetResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpenmallItemSaleareaGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoOpenmallItemSaleareaGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpenmallItemSaleareaGetAPIResponse)
	},
}

// GetTaobaoOpenmallItemSaleareaGetAPIResponse 从 sync.Pool 获取 TaobaoOpenmallItemSaleareaGetAPIResponse
func GetTaobaoOpenmallItemSaleareaGetAPIResponse() *TaobaoOpenmallItemSaleareaGetAPIResponse {
	return poolTaobaoOpenmallItemSaleareaGetAPIResponse.Get().(*TaobaoOpenmallItemSaleareaGetAPIResponse)
}

// ReleaseTaobaoOpenmallItemSaleareaGetAPIResponse 将 TaobaoOpenmallItemSaleareaGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpenmallItemSaleareaGetAPIResponse(v *TaobaoOpenmallItemSaleareaGetAPIResponse) {
	v.Reset()
	poolTaobaoOpenmallItemSaleareaGetAPIResponse.Put(v)
}
