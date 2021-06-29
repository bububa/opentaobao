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
    serviceName   string
    // When serviceName=other, fill in the corresponding tracking website.
    trackingWebsite   string
    // order ID for delivery by the user
    outRef   string
    // Status including: all shipments (all), part of the delivery (part)
    sendType   string
    // Remarks (only in English, and the length is limited to 512 characters)
    description   string
    // logistics number
    logisticsNo   string
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
func (r *AliexpressSolutionOrderFulfillRequest) SetServiceName(serviceName string) error {
    r.serviceName = serviceName
    r.Set("service_name", serviceName)
    return nil
}

// ServiceName Getter
func (r AliexpressSolutionOrderFulfillRequest) GetServiceName() string {
    return r.serviceName
}
// TrackingWebsite Setter
// When serviceName=other, fill in the corresponding tracking website.
func (r *AliexpressSolutionOrderFulfillRequest) SetTrackingWebsite(trackingWebsite string) error {
    r.trackingWebsite = trackingWebsite
    r.Set("tracking_website", trackingWebsite)
    return nil
}

// TrackingWebsite Getter
func (r AliexpressSolutionOrderFulfillRequest) GetTrackingWebsite() string {
    return r.trackingWebsite
}
// OutRef Setter
// order ID for delivery by the user
func (r *AliexpressSolutionOrderFulfillRequest) SetOutRef(outRef string) error {
    r.outRef = outRef
    r.Set("out_ref", outRef)
    return nil
}

// OutRef Getter
func (r AliexpressSolutionOrderFulfillRequest) GetOutRef() string {
    return r.outRef
}
// SendType Setter
// Status including: all shipments (all), part of the delivery (part)
func (r *AliexpressSolutionOrderFulfillRequest) SetSendType(sendType string) error {
    r.sendType = sendType
    r.Set("send_type", sendType)
    return nil
}

// SendType Getter
func (r AliexpressSolutionOrderFulfillRequest) GetSendType() string {
    return r.sendType
}
// Description Setter
// Remarks (only in English, and the length is limited to 512 characters)
func (r *AliexpressSolutionOrderFulfillRequest) SetDescription(description string) error {
    r.description = description
    r.Set("description", description)
    return nil
}

// Description Getter
func (r AliexpressSolutionOrderFulfillRequest) GetDescription() string {
    return r.description
}
// LogisticsNo Setter
// logistics number
func (r *AliexpressSolutionOrderFulfillRequest) SetLogisticsNo(logisticsNo string) error {
    r.logisticsNo = logisticsNo
    r.Set("logistics_no", logisticsNo)
    return nil
}

// LogisticsNo Getter
func (r AliexpressSolutionOrderFulfillRequest) GetLogisticsNo() string {
    return r.logisticsNo
}
