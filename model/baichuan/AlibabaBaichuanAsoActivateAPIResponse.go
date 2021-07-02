package baichuan

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaBaichuanAsoActivateAPIResponse 设备安装活动激活 API返回值
// alibaba.baichuan.aso.activate
//
// 设备安装活动激活
type AlibabaBaichuanAsoActivateAPIResponse struct {
	model.CommonResponse
	AlibabaBaichuanAsoActivateAPIResponseModel
}

// AlibabaBaichuanAsoActivateAPIResponseModel is 设备安装活动激活 成功返回结果
type AlibabaBaichuanAsoActivateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_baichuan_aso_activate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AsoActivateDeviceResult `json:"result,omitempty" xml:"result,omitempty"`
}
