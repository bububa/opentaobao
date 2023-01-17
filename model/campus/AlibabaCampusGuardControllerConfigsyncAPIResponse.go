package campus

import (
	"encoding/xml"

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

// AlibabaCampusGuardControllerConfigsyncAPIResponseModel is 门禁控制器配置项同步 成功返回结果
type AlibabaCampusGuardControllerConfigsyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_guard_controller_configsync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果对象
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}
