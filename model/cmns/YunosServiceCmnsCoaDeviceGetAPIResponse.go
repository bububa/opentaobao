package cmns

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* YunosServiceCmnsCoaDeviceGetAPIResponse
设备详情查询 API返回值
yunos.service.cmns.coa.device.get

第三方应用开发者调用此接口查询设备详情 */
type YunosServiceCmnsCoaDeviceGetAPIResponse struct {
	model.CommonResponse
	YunosServiceCmnsCoaDeviceGetAPIResponseModel
}

// YunosServiceCmnsCoaDeviceGetAPIResponseModel is 设备详情查询 成功返回结果
type YunosServiceCmnsCoaDeviceGetAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_service_cmns_coa_device_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 设备详情
	DeviceList []DeviceResult `json:"device_list,omitempty" xml:"device_list>device_result,omitempty"`
	// 接口查询出错提示信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 200表示查询成功
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}
