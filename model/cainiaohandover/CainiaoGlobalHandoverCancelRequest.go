package cainiaohandover

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
取消交接单 API请求
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

// 初始化CainiaoGlobalHandoverCancelRequest对象
func NewCainiaoGlobalHandoverCancelRequest() *CainiaoGlobalHandoverCancelRequest{
    return &CainiaoGlobalHandoverCancelRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoGlobalHandoverCancelRequest) GetApiMethodName() string {
    return "cainiao.global.handover.cancel"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoGlobalHandoverCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserInfo Setter
// 系统自动生成
func (r *CainiaoGlobalHandoverCancelRequest) SetUserInfo(userInfo *UserInfoDto) error {
    r.userInfo = userInfo
    r.Set("user_info", userInfo)
    return nil
}

// UserInfo Getter
func (r CainiaoGlobalHandoverCancelRequest) GetUserInfo() *UserInfoDto {
    return r.userInfo
}
// TrackingNumber Setter
// 要取消的交接物运单号，即大包运单号
func (r *CainiaoGlobalHandoverCancelRequest) SetTrackingNumber(trackingNumber string) error {
    r.trackingNumber = trackingNumber
    r.Set("tracking_number", trackingNumber)
    return nil
}

// TrackingNumber Getter
func (r CainiaoGlobalHandoverCancelRequest) GetTrackingNumber() string {
    return r.trackingNumber
}
// HandoverOrderId Setter
// 要取消的交接单id
func (r *CainiaoGlobalHandoverCancelRequest) SetHandoverOrderId(handoverOrderId int64) error {
    r.handoverOrderId = handoverOrderId
    r.Set("handover_order_id", handoverOrderId)
    return nil
}

// HandoverOrderId Getter
func (r CainiaoGlobalHandoverCancelRequest) GetHandoverOrderId() int64 {
    return r.handoverOrderId
}
// HandoverContentId Setter
// 要取消的交接物id，即大包id
func (r *CainiaoGlobalHandoverCancelRequest) SetHandoverContentId(handoverContentId int64) error {
    r.handoverContentId = handoverContentId
    r.Set("handover_content_id", handoverContentId)
    return nil
}

// HandoverContentId Getter
func (r CainiaoGlobalHandoverCancelRequest) GetHandoverContentId() int64 {
    return r.handoverContentId
}
// Client Setter
// ISV名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
func (r *CainiaoGlobalHandoverCancelRequest) SetClient(client string) error {
    r.client = client
    r.Set("client", client)
    return nil
}

// Client Getter
func (r CainiaoGlobalHandoverCancelRequest) GetClient() string {
    return r.client
}
// Locale Setter
// 多语言
func (r *CainiaoGlobalHandoverCancelRequest) SetLocale(locale string) error {
    r.locale = locale
    r.Set("locale", locale)
    return nil
}

// Locale Getter
func (r CainiaoGlobalHandoverCancelRequest) GetLocale() string {
    return r.locale
}
