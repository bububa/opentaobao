package cainiaohandover

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoGlobalHandoverSavedraftAPIResponse 创建交接单草稿 API返回值
// cainiao.global.handover.savedraft
//
// 提供给ISV通过该接口创建交接单草稿
type CainiaoGlobalHandoverSavedraftAPIResponse struct {
	model.CommonResponse
	CainiaoGlobalHandoverSavedraftAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoGlobalHandoverSavedraftAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoGlobalHandoverSavedraftAPIResponseModel).Reset()
}

// CainiaoGlobalHandoverSavedraftAPIResponseModel is 创建交接单草稿 成功返回结果
type CainiaoGlobalHandoverSavedraftAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_global_handover_savedraft_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求结果
	Result *HsfResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoGlobalHandoverSavedraftAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolCainiaoGlobalHandoverSavedraftAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoGlobalHandoverSavedraftAPIResponse)
	},
}

// GetCainiaoGlobalHandoverSavedraftAPIResponse 从 sync.Pool 获取 CainiaoGlobalHandoverSavedraftAPIResponse
func GetCainiaoGlobalHandoverSavedraftAPIResponse() *CainiaoGlobalHandoverSavedraftAPIResponse {
	return poolCainiaoGlobalHandoverSavedraftAPIResponse.Get().(*CainiaoGlobalHandoverSavedraftAPIResponse)
}

// ReleaseCainiaoGlobalHandoverSavedraftAPIResponse 将 CainiaoGlobalHandoverSavedraftAPIResponse 保存到 sync.Pool
func ReleaseCainiaoGlobalHandoverSavedraftAPIResponse(v *CainiaoGlobalHandoverSavedraftAPIResponse) {
	v.Reset()
	poolCainiaoGlobalHandoverSavedraftAPIResponse.Put(v)
}
