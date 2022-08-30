package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressLocalLogisticsReportShippedAPIRequest report as shipped API请求
// aliexpress.local.logistics.report.shipped
//
// report as shipped
type AliexpressLocalLogisticsReportShippedAPIRequest struct {
	model.Params
	// param
	_param1 *ReportShippedRequestDto
}

// NewAliexpressLocalLogisticsReportShippedRequest 初始化AliexpressLocalLogisticsReportShippedAPIRequest对象
func NewAliexpressLocalLogisticsReportShippedRequest() *AliexpressLocalLogisticsReportShippedAPIRequest {
	return &AliexpressLocalLogisticsReportShippedAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressLocalLogisticsReportShippedAPIRequest) GetApiMethodName() string {
	return "aliexpress.local.logistics.report.shipped"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressLocalLogisticsReportShippedAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam1 is Param1 Setter
// param
func (r *AliexpressLocalLogisticsReportShippedAPIRequest) SetParam1(_param1 *ReportShippedRequestDto) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AliexpressLocalLogisticsReportShippedAPIRequest) GetParam1() *ReportShippedRequestDto {
	return r._param1
}
