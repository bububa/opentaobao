package iot

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoailabaicloudtopdevicegetstatusAPIResponse 获取设备状态 API返回值
// taobao.ailab.aicloud.top.device.getstatus
//
// 获取设备状态
type TaobaoailabaicloudtopdevicegetstatusAPIResponse struct {
	model.CommonResponse
	TaobaoailabaicloudtopdevicegetstatusAPIResponseModel
}

// TaobaoailabaicloudtopdevicegetstatusAPIResponseModel is 获取设备状态 成功返回结果
type TaobaoailabaicloudtopdevicegetstatusAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_device_getstatus_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AiCloudResult `json:"result,omitempty" xml:"result,omitempty"`
}
