package iot

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoailabaicloudtopdevicecontrollampAPIResponse 台灯控制 API返回值
// taobao.ailab.aicloud.top.device.control.lamp
//
// 台灯控制
type TaobaoailabaicloudtopdevicecontrollampAPIResponse struct {
	model.CommonResponse
	TaobaoailabaicloudtopdevicecontrollampAPIResponseModel
}

// TaobaoailabaicloudtopdevicecontrollampAPIResponseModel is 台灯控制 成功返回结果
type TaobaoailabaicloudtopdevicecontrollampAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_device_control_lamp_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AiCloudResult `json:"result,omitempty" xml:"result,omitempty"`
}
