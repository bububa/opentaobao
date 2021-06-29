package cainiaohandover

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
取消交接单 APIRequest
cainiao.global.handover.cancel

提供给ISV通过该接口取消交接单
*/
type CainiaoGlobalHandoverCancelRequest struct {
    model.Params

    // 系统自动生成
    userInfo   *UserInfoDto 

    // 要取消的交接物运单号，即大包运单号
    trackingNumber   string 

    // 要取消的交接单id
    handoverOrderId   int64 

    // 要取消的交接物id，即大包id
    handoverContentId   int64 

    // ISV名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
    client   string 

    // 多语言
    locale   string 

}

func NewCainiaoGlobalHandoverCancelRequest() *CainiaoGlobalHandoverCancelRequest{
    return &CainiaoGlobalHandoverCancelRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoGlobalHandoverCancelRequest) GetApiMethodName() string {
    return "cainiao.global.handover.cancel"
}

func (r CainiaoGlobalHandoverCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoGlobalHandoverCancelRequest) SetUserInfo(userInfo *UserInfoDto) error {
    r.userInfo = userInfo
    r.Set("user_info", userInfo)
    return nil
}

func (r CainiaoGlobalHandoverCancelRequest) GetUserInfo() *UserInfoDto {
    return r.userInfo
}

func (r *CainiaoGlobalHandoverCancelRequest) SetTrackingNumber(trackingNumber string) error {
    r.trackingNumber = trackingNumber
    r.Set("tracking_number", trackingNumber)
    return nil
}

func (r CainiaoGlobalHandoverCancelRequest) GetTrackingNumber() string {
    return r.trackingNumber
}

func (r *CainiaoGlobalHandoverCancelRequest) SetHandoverOrderId(handoverOrderId int64) error {
    r.handoverOrderId = handoverOrderId
    r.Set("handover_order_id", handoverOrderId)
    return nil
}

func (r CainiaoGlobalHandoverCancelRequest) GetHandoverOrderId() int64 {
    return r.handoverOrderId
}

func (r *CainiaoGlobalHandoverCancelRequest) SetHandoverContentId(handoverContentId int64) error {
    r.handoverContentId = handoverContentId
    r.Set("handover_content_id", handoverContentId)
    return nil
}

func (r CainiaoGlobalHandoverCancelRequest) GetHandoverContentId() int64 {
    return r.handoverContentId
}

func (r *CainiaoGlobalHandoverCancelRequest) SetClient(client string) error {
    r.client = client
    r.Set("client", client)
    return nil
}

func (r CainiaoGlobalHandoverCancelRequest) GetClient() string {
    return r.client
}

func (r *CainiaoGlobalHandoverCancelRequest) SetLocale(locale string) error {
    r.locale = locale
    r.Set("locale", locale)
    return nil
}

func (r CainiaoGlobalHandoverCancelRequest) GetLocale() string {
    return r.locale
}

