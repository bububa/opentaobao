package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusGuardDataSyncAPIResponse 卡巴数据同步 API返回值
// alibaba.campus.guard.data.sync
//
// 数据同步门禁系统
type AlibabaCampusGuardDataSyncAPIResponse struct {
	model.CommonResponse
	AlibabaCampusGuardDataSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusGuardDataSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusGuardDataSyncAPIResponseModel).Reset()
}

// AlibabaCampusGuardDataSyncAPIResponseModel is 卡巴数据同步 成功返回结果
type AlibabaCampusGuardDataSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_guard_data_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusGuardDataSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusGuardDataSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusGuardDataSyncAPIResponse)
	},
}

// GetAlibabaCampusGuardDataSyncAPIResponse 从 sync.Pool 获取 AlibabaCampusGuardDataSyncAPIResponse
func GetAlibabaCampusGuardDataSyncAPIResponse() *AlibabaCampusGuardDataSyncAPIResponse {
	return poolAlibabaCampusGuardDataSyncAPIResponse.Get().(*AlibabaCampusGuardDataSyncAPIResponse)
}

// ReleaseAlibabaCampusGuardDataSyncAPIResponse 将 AlibabaCampusGuardDataSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusGuardDataSyncAPIResponse(v *AlibabaCampusGuardDataSyncAPIResponse) {
	v.Reset()
	poolAlibabaCampusGuardDataSyncAPIResponse.Put(v)
}
