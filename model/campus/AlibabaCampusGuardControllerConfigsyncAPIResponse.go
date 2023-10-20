package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusGuardControllerConfigsyncAPIResponse 门禁控制器配置项同步 API返回值
// alibaba.campus.guard.controller.configsync
//
// 门禁控制器配置项同步
type AlibabaCampusGuardControllerConfigsyncAPIResponse struct {
	model.CommonResponse
	AlibabaCampusGuardControllerConfigsyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusGuardControllerConfigsyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusGuardControllerConfigsyncAPIResponseModel).Reset()
}

// AlibabaCampusGuardControllerConfigsyncAPIResponseModel is 门禁控制器配置项同步 成功返回结果
type AlibabaCampusGuardControllerConfigsyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_guard_controller_configsync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果对象
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusGuardControllerConfigsyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusGuardControllerConfigsyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusGuardControllerConfigsyncAPIResponse)
	},
}

// GetAlibabaCampusGuardControllerConfigsyncAPIResponse 从 sync.Pool 获取 AlibabaCampusGuardControllerConfigsyncAPIResponse
func GetAlibabaCampusGuardControllerConfigsyncAPIResponse() *AlibabaCampusGuardControllerConfigsyncAPIResponse {
	return poolAlibabaCampusGuardControllerConfigsyncAPIResponse.Get().(*AlibabaCampusGuardControllerConfigsyncAPIResponse)
}

// ReleaseAlibabaCampusGuardControllerConfigsyncAPIResponse 将 AlibabaCampusGuardControllerConfigsyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusGuardControllerConfigsyncAPIResponse(v *AlibabaCampusGuardControllerConfigsyncAPIResponse) {
	v.Reset()
	poolAlibabaCampusGuardControllerConfigsyncAPIResponse.Put(v)
}
