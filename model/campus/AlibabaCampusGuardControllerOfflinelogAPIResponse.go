package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusGuardControllerOfflinelogAPIResponse 门禁控制器离线日志同步 API返回值
// alibaba.campus.guard.controller.offlinelog
//
// 门禁控制器离线日志同步
type AlibabaCampusGuardControllerOfflinelogAPIResponse struct {
	model.CommonResponse
	AlibabaCampusGuardControllerOfflinelogAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusGuardControllerOfflinelogAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusGuardControllerOfflinelogAPIResponseModel).Reset()
}

// AlibabaCampusGuardControllerOfflinelogAPIResponseModel is 门禁控制器离线日志同步 成功返回结果
type AlibabaCampusGuardControllerOfflinelogAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_guard_controller_offlinelog_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusGuardControllerOfflinelogAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusGuardControllerOfflinelogAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusGuardControllerOfflinelogAPIResponse)
	},
}

// GetAlibabaCampusGuardControllerOfflinelogAPIResponse 从 sync.Pool 获取 AlibabaCampusGuardControllerOfflinelogAPIResponse
func GetAlibabaCampusGuardControllerOfflinelogAPIResponse() *AlibabaCampusGuardControllerOfflinelogAPIResponse {
	return poolAlibabaCampusGuardControllerOfflinelogAPIResponse.Get().(*AlibabaCampusGuardControllerOfflinelogAPIResponse)
}

// ReleaseAlibabaCampusGuardControllerOfflinelogAPIResponse 将 AlibabaCampusGuardControllerOfflinelogAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusGuardControllerOfflinelogAPIResponse(v *AlibabaCampusGuardControllerOfflinelogAPIResponse) {
	v.Reset()
	poolAlibabaCampusGuardControllerOfflinelogAPIResponse.Put(v)
}
