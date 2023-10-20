package tbitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItemSkuAddAPIResponse 添加SKU API返回值
// taobao.item.sku.add
//
// 新增一个sku到num_iid指定的商品中 &lt;br/&gt;传入的iid所对应的商品必须属于当前会话的用户
type TaobaoItemSkuAddAPIResponse struct {
	model.CommonResponse
	TaobaoItemSkuAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoItemSkuAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoItemSkuAddAPIResponseModel).Reset()
}

// TaobaoItemSkuAddAPIResponseModel is 添加SKU 成功返回结果
type TaobaoItemSkuAddAPIResponseModel struct {
	XMLName xml.Name `xml:"item_sku_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// sku
	Sku *Sku `json:"sku,omitempty" xml:"sku,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoItemSkuAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Sku = nil
}

var poolTaobaoItemSkuAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoItemSkuAddAPIResponse)
	},
}

// GetTaobaoItemSkuAddAPIResponse 从 sync.Pool 获取 TaobaoItemSkuAddAPIResponse
func GetTaobaoItemSkuAddAPIResponse() *TaobaoItemSkuAddAPIResponse {
	return poolTaobaoItemSkuAddAPIResponse.Get().(*TaobaoItemSkuAddAPIResponse)
}

// ReleaseTaobaoItemSkuAddAPIResponse 将 TaobaoItemSkuAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoItemSkuAddAPIResponse(v *TaobaoItemSkuAddAPIResponse) {
	v.Reset()
	poolTaobaoItemSkuAddAPIResponse.Put(v)
}
