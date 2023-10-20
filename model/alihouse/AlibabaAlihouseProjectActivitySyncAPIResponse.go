package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseprojectactivitysyncAPIResponse 电商券数据同步 API返回值
// alibaba.alihouse.project.activity.sync
//
// 电商券数据同步
type AlibabaalihouseprojectactivitysyncAPIResponse struct {
	model.CommonResponse
	AlibabaalihouseprojectactivitysyncAPIResponseModel
}

// AlibabaalihouseprojectactivitysyncAPIResponseModel is 电商券数据同步 成功返回结果
type AlibabaalihouseprojectactivitysyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_project_activity_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaalihouseprojectactivitysyncResult `json:"result,omitempty" xml:"result,omitempty"`
}
