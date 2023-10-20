package tmallgenie

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsIotCloudDeviceReportAPIRequest 天猫精灵云云接入设备状态、事件上报接口 API请求
// alibaba.ailabs.iot.cloud.device.report
//
// 承接天猫精灵云云接入设备的状态、事件上报
type AlibabaAilabsIotCloudDeviceReportAPIRequest struct {
	model.Params
	// 上报总入参
	_cloudReportParam *CloudReportParam
}

// NewAlibabaAilabsIotCloudDeviceReportRequest 初始化AlibabaAilabsIotCloudDeviceReportAPIRequest对象
func NewAlibabaAilabsIotCloudDeviceReportRequest() *AlibabaAilabsIotCloudDeviceReportAPIRequest {
	return &AlibabaAilabsIotCloudDeviceReportAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAilabsIotCloudDeviceReportAPIRequest) Reset() {
	r._cloudReportParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabsIotCloudDeviceReportAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.iot.cloud.device.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabsIotCloudDeviceReportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAilabsIotCloudDeviceReportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCloudReportParam is CloudReportParam Setter
// 上报总入参
func (r *AlibabaAilabsIotCloudDeviceReportAPIRequest) SetCloudReportParam(_cloudReportParam *CloudReportParam) error {
	r._cloudReportParam = _cloudReportParam
	r.Set("cloud_report_param", _cloudReportParam)
	return nil
}

// GetCloudReportParam CloudReportParam Getter
func (r AlibabaAilabsIotCloudDeviceReportAPIRequest) GetCloudReportParam() *CloudReportParam {
	return r._cloudReportParam
}

var poolAlibabaAilabsIotCloudDeviceReportAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAilabsIotCloudDeviceReportRequest()
	},
}

// GetAlibabaAilabsIotCloudDeviceReportRequest 从 sync.Pool 获取 AlibabaAilabsIotCloudDeviceReportAPIRequest
func GetAlibabaAilabsIotCloudDeviceReportAPIRequest() *AlibabaAilabsIotCloudDeviceReportAPIRequest {
	return poolAlibabaAilabsIotCloudDeviceReportAPIRequest.Get().(*AlibabaAilabsIotCloudDeviceReportAPIRequest)
}

// ReleaseAlibabaAilabsIotCloudDeviceReportAPIRequest 将 AlibabaAilabsIotCloudDeviceReportAPIRequest 放入 sync.Pool
func ReleaseAlibabaAilabsIotCloudDeviceReportAPIRequest(v *AlibabaAilabsIotCloudDeviceReportAPIRequest) {
	v.Reset()
	poolAlibabaAilabsIotCloudDeviceReportAPIRequest.Put(v)
}
