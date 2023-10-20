package iot

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoailabaicloudtopdevicestatusinfogetAPIResponse 获取设备状态信息 API返回值
// taobao.ailab.aicloud.top.device.statusinfo.get
//
// 获取设备状态信息
type TaobaoailabaicloudtopdevicestatusinfogetAPIResponse struct {
	model.CommonResponse
	TaobaoailabaicloudtopdevicestatusinfogetAPIResponseModel
}

// TaobaoailabaicloudtopdevicestatusinfogetAPIResponseModel is 获取设备状态信息 成功返回结果
type TaobaoailabaicloudtopdevicestatusinfogetAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_device_statusinfo_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoailabaicloudtopdevicestatusinfogetResult `json:"result,omitempty" xml:"result,omitempty"`
}
