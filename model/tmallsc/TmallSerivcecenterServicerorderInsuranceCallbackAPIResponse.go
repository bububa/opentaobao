package tmallsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallserivcecenterservicerorderinsurancecallbackAPIResponse 服务商回传保单信息 API返回值
// tmall.serivcecenter.servicerorder.insurance.callback
//
// 服务商回传保单信息
type TmallserivcecenterservicerorderinsurancecallbackAPIResponse struct {
	model.CommonResponse
	TmallserivcecenterservicerorderinsurancecallbackAPIResponseModel
}

// TmallserivcecenterservicerorderinsurancecallbackAPIResponseModel is 服务商回传保单信息 成功返回结果
type TmallserivcecenterservicerorderinsurancecallbackAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_serivcecenter_servicerorder_insurance_callback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
