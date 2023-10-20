package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterSpserviceorderCreateAPIResponse 服务单创建 API返回值
// tmall.servicecenter.spserviceorder.create
//
// 服务单创建
type TmallServicecenterSpserviceorderCreateAPIResponse struct {
	model.CommonResponse
	TmallServicecenterSpserviceorderCreateAPIResponseModel
}

// TmallServicecenterSpserviceorderCreateAPIResponseModel is 服务单创建 成功返回结果
type TmallServicecenterSpserviceorderCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_spserviceorder_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 工单id列表
	WorkcardIds []int64 `json:"workcard_ids,omitempty" xml:"workcard_ids>int64,omitempty"`
	// 错误信息
	DisplayMsg string `json:"display_msg,omitempty" xml:"display_msg,omitempty"`
	// 服务单id
	ServiceOrderId int64 `json:"service_order_id,omitempty" xml:"service_order_id,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
