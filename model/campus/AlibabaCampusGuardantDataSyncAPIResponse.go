package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusGuardantDataSyncAPIResponse 刷卡数据同步 API返回值
// alibaba.campus.guardant.data.sync
//
// 数据同步门禁系统
type AlibabaCampusGuardantDataSyncAPIResponse struct {
	model.CommonResponse
	AlibabaCampusGuardantDataSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusGuardantDataSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusGuardantDataSyncAPIResponseModel).Reset()
}

// AlibabaCampusGuardantDataSyncAPIResponseModel is 刷卡数据同步 成功返回结果
type AlibabaCampusGuardantDataSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_guardant_data_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusGuardantDataSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusGuardantDataSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusGuardantDataSyncAPIResponse)
	},
}

// GetAlibabaCampusGuardantDataSyncAPIResponse 从 sync.Pool 获取 AlibabaCampusGuardantDataSyncAPIResponse
func GetAlibabaCampusGuardantDataSyncAPIResponse() *AlibabaCampusGuardantDataSyncAPIResponse {
	return poolAlibabaCampusGuardantDataSyncAPIResponse.Get().(*AlibabaCampusGuardantDataSyncAPIResponse)
}

// ReleaseAlibabaCampusGuardantDataSyncAPIResponse 将 AlibabaCampusGuardantDataSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusGuardantDataSyncAPIResponse(v *AlibabaCampusGuardantDataSyncAPIResponse) {
	v.Reset()
	poolAlibabaCampusGuardantDataSyncAPIResponse.Put(v)
}
