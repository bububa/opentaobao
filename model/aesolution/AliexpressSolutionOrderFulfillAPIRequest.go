package aesolution

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSolutionOrderFulfillAPIRequest fulfill order API请求
// aliexpress.solution.order.fulfill
//
// fulfill order for seller
type AliexpressSolutionOrderFulfillAPIRequest struct {
	model.Params
	// Actual logistics service selected by the user (logistics service key: This interface obtains the currently supportable logistics according to all the supportable logistics services listed by api.listLogisticsService. Please visit the forum link http://bbs.seller.aliexpress.com/bbs/read.php?tid=266120&page=1&toread=1#tpc for the detailed list of logistics services supported by the platform.)
	_serviceName string
	// When serviceName=other, fill in the corresponding tracking website.
	_trackingWebsite string
	// order ID for delivery by the user
	_outRef string
	// Status including: all shipments (all), part of the delivery (part)
	_sendType string
	// Remarks (only in English, and the length is limited to 512 characters)
	_description string
	// logistics number
	_logisticsNo string
}

// NewAliexpressSolutionOrderFulfillRequest 初始化AliexpressSolutionOrderFulfillAPIRequest对象
func NewAliexpressSolutionOrderFulfillRequest() *AliexpressSolutionOrderFulfillAPIRequest {
	return &AliexpressSolutionOrderFulfillAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressSolutionOrderFulfillAPIRequest) GetApiMethodName() string {
	return "aliexpress.solution.order.fulfill"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressSolutionOrderFulfillAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressSolutionOrderFulfillAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetServiceName is ServiceName Setter
// Actual logistics service selected by the user (logistics service key: This interface obtains the currently supportable logistics according to all the supportable logistics services listed by api.listLogisticsService. Please visit the forum link http://bbs.seller.aliexpress.com/bbs/read.php?tid=266120&amp;page=1&amp;toread=1#tpc for the detailed list of logistics services supported by the platform.)
func (r *AliexpressSolutionOrderFulfillAPIRequest) SetServiceName(_serviceName string) error {
	r._serviceName = _serviceName
	r.Set("service_name", _serviceName)
	return nil
}

// GetServiceName ServiceName Getter
func (r AliexpressSolutionOrderFulfillAPIRequest) GetServiceName() string {
	return r._serviceName
}

// SetTrackingWebsite is TrackingWebsite Setter
// When serviceName=other, fill in the corresponding tracking website.
func (r *AliexpressSolutionOrderFulfillAPIRequest) SetTrackingWebsite(_trackingWebsite string) error {
	r._trackingWebsite = _trackingWebsite
	r.Set("tracking_website", _trackingWebsite)
	return nil
}

// GetTrackingWebsite TrackingWebsite Getter
func (r AliexpressSolutionOrderFulfillAPIRequest) GetTrackingWebsite() string {
	return r._trackingWebsite
}

// SetOutRef is OutRef Setter
// order ID for delivery by the user
func (r *AliexpressSolutionOrderFulfillAPIRequest) SetOutRef(_outRef string) error {
	r._outRef = _outRef
	r.Set("out_ref", _outRef)
	return nil
}

// GetOutRef OutRef Getter
func (r AliexpressSolutionOrderFulfillAPIRequest) GetOutRef() string {
	return r._outRef
}

// SetSendType is SendType Setter
// Status including: all shipments (all), part of the delivery (part)
func (r *AliexpressSolutionOrderFulfillAPIRequest) SetSendType(_sendType string) error {
	r._sendType = _sendType
	r.Set("send_type", _sendType)
	return nil
}

// GetSendType SendType Getter
func (r AliexpressSolutionOrderFulfillAPIRequest) GetSendType() string {
	return r._sendType
}

// SetDescription is Description Setter
// Remarks (only in English, and the length is limited to 512 characters)
func (r *AliexpressSolutionOrderFulfillAPIRequest) SetDescription(_description string) error {
	r._description = _description
	r.Set("description", _description)
	return nil
}

// GetDescription Description Getter
func (r AliexpressSolutionOrderFulfillAPIRequest) GetDescription() string {
	return r._description
}

// SetLogisticsNo is LogisticsNo Setter
// logistics number
func (r *AliexpressSolutionOrderFulfillAPIRequest) SetLogisticsNo(_logisticsNo string) error {
	r._logisticsNo = _logisticsNo
	r.Set("logistics_no", _logisticsNo)
	return nil
}

// GetLogisticsNo LogisticsNo Getter
func (r AliexpressSolutionOrderFulfillAPIRequest) GetLogisticsNo() string {
	return r._logisticsNo
}
