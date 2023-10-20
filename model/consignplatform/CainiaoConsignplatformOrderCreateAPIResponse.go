package consignplatform

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoConsignplatformOrderCreateAPIResponse 菜鸟发货工作台创建订单 API返回值
// cainiao.consignplatform.order.create
//
// 菜鸟发货工作台，商家或者isv通过api进行订单写入操作
type CainiaoConsignplatformOrderCreateAPIResponse struct {
	model.CommonResponse
	CainiaoConsignplatformOrderCreateAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoConsignplatformOrderCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoConsignplatformOrderCreateAPIResponseModel).Reset()
}

// CainiaoConsignplatformOrderCreateAPIResponseModel is 菜鸟发货工作台创建订单 成功返回结果
type CainiaoConsignplatformOrderCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_consignplatform_order_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 失败信息
	FailMessage string `json:"fail_message,omitempty" xml:"fail_message,omitempty"`
	// 失败code
	FailCode string `json:"fail_code,omitempty" xml:"fail_code,omitempty"`
	// 创建是否成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoConsignplatformOrderCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.FailMessage = ""
	m.FailCode = ""
	m.Result = false
}

var poolCainiaoConsignplatformOrderCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoConsignplatformOrderCreateAPIResponse)
	},
}

// GetCainiaoConsignplatformOrderCreateAPIResponse 从 sync.Pool 获取 CainiaoConsignplatformOrderCreateAPIResponse
func GetCainiaoConsignplatformOrderCreateAPIResponse() *CainiaoConsignplatformOrderCreateAPIResponse {
	return poolCainiaoConsignplatformOrderCreateAPIResponse.Get().(*CainiaoConsignplatformOrderCreateAPIResponse)
}

// ReleaseCainiaoConsignplatformOrderCreateAPIResponse 将 CainiaoConsignplatformOrderCreateAPIResponse 保存到 sync.Pool
func ReleaseCainiaoConsignplatformOrderCreateAPIResponse(v *CainiaoConsignplatformOrderCreateAPIResponse) {
	v.Reset()
	poolCainiaoConsignplatformOrderCreateAPIResponse.Put(v)
}
