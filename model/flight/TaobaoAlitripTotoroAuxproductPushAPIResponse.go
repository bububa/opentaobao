package flight

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTotoroAuxproductPushAPIResponse 廉航辅营产品投放 API返回值
// taobao.alitrip.totoro.auxproduct.push
//
// 廉航辅营产品投放接口
type TaobaoAlitripTotoroAuxproductPushAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTotoroAuxproductPushAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripTotoroAuxproductPushAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripTotoroAuxproductPushAPIResponseModel).Reset()
}

// TaobaoAlitripTotoroAuxproductPushAPIResponseModel is 廉航辅营产品投放 成功返回结果
type TaobaoAlitripTotoroAuxproductPushAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_totoro_auxproduct_push_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作日志id，商家可通过该id在后台查看本次操作的具体结果
	TracerId string `json:"tracer_id,omitempty" xml:"tracer_id,omitempty"`
	// 备注
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 是否操作成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripTotoroAuxproductPushAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TracerId = ""
	m.Message = ""
	m.IsSuccess = false
}

var poolTaobaoAlitripTotoroAuxproductPushAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTotoroAuxproductPushAPIResponse)
	},
}

// GetTaobaoAlitripTotoroAuxproductPushAPIResponse 从 sync.Pool 获取 TaobaoAlitripTotoroAuxproductPushAPIResponse
func GetTaobaoAlitripTotoroAuxproductPushAPIResponse() *TaobaoAlitripTotoroAuxproductPushAPIResponse {
	return poolTaobaoAlitripTotoroAuxproductPushAPIResponse.Get().(*TaobaoAlitripTotoroAuxproductPushAPIResponse)
}

// ReleaseTaobaoAlitripTotoroAuxproductPushAPIResponse 将 TaobaoAlitripTotoroAuxproductPushAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripTotoroAuxproductPushAPIResponse(v *TaobaoAlitripTotoroAuxproductPushAPIResponse) {
	v.Reset()
	poolTaobaoAlitripTotoroAuxproductPushAPIResponse.Put(v)
}
