package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaServicecenterFulfiltaskCreateAPIResponse 合单生成核销单 API返回值
// alibaba.servicecenter.fulfiltask.create
//
// 服务对工单进行合单，合单的结果是生成核销单
type AlibabaServicecenterFulfiltaskCreateAPIResponse struct {
	model.CommonResponse
	AlibabaServicecenterFulfiltaskCreateAPIResponseModel
}

// AlibabaServicecenterFulfiltaskCreateAPIResponseModel is 合单生成核销单 成功返回结果
type AlibabaServicecenterFulfiltaskCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_servicecenter_fulfiltask_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaServicecenterFulfiltaskCreateResult `json:"result,omitempty" xml:"result,omitempty"`
}
