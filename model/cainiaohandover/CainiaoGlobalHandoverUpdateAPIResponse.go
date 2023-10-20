package cainiaohandover

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoGlobalHandoverUpdateAPIResponse 修改交接单 API返回值
// cainiao.global.handover.update
//
// 提供给ISV通过该接口修改交接单
type CainiaoGlobalHandoverUpdateAPIResponse struct {
	model.CommonResponse
	CainiaoGlobalHandoverUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoGlobalHandoverUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoGlobalHandoverUpdateAPIResponseModel).Reset()
}

// CainiaoGlobalHandoverUpdateAPIResponseModel is 修改交接单 成功返回结果
type CainiaoGlobalHandoverUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_global_handover_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求结果
	Result *HsfResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoGlobalHandoverUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolCainiaoGlobalHandoverUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoGlobalHandoverUpdateAPIResponse)
	},
}

// GetCainiaoGlobalHandoverUpdateAPIResponse 从 sync.Pool 获取 CainiaoGlobalHandoverUpdateAPIResponse
func GetCainiaoGlobalHandoverUpdateAPIResponse() *CainiaoGlobalHandoverUpdateAPIResponse {
	return poolCainiaoGlobalHandoverUpdateAPIResponse.Get().(*CainiaoGlobalHandoverUpdateAPIResponse)
}

// ReleaseCainiaoGlobalHandoverUpdateAPIResponse 将 CainiaoGlobalHandoverUpdateAPIResponse 保存到 sync.Pool
func ReleaseCainiaoGlobalHandoverUpdateAPIResponse(v *CainiaoGlobalHandoverUpdateAPIResponse) {
	v.Reset()
	poolCainiaoGlobalHandoverUpdateAPIResponse.Put(v)
}
