package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusDevicehubOpenapiReportdataAPIResponse 设备数据上报 API返回值
// alibaba.campus.devicehub.openapi.reportdata
//
// 设备数据上报
type AlibabaCampusDevicehubOpenapiReportdataAPIResponse struct {
	model.CommonResponse
	AlibabaCampusDevicehubOpenapiReportdataAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusDevicehubOpenapiReportdataAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusDevicehubOpenapiReportdataAPIResponseModel).Reset()
}

// AlibabaCampusDevicehubOpenapiReportdataAPIResponseModel is 设备数据上报 成功返回结果
type AlibabaCampusDevicehubOpenapiReportdataAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_devicehub_openapi_reportdata_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 自动生成
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusDevicehubOpenapiReportdataAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusDevicehubOpenapiReportdataAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusDevicehubOpenapiReportdataAPIResponse)
	},
}

// GetAlibabaCampusDevicehubOpenapiReportdataAPIResponse 从 sync.Pool 获取 AlibabaCampusDevicehubOpenapiReportdataAPIResponse
func GetAlibabaCampusDevicehubOpenapiReportdataAPIResponse() *AlibabaCampusDevicehubOpenapiReportdataAPIResponse {
	return poolAlibabaCampusDevicehubOpenapiReportdataAPIResponse.Get().(*AlibabaCampusDevicehubOpenapiReportdataAPIResponse)
}

// ReleaseAlibabaCampusDevicehubOpenapiReportdataAPIResponse 将 AlibabaCampusDevicehubOpenapiReportdataAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusDevicehubOpenapiReportdataAPIResponse(v *AlibabaCampusDevicehubOpenapiReportdataAPIResponse) {
	v.Reset()
	poolAlibabaCampusDevicehubOpenapiReportdataAPIResponse.Put(v)
}
