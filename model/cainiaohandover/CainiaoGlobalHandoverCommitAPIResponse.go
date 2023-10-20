package cainiaohandover

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoGlobalHandoverCommitAPIResponse 提交发布交接单 API返回值
// cainiao.global.handover.commit
//
// 提供给ISV通过该接口提交发布交接单
type CainiaoGlobalHandoverCommitAPIResponse struct {
	model.CommonResponse
	CainiaoGlobalHandoverCommitAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoGlobalHandoverCommitAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoGlobalHandoverCommitAPIResponseModel).Reset()
}

// CainiaoGlobalHandoverCommitAPIResponseModel is 提交发布交接单 成功返回结果
type CainiaoGlobalHandoverCommitAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_global_handover_commit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求结果
	Result *HsfResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoGlobalHandoverCommitAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolCainiaoGlobalHandoverCommitAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoGlobalHandoverCommitAPIResponse)
	},
}

// GetCainiaoGlobalHandoverCommitAPIResponse 从 sync.Pool 获取 CainiaoGlobalHandoverCommitAPIResponse
func GetCainiaoGlobalHandoverCommitAPIResponse() *CainiaoGlobalHandoverCommitAPIResponse {
	return poolCainiaoGlobalHandoverCommitAPIResponse.Get().(*CainiaoGlobalHandoverCommitAPIResponse)
}

// ReleaseCainiaoGlobalHandoverCommitAPIResponse 将 CainiaoGlobalHandoverCommitAPIResponse 保存到 sync.Pool
func ReleaseCainiaoGlobalHandoverCommitAPIResponse(v *CainiaoGlobalHandoverCommitAPIResponse) {
	v.Reset()
	poolCainiaoGlobalHandoverCommitAPIResponse.Put(v)
}
