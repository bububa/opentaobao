package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidletenderuploadreportAPIRequest 服务商上传验货报告同步给闲鱼 API请求
// alibaba.idle.tender.upload.report
//
// 服务商上传验货报告同步给闲鱼
type AlibabaidletenderuploadreportAPIRequest struct {
	model.Params
	// 回收商质检报告
	_param0 *InspectionReport
}

// NewAlibabaidletenderuploadreportRequest 初始化AlibabaidletenderuploadreportAPIRequest对象
func NewAlibabaidletenderuploadreportRequest() *AlibabaidletenderuploadreportAPIRequest {
	return &AlibabaidletenderuploadreportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidletenderuploadreportAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.tender.upload.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidletenderuploadreportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidletenderuploadreportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 回收商质检报告
func (r *AlibabaidletenderuploadreportAPIRequest) SetParam0(_param0 *InspectionReport) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaidletenderuploadreportAPIRequest) GetParam0() *InspectionReport {
	return r._param0
}
