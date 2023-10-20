package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterFulfiltaskInsuranceActionAPIResponse 供应链保险链路动作 API返回值
// tmall.servicecenter.fulfiltask.insurance.action
//
// 服务供应链履约链路 保险类业务履约接口
type TmallServicecenterFulfiltaskInsuranceActionAPIResponse struct {
	model.CommonResponse
	TmallServicecenterFulfiltaskInsuranceActionAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterFulfiltaskInsuranceActionAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterFulfiltaskInsuranceActionAPIResponseModel).Reset()
}

// TmallServicecenterFulfiltaskInsuranceActionAPIResponseModel is 供应链保险链路动作 成功返回结果
type TmallServicecenterFulfiltaskInsuranceActionAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_fulfiltask_insurance_action_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TmallServicecenterFulfiltaskInsuranceActionResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterFulfiltaskInsuranceActionAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServicecenterFulfiltaskInsuranceActionAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterFulfiltaskInsuranceActionAPIResponse)
	},
}

// GetTmallServicecenterFulfiltaskInsuranceActionAPIResponse 从 sync.Pool 获取 TmallServicecenterFulfiltaskInsuranceActionAPIResponse
func GetTmallServicecenterFulfiltaskInsuranceActionAPIResponse() *TmallServicecenterFulfiltaskInsuranceActionAPIResponse {
	return poolTmallServicecenterFulfiltaskInsuranceActionAPIResponse.Get().(*TmallServicecenterFulfiltaskInsuranceActionAPIResponse)
}

// ReleaseTmallServicecenterFulfiltaskInsuranceActionAPIResponse 将 TmallServicecenterFulfiltaskInsuranceActionAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterFulfiltaskInsuranceActionAPIResponse(v *TmallServicecenterFulfiltaskInsuranceActionAPIResponse) {
	v.Reset()
	poolTmallServicecenterFulfiltaskInsuranceActionAPIResponse.Put(v)
}
