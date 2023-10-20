package cainiaohandover

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoGlobalHandoverCancelAPIResponse 取消交接单 API返回值
// cainiao.global.handover.cancel
//
// 提供给ISV通过该接口取消交接单
type CainiaoGlobalHandoverCancelAPIResponse struct {
	model.CommonResponse
	CainiaoGlobalHandoverCancelAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoGlobalHandoverCancelAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoGlobalHandoverCancelAPIResponseModel).Reset()
}

// CainiaoGlobalHandoverCancelAPIResponseModel is 取消交接单 成功返回结果
type CainiaoGlobalHandoverCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_global_handover_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求结果
	Result *HsfResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoGlobalHandoverCancelAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolCainiaoGlobalHandoverCancelAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoGlobalHandoverCancelAPIResponse)
	},
}

// GetCainiaoGlobalHandoverCancelAPIResponse 从 sync.Pool 获取 CainiaoGlobalHandoverCancelAPIResponse
func GetCainiaoGlobalHandoverCancelAPIResponse() *CainiaoGlobalHandoverCancelAPIResponse {
	return poolCainiaoGlobalHandoverCancelAPIResponse.Get().(*CainiaoGlobalHandoverCancelAPIResponse)
}

// ReleaseCainiaoGlobalHandoverCancelAPIResponse 将 CainiaoGlobalHandoverCancelAPIResponse 保存到 sync.Pool
func ReleaseCainiaoGlobalHandoverCancelAPIResponse(v *CainiaoGlobalHandoverCancelAPIResponse) {
	v.Reset()
	poolCainiaoGlobalHandoverCancelAPIResponse.Put(v)
}
