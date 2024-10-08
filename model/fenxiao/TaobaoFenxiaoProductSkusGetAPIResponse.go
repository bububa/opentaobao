package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoProductSkusGetAPIResponse SKU查询接口 API返回值
// taobao.fenxiao.product.skus.get
//
// 产品sku查询
type TaobaoFenxiaoProductSkusGetAPIResponse struct {
	model.CommonResponse
	TaobaoFenxiaoProductSkusGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFenxiaoProductSkusGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFenxiaoProductSkusGetAPIResponseModel).Reset()
}

// TaobaoFenxiaoProductSkusGetAPIResponseModel is SKU查询接口 成功返回结果
type TaobaoFenxiaoProductSkusGetAPIResponseModel struct {
	XMLName xml.Name `xml:"fenxiao_product_skus_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// sku信息
	Skus []FenxiaoSku `json:"skus,omitempty" xml:"skus>fenxiao_sku,omitempty"`
	// 记录数量
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFenxiaoProductSkusGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Skus = m.Skus[:0]
	m.TotalResults = 0
}

var poolTaobaoFenxiaoProductSkusGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFenxiaoProductSkusGetAPIResponse)
	},
}

// GetTaobaoFenxiaoProductSkusGetAPIResponse 从 sync.Pool 获取 TaobaoFenxiaoProductSkusGetAPIResponse
func GetTaobaoFenxiaoProductSkusGetAPIResponse() *TaobaoFenxiaoProductSkusGetAPIResponse {
	return poolTaobaoFenxiaoProductSkusGetAPIResponse.Get().(*TaobaoFenxiaoProductSkusGetAPIResponse)
}

// ReleaseTaobaoFenxiaoProductSkusGetAPIResponse 将 TaobaoFenxiaoProductSkusGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFenxiaoProductSkusGetAPIResponse(v *TaobaoFenxiaoProductSkusGetAPIResponse) {
	v.Reset()
	poolTaobaoFenxiaoProductSkusGetAPIResponse.Put(v)
}
