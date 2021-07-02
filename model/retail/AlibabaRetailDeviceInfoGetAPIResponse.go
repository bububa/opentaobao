package retail

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailDeviceInfoGetAPIResponse 贩卖机设备信息获取 API返回值
// alibaba.retail.device.info.get
//
// 贩卖机设备信息获取
type AlibabaRetailDeviceInfoGetAPIResponse struct {
	model.CommonResponse
	AlibabaRetailDeviceInfoGetAPIResponseModel
}

// AlibabaRetailDeviceInfoGetAPIResponseModel is 贩卖机设备信息获取 成功返回结果
type AlibabaRetailDeviceInfoGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_retail_device_info_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	Data string `json:"data,omitempty" xml:"data,omitempty"`
}
