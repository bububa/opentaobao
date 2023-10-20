package alihealthmedical

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthmedicalitempublishAPIResponse 三方入驻-开通服务 API返回值
// alibaba.alihealth.medical.item.publish
//
// 三方入驻-开通服务
type AlibabaalihealthmedicalitempublishAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthmedicalitempublishAPIResponseModel
}

// AlibabaalihealthmedicalitempublishAPIResponseModel is 三方入驻-开通服务 成功返回结果
type AlibabaalihealthmedicalitempublishAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_medical_item_publish_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统返回的通用结果类
	Result1 *ServiceResult `json:"result1,omitempty" xml:"result1,omitempty"`
}
