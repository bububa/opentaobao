package waybill

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoWaybillIiUpdateAPIResponse 电子面单云打印更新接口 API返回值
// cainiao.waybill.ii.update
//
// 商家更新电子面单号对应的面单信息。
type CainiaoWaybillIiUpdateAPIResponse struct {
	model.CommonResponse
	CainiaoWaybillIiUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoWaybillIiUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoWaybillIiUpdateAPIResponseModel).Reset()
}

// CainiaoWaybillIiUpdateAPIResponseModel is 电子面单云打印更新接口 成功返回结果
type CainiaoWaybillIiUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_waybill_ii_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 模板内容
	PrintData string `json:"print_data,omitempty" xml:"print_data,omitempty"`
	// 面单号
	WaybillCode string `json:"waybill_code,omitempty" xml:"waybill_code,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoWaybillIiUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.PrintData = ""
	m.WaybillCode = ""
}

var poolCainiaoWaybillIiUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoWaybillIiUpdateAPIResponse)
	},
}

// GetCainiaoWaybillIiUpdateAPIResponse 从 sync.Pool 获取 CainiaoWaybillIiUpdateAPIResponse
func GetCainiaoWaybillIiUpdateAPIResponse() *CainiaoWaybillIiUpdateAPIResponse {
	return poolCainiaoWaybillIiUpdateAPIResponse.Get().(*CainiaoWaybillIiUpdateAPIResponse)
}

// ReleaseCainiaoWaybillIiUpdateAPIResponse 将 CainiaoWaybillIiUpdateAPIResponse 保存到 sync.Pool
func ReleaseCainiaoWaybillIiUpdateAPIResponse(v *CainiaoWaybillIiUpdateAPIResponse) {
	v.Reset()
	poolCainiaoWaybillIiUpdateAPIResponse.Put(v)
}
