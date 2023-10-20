package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoProductSkuUpdateAPIResponse 产品sku编辑接口 API返回值
// taobao.fenxiao.product.sku.update
//
// 产品SKU信息更新
type TaobaoFenxiaoProductSkuUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoFenxiaoProductSkuUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFenxiaoProductSkuUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFenxiaoProductSkuUpdateAPIResponseModel).Reset()
}

// TaobaoFenxiaoProductSkuUpdateAPIResponseModel is 产品sku编辑接口 成功返回结果
type TaobaoFenxiaoProductSkuUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"fenxiao_product_sku_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作时间
	Created string `json:"created,omitempty" xml:"created,omitempty"`
	// 操作结果
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFenxiaoProductSkuUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Created = ""
	m.Result = false
}

var poolTaobaoFenxiaoProductSkuUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFenxiaoProductSkuUpdateAPIResponse)
	},
}

// GetTaobaoFenxiaoProductSkuUpdateAPIResponse 从 sync.Pool 获取 TaobaoFenxiaoProductSkuUpdateAPIResponse
func GetTaobaoFenxiaoProductSkuUpdateAPIResponse() *TaobaoFenxiaoProductSkuUpdateAPIResponse {
	return poolTaobaoFenxiaoProductSkuUpdateAPIResponse.Get().(*TaobaoFenxiaoProductSkuUpdateAPIResponse)
}

// ReleaseTaobaoFenxiaoProductSkuUpdateAPIResponse 将 TaobaoFenxiaoProductSkuUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFenxiaoProductSkuUpdateAPIResponse(v *TaobaoFenxiaoProductSkuUpdateAPIResponse) {
	v.Reset()
	poolTaobaoFenxiaoProductSkuUpdateAPIResponse.Put(v)
}
