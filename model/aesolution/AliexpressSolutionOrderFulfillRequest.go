package aesolution

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
fulfill order APIRequest
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

func NewAliexpressSolutionOrderFulfillRequest() *AliexpressSolutionOrderFulfillRequest{
    return &AliexpressSolutionOrderFulfillRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressSolutionOrderFulfillRequest) GetApiMethodName() string {
    return "aliexpress.solution.order.fulfill"
}

func (r AliexpressSolutionOrderFulfillRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressSolutionOrderFulfillRequest) SetServiceName(serviceName string) error {
    r.serviceName = serviceName
    r.Set("service_name", serviceName)
    return nil
}

func (r AliexpressSolutionOrderFulfillRequest) GetServiceName() string {
    return r.serviceName
}

func (r *AliexpressSolutionOrderFulfillRequest) SetTrackingWebsite(trackingWebsite string) error {
    r.trackingWebsite = trackingWebsite
    r.Set("tracking_website", trackingWebsite)
    return nil
}

func (r AliexpressSolutionOrderFulfillRequest) GetTrackingWebsite() string {
    return r.trackingWebsite
}

func (r *AliexpressSolutionOrderFulfillRequest) SetOutRef(outRef string) error {
    r.outRef = outRef
    r.Set("out_ref", outRef)
    return nil
}

func (r AliexpressSolutionOrderFulfillRequest) GetOutRef() string {
    return r.outRef
}

func (r *AliexpressSolutionOrderFulfillRequest) SetSendType(sendType string) error {
    r.sendType = sendType
    r.Set("send_type", sendType)
    return nil
}

func (r AliexpressSolutionOrderFulfillRequest) GetSendType() string {
    return r.sendType
}

func (r *AliexpressSolutionOrderFulfillRequest) SetDescription(description string) error {
    r.description = description
    r.Set("description", description)
    return nil
}

func (r AliexpressSolutionOrderFulfillRequest) GetDescription() string {
    return r.description
}

func (r *AliexpressSolutionOrderFulfillRequest) SetLogisticsNo(logisticsNo string) error {
    r.logisticsNo = logisticsNo
    r.Set("logistics_no", logisticsNo)
    return nil
}

func (r AliexpressSolutionOrderFulfillRequest) GetLogisticsNo() string {
    return r.logisticsNo
}

