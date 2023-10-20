package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpresslocallogisticsreportshippedAPIRequest report as shipped API请求
// aliexpress.local.logistics.report.shipped
//
// report as shipped
type AliexpresslocallogisticsreportshippedAPIRequest struct {
	model.Params
	// param
	_param1 *ReportShippedRequestDto
}

// NewAliexpresslocallogisticsreportshippedRequest 初始化AliexpresslocallogisticsreportshippedAPIRequest对象
func NewAliexpresslocallogisticsreportshippedRequest() *AliexpresslocallogisticsreportshippedAPIRequest {
	return &AliexpresslocallogisticsreportshippedAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpresslocallogisticsreportshippedAPIRequest) GetApiMethodName() string {
	return "aliexpress.local.logistics.report.shipped"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpresslocallogisticsreportshippedAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpresslocallogisticsreportshippedAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam1 is Param1 Setter
// param
func (r *AliexpresslocallogisticsreportshippedAPIRequest) SetParam1(_param1 *ReportShippedRequestDto) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AliexpresslocallogisticsreportshippedAPIRequest) GetParam1() *ReportShippedRequestDto {
	return r._param1
}
