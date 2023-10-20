package tmallsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardLogisticsinfoCreateAPIResponse 创建服务履约物流单 API返回值
// tmall.servicecenter.workcard.logisticsinfo.create
//
// 创建服务履约物流单
type TmallServicecenterWorkcardLogisticsinfoCreateAPIResponse struct {
	model.CommonResponse
	TmallServicecenterWorkcardLogisticsinfoCreateAPIResponseModel
}

// TmallServicecenterWorkcardLogisticsinfoCreateAPIResponseModel is 创建服务履约物流单 成功返回结果
type TmallServicecenterWorkcardLogisticsinfoCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_workcard_logisticsinfo_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
