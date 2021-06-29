package aesolution

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
fulfill order API请求
aliexpress.solution.order.fulfill

fulfill order for seller
*/
type AliexpressSolutionOrderFulfillRequest struct {
    model.Params
    // Actual logistics service selected by the user (logistics service key: This interface obtains the currently supportable logistics according to all the supportable logistics services listed by api.listLogisticsService. Please visit the forum link http://bbs.seller.aliexpress.com/bbs/read.php?tid=266120&page=1&toread=1#tpc for the detailed list of logistics services supported by the platform.)
    _serviceName   string
    // When serviceName=other, fill in the corresponding tracking website.
    _trackingWebsite   string
    // order ID for delivery by the user
    _outRef   string
    // Status including: all shipments (all), part of the delivery (part)
    _sendType   string
    // Remarks (only in English, and the length is limited to 512 characters)
    _description   string
    // logistics number
    _logisticsNo   string
}

// 初始化AliexpressSolutionOrderFulfillRequest对象
func NewAliexpressSolutionOrderFulfillRequest() *AliexpressSolutionOrderFulfillRequest{
    return &AliexpressSolutionOrderFulfillRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressSolutionOrderFulfillRequest) GetApiMethodName() string {
    return "aliexpress.solution.order.fulfill"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressSolutionOrderFulfillRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ServiceName Setter
// Actual logistics service selected by the user (logistics service key: This interface obtains the currently supportable logistics according to all the supportable logistics services listed by api.listLogisticsService. Please visit the forum link http://bbs.seller.aliexpress.com/bbs/read.php?tid=266120&page=1&toread=1#tpc for the detailed list of logistics services supported by the platform.)
func (r *AliexpressSolutionOrderFulfillRequest) SetServiceName(_serviceName string) error {
    r._serviceName = _serviceName
    r.Set("service_name", _serviceName)
    return nil
}

// ServiceName Getter
func (r AliexpressSolutionOrderFulfillRequest) GetServiceName() string {
    return r._serviceName
}
// TrackingWebsite Setter
// When serviceName=other, fill in the corresponding tracking website.
func (r *AliexpressSolutionOrderFulfillRequest) SetTrackingWebsite(_trackingWebsite string) error {
    r._trackingWebsite = _trackingWebsite
    r.Set("tracking_website", _trackingWebsite)
    return nil
}

// TrackingWebsite Getter
func (r AliexpressSolutionOrderFulfillRequest) GetTrackingWebsite() string {
    return r._trackingWebsite
}
// OutRef Setter
// order ID for delivery by the user
func (r *AliexpressSolutionOrderFulfillRequest) SetOutRef(_outRef string) error {
    r._outRef = _outRef
    r.Set("out_ref", _outRef)
    return nil
}

// OutRef Getter
func (r AliexpressSolutionOrderFulfillRequest) GetOutRef() string {
    return r._outRef
}
// SendType Setter
// Status including: all shipments (all), part of the delivery (part)
func (r *AliexpressSolutionOrderFulfillRequest) SetSendType(_sendType string) error {
    r._sendType = _sendType
    r.Set("send_type", _sendType)
    return nil
}

// SendType Getter
func (r AliexpressSolutionOrderFulfillRequest) GetSendType() string {
    return r._sendType
}
// Description Setter
// Remarks (only in English, and the length is limited to 512 characters)
func (r *AliexpressSolutionOrderFulfillRequest) SetDescription(_description string) error {
    r._description = _description
    r.Set("description", _description)
    return nil
}

// Description Getter
func (r AliexpressSolutionOrderFulfillRequest) GetDescription() string {
    return r._description
}
// LogisticsNo Setter
// logistics number
func (r *AliexpressSolutionOrderFulfillRequest) SetLogisticsNo(_logisticsNo string) error {
    r._logisticsNo = _logisticsNo
    r.Set("logistics_no", _logisticsNo)
    return nil
}

// LogisticsNo Getter
func (r AliexpressSolutionOrderFulfillRequest) GetLogisticsNo() string {
    return r._logisticsNo
}
