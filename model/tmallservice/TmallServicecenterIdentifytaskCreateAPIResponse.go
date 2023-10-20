package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenteridentifytaskcreateAPIResponse 服务商创建核销单 API返回值
// tmall.servicecenter.identifytask.create
//
// 服务商调用该接口进行创建核销单操作
type TmallservicecenteridentifytaskcreateAPIResponse struct {
	model.CommonResponse
	TmallservicecenteridentifytaskcreateAPIResponseModel
}

// TmallservicecenteridentifytaskcreateAPIResponseModel is 服务商创建核销单 成功返回结果
type TmallservicecenteridentifytaskcreateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_identifytask_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// -
	Result *FulfilplatformResult `json:"result,omitempty" xml:"result,omitempty"`
}
