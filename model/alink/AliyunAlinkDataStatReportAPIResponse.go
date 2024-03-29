package alink

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliyunAlinkDataStatReportAPIResponse 外部离线统计数据上报 API返回值
// aliyun.alink.data.stat.report
//
// 外部合作厂商上报设备的明细数据，或者离线统计数据。
type AliyunAlinkDataStatReportAPIResponse struct {
	model.CommonResponse
	AliyunAlinkDataStatReportAPIResponseModel
}

// Reset 清空结构体
func (m *AliyunAlinkDataStatReportAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliyunAlinkDataStatReportAPIResponseModel).Reset()
}

// AliyunAlinkDataStatReportAPIResponseModel is 外部离线统计数据上报 成功返回结果
type AliyunAlinkDataStatReportAPIResponseModel struct {
	XMLName xml.Name `xml:"aliyun_alink_data_stat_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误消息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 数据入库状态
	Module bool `json:"module,omitempty" xml:"module,omitempty"`
	// 调用是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AliyunAlinkDataStatReportAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.Status = 0
	m.Module = false
	m.IsSuccess = false
}

var poolAliyunAlinkDataStatReportAPIResponse = sync.Pool{
	New: func() any {
		return new(AliyunAlinkDataStatReportAPIResponse)
	},
}

// GetAliyunAlinkDataStatReportAPIResponse 从 sync.Pool 获取 AliyunAlinkDataStatReportAPIResponse
func GetAliyunAlinkDataStatReportAPIResponse() *AliyunAlinkDataStatReportAPIResponse {
	return poolAliyunAlinkDataStatReportAPIResponse.Get().(*AliyunAlinkDataStatReportAPIResponse)
}

// ReleaseAliyunAlinkDataStatReportAPIResponse 将 AliyunAlinkDataStatReportAPIResponse 保存到 sync.Pool
func ReleaseAliyunAlinkDataStatReportAPIResponse(v *AliyunAlinkDataStatReportAPIResponse) {
	v.Reset()
	poolAliyunAlinkDataStatReportAPIResponse.Put(v)
}
