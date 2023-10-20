package alink

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalinkdeviceunifystatussetAPIResponse 设置设备标准属性状态 API返回值
// alibaba.alink.device.unify.status.set
//
// 操作用户绑定的设备
type AlibabaalinkdeviceunifystatussetAPIResponse struct {
	model.CommonResponse
	AlibabaalinkdeviceunifystatussetAPIResponseModel
}

// AlibabaalinkdeviceunifystatussetAPIResponseModel is 设置设备标准属性状态 成功返回结果
type AlibabaalinkdeviceunifystatussetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alink_device_unify_status_set_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *TopServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}
