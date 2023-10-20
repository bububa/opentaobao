package tmallsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardLogisticsinfoUpdateAPIResponse 物流单信息回传接口 API返回值
// tmall.servicecenter.workcard.logisticsinfo.update
//
// 物流单信息回传接口
type TmallServicecenterWorkcardLogisticsinfoUpdateAPIResponse struct {
	model.CommonResponse
	TmallServicecenterWorkcardLogisticsinfoUpdateAPIResponseModel
}

// TmallServicecenterWorkcardLogisticsinfoUpdateAPIResponseModel is 物流单信息回传接口 成功返回结果
type TmallServicecenterWorkcardLogisticsinfoUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_workcard_logisticsinfo_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
