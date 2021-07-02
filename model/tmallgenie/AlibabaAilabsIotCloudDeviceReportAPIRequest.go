package tmallgenie

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabsIotCloudDeviceReportAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.iot.cloud.device.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabsIotCloudDeviceReportAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CloudReportParam Setter
// 上报总入参
func (r *AlibabaAilabsIotCloudDeviceReportAPIRequest) SetCloudReportParam(_cloudReportParam *CloudReportParam) error {
	r._cloudReportParam = _cloudReportParam
	r.Set("cloud_report_param", _cloudReportParam)
	return nil
}

// Get CloudReportParam Getter
func (r AlibabaAilabsIotCloudDeviceReportAPIRequest) GetCloudReportParam() *CloudReportParam {
	return r._cloudReportParam
}
