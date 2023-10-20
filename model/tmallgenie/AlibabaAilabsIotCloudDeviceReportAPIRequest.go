package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaailabsiotclouddevicereportAPIRequest 天猫精灵云云接入设备状态、事件上报接口 API请求
// alibaba.ailabs.iot.cloud.device.report
//
// 承接天猫精灵云云接入设备的状态、事件上报
type AlibabaailabsiotclouddevicereportAPIRequest struct {
	model.Params
	// 上报总入参
	_cloudReportParam *CloudReportParam
}

// NewAlibabaailabsiotclouddevicereportRequest 初始化AlibabaailabsiotclouddevicereportAPIRequest对象
func NewAlibabaailabsiotclouddevicereportRequest() *AlibabaailabsiotclouddevicereportAPIRequest {
	return &AlibabaailabsiotclouddevicereportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaailabsiotclouddevicereportAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.iot.cloud.device.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaailabsiotclouddevicereportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaailabsiotclouddevicereportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCloudReportParam is CloudReportParam Setter
// 上报总入参
func (r *AlibabaailabsiotclouddevicereportAPIRequest) SetCloudReportParam(_cloudReportParam *CloudReportParam) error {
	r._cloudReportParam = _cloudReportParam
	r.Set("cloud_report_param", _cloudReportParam)
	return nil
}

// GetCloudReportParam CloudReportParam Getter
func (r AlibabaailabsiotclouddevicereportAPIRequest) GetCloudReportParam() *CloudReportParam {
	return r._cloudReportParam
}
