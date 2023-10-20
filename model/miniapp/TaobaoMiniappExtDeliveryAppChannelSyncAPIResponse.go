package miniapp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaominiappextdeliveryappchannelsyncAPIResponse ISV写入应用的渠道信息 API返回值
// taobao.miniapp.ext.delivery.app.channel.sync
//
// ISV写入应用的渠道信息
type TaobaominiappextdeliveryappchannelsyncAPIResponse struct {
	model.CommonResponse
	TaobaominiappextdeliveryappchannelsyncAPIResponseModel
}

// TaobaominiappextdeliveryappchannelsyncAPIResponseModel is ISV写入应用的渠道信息 成功返回结果
type TaobaominiappextdeliveryappchannelsyncAPIResponseModel struct {
	XMLName xml.Name `xml:"miniapp_ext_delivery_app_channel_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorType int64 `json:"error_type,omitempty" xml:"error_type,omitempty"`
	// configId
	Model int64 `json:"model,omitempty" xml:"model,omitempty"`
	// true or false
	Successful bool `json:"successful,omitempty" xml:"successful,omitempty"`
}
