package examination

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthexaminationreservestatenotifyAPIResponse 体检机构对接_体检状态主动通知 API返回值
// alibaba.alihealth.examination.reserve.state.notify
//
// 到了体检当天后，服务商主动通知体检预约状态
type AlibabaalihealthexaminationreservestatenotifyAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthexaminationreservestatenotifyAPIResponseModel
}

// AlibabaalihealthexaminationreservestatenotifyAPIResponseModel is 体检机构对接_体检状态主动通知 成功返回结果
type AlibabaalihealthexaminationreservestatenotifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_examination_reserve_state_notify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}
