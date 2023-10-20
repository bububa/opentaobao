package idle

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidlerentordercheckstatusuploadAPIResponse 上传验收结果 API返回值
// alibaba.idle.rent.order.checkstatus.upload
//
// 上传验收结果
type AlibabaidlerentordercheckstatusuploadAPIResponse struct {
	model.CommonResponse
	AlibabaidlerentordercheckstatusuploadAPIResponseModel
}

// AlibabaidlerentordercheckstatusuploadAPIResponseModel is 上传验收结果 成功返回结果
type AlibabaidlerentordercheckstatusuploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_rent_order_checkstatus_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Result *AlibabaidlerentordercheckstatusuploadTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
