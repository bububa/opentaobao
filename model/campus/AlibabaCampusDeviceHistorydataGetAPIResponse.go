package campus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusDeviceHistorydataGetAPIResponse 设备历史数据批量获取 API返回值
// alibaba.campus.device.historydata.get
//
// 设备历史数据批量获取
type AlibabaCampusDeviceHistorydataGetAPIResponse struct {
	model.CommonResponse
	AlibabaCampusDeviceHistorydataGetAPIResponseModel
}

// AlibabaCampusDeviceHistorydataGetAPIResponseModel is 设备历史数据批量获取 成功返回结果
type AlibabaCampusDeviceHistorydataGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_device_historydata_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ListResult `json:"result,omitempty" xml:"result,omitempty"`
}
