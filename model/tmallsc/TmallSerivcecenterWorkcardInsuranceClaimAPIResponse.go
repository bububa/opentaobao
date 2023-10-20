package tmallsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallserivcecenterworkcardinsuranceclaimAPIResponse 保险理赔回传工单记录 API返回值
// tmall.serivcecenter.workcard.insurance.claim
//
// 保险理赔回传工单记录
type TmallserivcecenterworkcardinsuranceclaimAPIResponse struct {
	model.CommonResponse
	TmallserivcecenterworkcardinsuranceclaimAPIResponseModel
}

// TmallserivcecenterworkcardinsuranceclaimAPIResponseModel is 保险理赔回传工单记录 成功返回结果
type TmallserivcecenterworkcardinsuranceclaimAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_serivcecenter_workcard_insurance_claim_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
