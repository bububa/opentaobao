package alink

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliyunAlinkDataStatReportAPIRequest 外部离线统计数据上报 API请求
// aliyun.alink.data.stat.report
//
// 外部合作厂商上报设备的明细数据，或者离线统计数据。
type AliyunAlinkDataStatReportAPIRequest struct {
	model.Params
	// 入参对象
	_paramBean *OuterDataBean
}

// NewAliyunAlinkDataStatReportRequest 初始化AliyunAlinkDataStatReportAPIRequest对象
func NewAliyunAlinkDataStatReportRequest() *AliyunAlinkDataStatReportAPIRequest {
	return &AliyunAlinkDataStatReportAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliyunAlinkDataStatReportAPIRequest) Reset() {
	r._paramBean = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliyunAlinkDataStatReportAPIRequest) GetApiMethodName() string {
	return "aliyun.alink.data.stat.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliyunAlinkDataStatReportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliyunAlinkDataStatReportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBean is ParamBean Setter
// 入参对象
func (r *AliyunAlinkDataStatReportAPIRequest) SetParamBean(_paramBean *OuterDataBean) error {
	r._paramBean = _paramBean
	r.Set("param_bean", _paramBean)
	return nil
}

// GetParamBean ParamBean Getter
func (r AliyunAlinkDataStatReportAPIRequest) GetParamBean() *OuterDataBean {
	return r._paramBean
}

var poolAliyunAlinkDataStatReportAPIRequest = sync.Pool{
	New: func() any {
		return NewAliyunAlinkDataStatReportRequest()
	},
}

// GetAliyunAlinkDataStatReportRequest 从 sync.Pool 获取 AliyunAlinkDataStatReportAPIRequest
func GetAliyunAlinkDataStatReportAPIRequest() *AliyunAlinkDataStatReportAPIRequest {
	return poolAliyunAlinkDataStatReportAPIRequest.Get().(*AliyunAlinkDataStatReportAPIRequest)
}

// ReleaseAliyunAlinkDataStatReportAPIRequest 将 AliyunAlinkDataStatReportAPIRequest 放入 sync.Pool
func ReleaseAliyunAlinkDataStatReportAPIRequest(v *AliyunAlinkDataStatReportAPIRequest) {
	v.Reset()
	poolAliyunAlinkDataStatReportAPIRequest.Put(v)
}
