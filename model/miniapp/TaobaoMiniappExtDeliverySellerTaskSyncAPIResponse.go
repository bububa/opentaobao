package miniapp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaominiappextdeliverysellertasksyncAPIResponse 同步外投任务 API返回值
// taobao.miniapp.ext.delivery.seller.task.sync
//
// 同步外投任务
type TaobaominiappextdeliverysellertasksyncAPIResponse struct {
	model.CommonResponse
	TaobaominiappextdeliverysellertasksyncAPIResponseModel
}

// TaobaominiappextdeliverysellertasksyncAPIResponseModel is 同步外投任务 成功返回结果
type TaobaominiappextdeliverysellertasksyncAPIResponseModel struct {
	XMLName xml.Name `xml:"miniapp_ext_delivery_seller_task_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorType int64 `json:"error_type,omitempty" xml:"error_type,omitempty"`
	// 任务id
	Model int64 `json:"model,omitempty" xml:"model,omitempty"`
	// true or false
	Successful bool `json:"successful,omitempty" xml:"successful,omitempty"`
}
