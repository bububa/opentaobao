package alink

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliyunAlinkDataStatReportAPIRequest
外部离线统计数据上报 API请求
aliyun.alink.data.stat.report

外部合作厂商上报设备的明细数据，或者离线统计数据。 */
type AliyunAlinkDataStatReportAPIRequest struct {
	model.Params
	// 入参对象
	_paramBean *OuterDataBean
}

// NewAliyunAlinkDataStatReportRequest 初始化AliyunAlinkDataStatReportAPIRequest对象
func NewAliyunAlinkDataStatReportRequest() *AliyunAlinkDataStatReportAPIRequest {
	return &AliyunAlinkDataStatReportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliyunAlinkDataStatReportAPIRequest) GetApiMethodName() string {
	return "aliyun.alink.data.stat.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliyunAlinkDataStatReportAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamBean Setter
// 入参对象
func (r *AliyunAlinkDataStatReportAPIRequest) SetParamBean(_paramBean *OuterDataBean) error {
	r._paramBean = _paramBean
	r.Set("param_bean", _paramBean)
	return nil
}

// Get ParamBean Getter
func (r AliyunAlinkDataStatReportAPIRequest) GetParamBean() *OuterDataBean {
	return r._paramBean
}
