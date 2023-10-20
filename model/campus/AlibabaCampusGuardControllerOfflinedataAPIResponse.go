package campus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusguardcontrollerofflinedataAPIResponse 点位离线数据拉取 API返回值
// alibaba.campus.guard.controller.offlinedata
//
// 点位离线数据拉取
type AlibabacampusguardcontrollerofflinedataAPIResponse struct {
	model.CommonResponse
	AlibabacampusguardcontrollerofflinedataAPIResponseModel
}

// AlibabacampusguardcontrollerofflinedataAPIResponseModel is 点位离线数据拉取 成功返回结果
type AlibabacampusguardcontrollerofflinedataAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_guard_controller_offlinedata_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}
