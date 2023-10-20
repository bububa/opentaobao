package servicecenter

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallCarFpcarGetcarNotifyAPIResponse 门店通知用户提车 API返回值
// tmall.car.fpcar.getcar.notify
//
// 提供给外部(大搜或其它合作方)的接口-门店通知用户提车
type TmallCarFpcarGetcarNotifyAPIResponse struct {
	model.CommonResponse
	TmallCarFpcarGetcarNotifyAPIResponseModel
}

// TmallCarFpcarGetcarNotifyAPIResponseModel is 门店通知用户提车 成功返回结果
type TmallCarFpcarGetcarNotifyAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_car_fpcar_getcar_notify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回的数据结果
	Object string `json:"object,omitempty" xml:"object,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 是否成功
	Succes bool `json:"succes,omitempty" xml:"succes,omitempty"`
}
