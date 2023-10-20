package tmallnr

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabalsycrmactivitypageupdateAPIResponse ISV活动页面创建修改 API返回值
// alibaba.lsy.crm.activity.page.update
//
// ISV活动页面创建修改
type AlibabalsycrmactivitypageupdateAPIResponse struct {
	model.CommonResponse
	AlibabalsycrmactivitypageupdateAPIResponseModel
}

// AlibabalsycrmactivitypageupdateAPIResponseModel is ISV活动页面创建修改 成功返回结果
type AlibabalsycrmactivitypageupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lsy_crm_activity_page_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabalsycrmactivitypageupdateResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
