package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaservicecenteridentifytaskcreateAPIResponse 创建核销单 API返回值
// alibaba.servicecenter.identifytask.create
//
// 创建核销单
type AlibabaservicecenteridentifytaskcreateAPIResponse struct {
	model.CommonResponse
	AlibabaservicecenteridentifytaskcreateAPIResponseModel
}

// AlibabaservicecenteridentifytaskcreateAPIResponseModel is 创建核销单 成功返回结果
type AlibabaservicecenteridentifytaskcreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_servicecenter_identifytask_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求结果
	Result *FulfilplatformResult `json:"result,omitempty" xml:"result,omitempty"`
}
