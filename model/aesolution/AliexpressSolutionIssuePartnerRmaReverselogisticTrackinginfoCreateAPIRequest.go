package aesolution

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpresssolutionissuepartnerrmareverselogistictrackinginfocreateAPIRequest aliexpress.solution.issue.partner.rma.reverselogistic.trackinginfo.create API请求
// aliexpress.solution.issue.partner.rma.reverselogistic.trackinginfo.create
//
// Receives information about reverse logistics tracking info
type AliexpresssolutionissuepartnerrmareverselogistictrackinginfocreateAPIRequest struct {
	model.Params
	// Logistic's order creation request
	_logisticsOrderCreationRequest *LogisticOrderCreationForRmaRequest
}

// NewAliexpresssolutionissuepartnerrmareverselogistictrackinginfocreateRequest 初始化AliexpresssolutionissuepartnerrmareverselogistictrackinginfocreateAPIRequest对象
func NewAliexpresssolutionissuepartnerrmareverselogistictrackinginfocreateRequest() *AliexpresssolutionissuepartnerrmareverselogistictrackinginfocreateAPIRequest {
	return &AliexpresssolutionissuepartnerrmareverselogistictrackinginfocreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpresssolutionissuepartnerrmareverselogistictrackinginfocreateAPIRequest) GetApiMethodName() string {
	return "aliexpress.solution.issue.partner.rma.reverselogistic.trackinginfo.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpresssolutionissuepartnerrmareverselogistictrackinginfocreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpresssolutionissuepartnerrmareverselogistictrackinginfocreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLogisticsOrderCreationRequest is LogisticsOrderCreationRequest Setter
// Logistic&#39;s order creation request
func (r *AliexpresssolutionissuepartnerrmareverselogistictrackinginfocreateAPIRequest) SetLogisticsOrderCreationRequest(_logisticsOrderCreationRequest *LogisticOrderCreationForRmaRequest) error {
	r._logisticsOrderCreationRequest = _logisticsOrderCreationRequest
	r.Set("logistics_order_creation_request", _logisticsOrderCreationRequest)
	return nil
}

// GetLogisticsOrderCreationRequest LogisticsOrderCreationRequest Getter
func (r AliexpresssolutionissuepartnerrmareverselogistictrackinginfocreateAPIRequest) GetLogisticsOrderCreationRequest() *LogisticOrderCreationForRmaRequest {
	return r._logisticsOrderCreationRequest
}
