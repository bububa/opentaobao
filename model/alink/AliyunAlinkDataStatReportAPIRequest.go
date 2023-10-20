package alink

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliyunalinkdatastatreportAPIRequest 外部离线统计数据上报 API请求
// aliyun.alink.data.stat.report
//
// 外部合作厂商上报设备的明细数据，或者离线统计数据。
type AliyunalinkdatastatreportAPIRequest struct {
	model.Params
	// 入参对象
	_paramBean *OuterDataBean
}

// NewAliyunalinkdatastatreportRequest 初始化AliyunalinkdatastatreportAPIRequest对象
func NewAliyunalinkdatastatreportRequest() *AliyunalinkdatastatreportAPIRequest {
	return &AliyunalinkdatastatreportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliyunalinkdatastatreportAPIRequest) GetApiMethodName() string {
	return "aliyun.alink.data.stat.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliyunalinkdatastatreportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliyunalinkdatastatreportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBean is ParamBean Setter
// 入参对象
func (r *AliyunalinkdatastatreportAPIRequest) SetParamBean(_paramBean *OuterDataBean) error {
	r._paramBean = _paramBean
	r.Set("param_bean", _paramBean)
	return nil
}

// GetParamBean ParamBean Getter
func (r AliyunalinkdatastatreportAPIRequest) GetParamBean() *OuterDataBean {
	return r._paramBean
}
