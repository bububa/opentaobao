package tmallnr

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabalsycrmactivityupdateAPIResponse ISV活动修改 API返回值
// alibaba.lsy.crm.activity.update
//
// ISV活动修改
type AlibabalsycrmactivityupdateAPIResponse struct {
	model.CommonResponse
	AlibabalsycrmactivityupdateAPIResponseModel
}

// AlibabalsycrmactivityupdateAPIResponseModel is ISV活动修改 成功返回结果
type AlibabalsycrmactivityupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lsy_crm_activity_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabalsycrmactivityupdateResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
