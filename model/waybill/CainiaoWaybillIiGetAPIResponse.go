package waybill

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoWaybillIiGetAPIResponse 电子面单云打印接口 API返回值
// cainiao.waybill.ii.get
//
// 菜鸟电子面单的云打印申请电子面单号的方法
type CainiaoWaybillIiGetAPIResponse struct {
	model.CommonResponse
	CainiaoWaybillIiGetAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoWaybillIiGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoWaybillIiGetAPIResponseModel).Reset()
}

// CainiaoWaybillIiGetAPIResponseModel is 电子面单云打印接口 成功返回结果
type CainiaoWaybillIiGetAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_waybill_ii_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Modules []WaybillCloudPrintResponse `json:"modules,omitempty" xml:"modules>waybill_cloud_print_response,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoWaybillIiGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Modules = m.Modules[:0]
}

var poolCainiaoWaybillIiGetAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoWaybillIiGetAPIResponse)
	},
}

// GetCainiaoWaybillIiGetAPIResponse 从 sync.Pool 获取 CainiaoWaybillIiGetAPIResponse
func GetCainiaoWaybillIiGetAPIResponse() *CainiaoWaybillIiGetAPIResponse {
	return poolCainiaoWaybillIiGetAPIResponse.Get().(*CainiaoWaybillIiGetAPIResponse)
}

// ReleaseCainiaoWaybillIiGetAPIResponse 将 CainiaoWaybillIiGetAPIResponse 保存到 sync.Pool
func ReleaseCainiaoWaybillIiGetAPIResponse(v *CainiaoWaybillIiGetAPIResponse) {
	v.Reset()
	poolCainiaoWaybillIiGetAPIResponse.Put(v)
}
