package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoProductSkuAddAPIResponse 产品sku添加接口 API返回值
// taobao.fenxiao.product.sku.add
//
// 添加产品SKU信息
type TaobaoFenxiaoProductSkuAddAPIResponse struct {
	model.CommonResponse
	TaobaoFenxiaoProductSkuAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFenxiaoProductSkuAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFenxiaoProductSkuAddAPIResponseModel).Reset()
}

// TaobaoFenxiaoProductSkuAddAPIResponseModel is 产品sku添加接口 成功返回结果
type TaobaoFenxiaoProductSkuAddAPIResponseModel struct {
	XMLName xml.Name `xml:"fenxiao_product_sku_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作时间
	Created string `json:"created,omitempty" xml:"created,omitempty"`
	// 操作结果
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFenxiaoProductSkuAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Created = ""
	m.Result = false
}

var poolTaobaoFenxiaoProductSkuAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFenxiaoProductSkuAddAPIResponse)
	},
}

// GetTaobaoFenxiaoProductSkuAddAPIResponse 从 sync.Pool 获取 TaobaoFenxiaoProductSkuAddAPIResponse
func GetTaobaoFenxiaoProductSkuAddAPIResponse() *TaobaoFenxiaoProductSkuAddAPIResponse {
	return poolTaobaoFenxiaoProductSkuAddAPIResponse.Get().(*TaobaoFenxiaoProductSkuAddAPIResponse)
}

// ReleaseTaobaoFenxiaoProductSkuAddAPIResponse 将 TaobaoFenxiaoProductSkuAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFenxiaoProductSkuAddAPIResponse(v *TaobaoFenxiaoProductSkuAddAPIResponse) {
	v.Reset()
	poolTaobaoFenxiaoProductSkuAddAPIResponse.Put(v)
}
