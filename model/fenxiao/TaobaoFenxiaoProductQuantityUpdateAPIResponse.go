package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoProductQuantityUpdateAPIResponse 产品库存修改 API返回值
// taobao.fenxiao.product.quantity.update
//
// 修改产品库存信息，支持全量修改以及增量修改两种方式
type TaobaoFenxiaoProductQuantityUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoFenxiaoProductQuantityUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFenxiaoProductQuantityUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFenxiaoProductQuantityUpdateAPIResponseModel).Reset()
}

// TaobaoFenxiaoProductQuantityUpdateAPIResponseModel is 产品库存修改 成功返回结果
type TaobaoFenxiaoProductQuantityUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"fenxiao_product_quantity_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作时间
	Created string `json:"created,omitempty" xml:"created,omitempty"`
	// 操作结果
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFenxiaoProductQuantityUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Created = ""
	m.Result = false
}

var poolTaobaoFenxiaoProductQuantityUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFenxiaoProductQuantityUpdateAPIResponse)
	},
}

// GetTaobaoFenxiaoProductQuantityUpdateAPIResponse 从 sync.Pool 获取 TaobaoFenxiaoProductQuantityUpdateAPIResponse
func GetTaobaoFenxiaoProductQuantityUpdateAPIResponse() *TaobaoFenxiaoProductQuantityUpdateAPIResponse {
	return poolTaobaoFenxiaoProductQuantityUpdateAPIResponse.Get().(*TaobaoFenxiaoProductQuantityUpdateAPIResponse)
}

// ReleaseTaobaoFenxiaoProductQuantityUpdateAPIResponse 将 TaobaoFenxiaoProductQuantityUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFenxiaoProductQuantityUpdateAPIResponse(v *TaobaoFenxiaoProductQuantityUpdateAPIResponse) {
	v.Reset()
	poolTaobaoFenxiaoProductQuantityUpdateAPIResponse.Put(v)
}
