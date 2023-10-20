package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascplogisticscpgetAPIRequest 快递公司资源列表查询接口 API请求
// alibaba.ascp.logistics.cp.get
//
// 快递公司资源列表查询接口
type AlibabaascplogisticscpgetAPIRequest struct {
	model.Params
	// 请求体
	_logisticsResourceRequest *LogisticsResourceRequest
}

// NewAlibabaascplogisticscpgetRequest 初始化AlibabaascplogisticscpgetAPIRequest对象
func NewAlibabaascplogisticscpgetRequest() *AlibabaascplogisticscpgetAPIRequest {
	return &AlibabaascplogisticscpgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascplogisticscpgetAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.logistics.cp.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascplogisticscpgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascplogisticscpgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLogisticsResourceRequest is LogisticsResourceRequest Setter
// 请求体
func (r *AlibabaascplogisticscpgetAPIRequest) SetLogisticsResourceRequest(_logisticsResourceRequest *LogisticsResourceRequest) error {
	r._logisticsResourceRequest = _logisticsResourceRequest
	r.Set("logistics_resource_request", _logisticsResourceRequest)
	return nil
}

// GetLogisticsResourceRequest LogisticsResourceRequest Getter
func (r AlibabaascplogisticscpgetAPIRequest) GetLogisticsResourceRequest() *LogisticsResourceRequest {
	return r._logisticsResourceRequest
}
