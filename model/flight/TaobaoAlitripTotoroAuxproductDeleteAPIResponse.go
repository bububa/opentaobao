package flight

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTotoroAuxproductDeleteAPIResponse 廉航辅营产品删除 API返回值
// taobao.alitrip.totoro.auxproduct.delete
//
// 廉航辅营产品删除接口
type TaobaoAlitripTotoroAuxproductDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTotoroAuxproductDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripTotoroAuxproductDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripTotoroAuxproductDeleteAPIResponseModel).Reset()
}

// TaobaoAlitripTotoroAuxproductDeleteAPIResponseModel is 廉航辅营产品删除 成功返回结果
type TaobaoAlitripTotoroAuxproductDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_totoro_auxproduct_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *DelAuxProductsRs `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripTotoroAuxproductDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAlitripTotoroAuxproductDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTotoroAuxproductDeleteAPIResponse)
	},
}

// GetTaobaoAlitripTotoroAuxproductDeleteAPIResponse 从 sync.Pool 获取 TaobaoAlitripTotoroAuxproductDeleteAPIResponse
func GetTaobaoAlitripTotoroAuxproductDeleteAPIResponse() *TaobaoAlitripTotoroAuxproductDeleteAPIResponse {
	return poolTaobaoAlitripTotoroAuxproductDeleteAPIResponse.Get().(*TaobaoAlitripTotoroAuxproductDeleteAPIResponse)
}

// ReleaseTaobaoAlitripTotoroAuxproductDeleteAPIResponse 将 TaobaoAlitripTotoroAuxproductDeleteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripTotoroAuxproductDeleteAPIResponse(v *TaobaoAlitripTotoroAuxproductDeleteAPIResponse) {
	v.Reset()
	poolTaobaoAlitripTotoroAuxproductDeleteAPIResponse.Put(v)
}
