package examination

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthexaminationreservestateAPIResponse 体检机构对接_体检状态查询 API返回值
// alibaba.alihealth.examination.reserve.state
//
// 体检机构对接_体检状态查询
type AlibabaalihealthexaminationreservestateAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthexaminationreservestateAPIResponseModel
}

// AlibabaalihealthexaminationreservestateAPIResponseModel is 体检机构对接_体检状态查询 成功返回结果
type AlibabaalihealthexaminationreservestateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_examination_reserve_state_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回结果编码
	ResponseCode string `json:"response_code,omitempty" xml:"response_code,omitempty"`
	// 预约信息
	CooperationOrderInfo *CooperationOrderInfo `json:"cooperation_order_info,omitempty" xml:"cooperation_order_info,omitempty"`
}
