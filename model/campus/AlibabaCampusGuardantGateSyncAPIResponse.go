package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusGuardantGateSyncAPIResponse 网点数据同步 API返回值
// alibaba.campus.guardant.gate.sync
//
// 门禁供应商创建网点同步
type AlibabaCampusGuardantGateSyncAPIResponse struct {
	model.CommonResponse
	AlibabaCampusGuardantGateSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusGuardantGateSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusGuardantGateSyncAPIResponseModel).Reset()
}

// AlibabaCampusGuardantGateSyncAPIResponseModel is 网点数据同步 成功返回结果
type AlibabaCampusGuardantGateSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_guardant_gate_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusGuardantGateSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusGuardantGateSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusGuardantGateSyncAPIResponse)
	},
}

// GetAlibabaCampusGuardantGateSyncAPIResponse 从 sync.Pool 获取 AlibabaCampusGuardantGateSyncAPIResponse
func GetAlibabaCampusGuardantGateSyncAPIResponse() *AlibabaCampusGuardantGateSyncAPIResponse {
	return poolAlibabaCampusGuardantGateSyncAPIResponse.Get().(*AlibabaCampusGuardantGateSyncAPIResponse)
}

// ReleaseAlibabaCampusGuardantGateSyncAPIResponse 将 AlibabaCampusGuardantGateSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusGuardantGateSyncAPIResponse(v *AlibabaCampusGuardantGateSyncAPIResponse) {
	v.Reset()
	poolAlibabaCampusGuardantGateSyncAPIResponse.Put(v)
}
