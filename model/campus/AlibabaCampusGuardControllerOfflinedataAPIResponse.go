package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusGuardControllerOfflinedataAPIResponse 点位离线数据拉取 API返回值
// alibaba.campus.guard.controller.offlinedata
//
// 点位离线数据拉取
type AlibabaCampusGuardControllerOfflinedataAPIResponse struct {
	model.CommonResponse
	AlibabaCampusGuardControllerOfflinedataAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusGuardControllerOfflinedataAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusGuardControllerOfflinedataAPIResponseModel).Reset()
}

// AlibabaCampusGuardControllerOfflinedataAPIResponseModel is 点位离线数据拉取 成功返回结果
type AlibabaCampusGuardControllerOfflinedataAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_guard_controller_offlinedata_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusGuardControllerOfflinedataAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusGuardControllerOfflinedataAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusGuardControllerOfflinedataAPIResponse)
	},
}

// GetAlibabaCampusGuardControllerOfflinedataAPIResponse 从 sync.Pool 获取 AlibabaCampusGuardControllerOfflinedataAPIResponse
func GetAlibabaCampusGuardControllerOfflinedataAPIResponse() *AlibabaCampusGuardControllerOfflinedataAPIResponse {
	return poolAlibabaCampusGuardControllerOfflinedataAPIResponse.Get().(*AlibabaCampusGuardControllerOfflinedataAPIResponse)
}

// ReleaseAlibabaCampusGuardControllerOfflinedataAPIResponse 将 AlibabaCampusGuardControllerOfflinedataAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusGuardControllerOfflinedataAPIResponse(v *AlibabaCampusGuardControllerOfflinedataAPIResponse) {
	v.Reset()
	poolAlibabaCampusGuardControllerOfflinedataAPIResponse.Put(v)
}
