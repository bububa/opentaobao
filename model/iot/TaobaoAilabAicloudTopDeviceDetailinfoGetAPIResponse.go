package iot

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopDeviceDetailinfoGetAPIResponse 获取设备详细信息 API返回值
// taobao.ailab.aicloud.top.device.detailinfo.get
//
// 获取设备详细信息
type TaobaoAilabAicloudTopDeviceDetailinfoGetAPIResponse struct {
	model.CommonResponse
	TaobaoAilabAicloudTopDeviceDetailinfoGetAPIResponseModel
}

// TaobaoAilabAicloudTopDeviceDetailinfoGetAPIResponseModel is 获取设备详细信息 成功返回结果
type TaobaoAilabAicloudTopDeviceDetailinfoGetAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_device_detailinfo_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoAilabAicloudTopDeviceDetailinfoGetResult `json:"result,omitempty" xml:"result,omitempty"`
}
