package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoProductSkuDeleteAPIResponse 产品SKU删除接口 API返回值
// taobao.fenxiao.product.sku.delete
//
// 根据sku properties删除sku数据
type TaobaoFenxiaoProductSkuDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoFenxiaoProductSkuDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFenxiaoProductSkuDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFenxiaoProductSkuDeleteAPIResponseModel).Reset()
}

// TaobaoFenxiaoProductSkuDeleteAPIResponseModel is 产品SKU删除接口 成功返回结果
type TaobaoFenxiaoProductSkuDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"fenxiao_product_sku_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作时间
	Created string `json:"created,omitempty" xml:"created,omitempty"`
	// 操作结果
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFenxiaoProductSkuDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Created = ""
	m.Result = false
}

var poolTaobaoFenxiaoProductSkuDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFenxiaoProductSkuDeleteAPIResponse)
	},
}

// GetTaobaoFenxiaoProductSkuDeleteAPIResponse 从 sync.Pool 获取 TaobaoFenxiaoProductSkuDeleteAPIResponse
func GetTaobaoFenxiaoProductSkuDeleteAPIResponse() *TaobaoFenxiaoProductSkuDeleteAPIResponse {
	return poolTaobaoFenxiaoProductSkuDeleteAPIResponse.Get().(*TaobaoFenxiaoProductSkuDeleteAPIResponse)
}

// ReleaseTaobaoFenxiaoProductSkuDeleteAPIResponse 将 TaobaoFenxiaoProductSkuDeleteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFenxiaoProductSkuDeleteAPIResponse(v *TaobaoFenxiaoProductSkuDeleteAPIResponse) {
	v.Reset()
	poolTaobaoFenxiaoProductSkuDeleteAPIResponse.Put(v)
}
