package cainiaohandover

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoGlobalHandoverContentSubbagAddAPIResponse 预约单下追加大包 API返回值
// cainiao.global.handover.content.subbag.add
//
// 预约单下追加大包
type CainiaoGlobalHandoverContentSubbagAddAPIResponse struct {
	model.CommonResponse
	CainiaoGlobalHandoverContentSubbagAddAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoGlobalHandoverContentSubbagAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoGlobalHandoverContentSubbagAddAPIResponseModel).Reset()
}

// CainiaoGlobalHandoverContentSubbagAddAPIResponseModel is 预约单下追加大包 成功返回结果
type CainiaoGlobalHandoverContentSubbagAddAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_global_handover_content_subbag_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求响应
	Result *HsfResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoGlobalHandoverContentSubbagAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolCainiaoGlobalHandoverContentSubbagAddAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoGlobalHandoverContentSubbagAddAPIResponse)
	},
}

// GetCainiaoGlobalHandoverContentSubbagAddAPIResponse 从 sync.Pool 获取 CainiaoGlobalHandoverContentSubbagAddAPIResponse
func GetCainiaoGlobalHandoverContentSubbagAddAPIResponse() *CainiaoGlobalHandoverContentSubbagAddAPIResponse {
	return poolCainiaoGlobalHandoverContentSubbagAddAPIResponse.Get().(*CainiaoGlobalHandoverContentSubbagAddAPIResponse)
}

// ReleaseCainiaoGlobalHandoverContentSubbagAddAPIResponse 将 CainiaoGlobalHandoverContentSubbagAddAPIResponse 保存到 sync.Pool
func ReleaseCainiaoGlobalHandoverContentSubbagAddAPIResponse(v *CainiaoGlobalHandoverContentSubbagAddAPIResponse) {
	v.Reset()
	poolCainiaoGlobalHandoverContentSubbagAddAPIResponse.Put(v)
}
