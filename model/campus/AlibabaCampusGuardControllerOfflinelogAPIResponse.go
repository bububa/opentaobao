package campus

import (
	"encoding/xml"

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

// AlibabaCampusGuardControllerOfflinelogAPIResponseModel is 门禁控制器离线日志同步 成功返回结果
type AlibabaCampusGuardControllerOfflinelogAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_guard_controller_offlinelog_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}
