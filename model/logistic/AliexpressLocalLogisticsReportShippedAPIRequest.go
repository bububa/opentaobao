package logistic

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliexpressLocalLogisticsReportShippedAPIRequest) Reset() {
	r._param1 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressLocalLogisticsReportShippedAPIRequest) GetApiMethodName() string {
	return "aliexpress.local.logistics.report.shipped"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressLocalLogisticsReportShippedAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressLocalLogisticsReportShippedAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAliexpressLocalLogisticsReportShippedAPIRequest = sync.Pool{
	New: func() any {
		return NewAliexpressLocalLogisticsReportShippedRequest()
	},
}

// GetAliexpressLocalLogisticsReportShippedRequest 从 sync.Pool 获取 AliexpressLocalLogisticsReportShippedAPIRequest
func GetAliexpressLocalLogisticsReportShippedAPIRequest() *AliexpressLocalLogisticsReportShippedAPIRequest {
	return poolAliexpressLocalLogisticsReportShippedAPIRequest.Get().(*AliexpressLocalLogisticsReportShippedAPIRequest)
}

// ReleaseAliexpressLocalLogisticsReportShippedAPIRequest 将 AliexpressLocalLogisticsReportShippedAPIRequest 放入 sync.Pool
func ReleaseAliexpressLocalLogisticsReportShippedAPIRequest(v *AliexpressLocalLogisticsReportShippedAPIRequest) {
	v.Reset()
	poolAliexpressLocalLogisticsReportShippedAPIRequest.Put(v)
}
