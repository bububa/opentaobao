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
    _userInfo   *UserInfoDto
    // 要取消的交接物运单号，即大包运单号
    _trackingNumber   string
    // 要取消的交接单id
    _handoverOrderId   int64
    // 要取消的交接物id，即大包id
    _handoverContentId   int64
    // ISV名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
    _client   string
    // 多语言
    _locale   string
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
func (r *CainiaoGlobalHandoverCancelRequest) SetUserInfo(_userInfo *UserInfoDto) error {
    r._userInfo = _userInfo
    r.Set("user_info", _userInfo)
    return nil
}

// UserInfo Getter
func (r CainiaoGlobalHandoverCancelRequest) GetUserInfo() *UserInfoDto {
    return r._userInfo
}
// TrackingNumber Setter
// 要取消的交接物运单号，即大包运单号
func (r *CainiaoGlobalHandoverCancelRequest) SetTrackingNumber(_trackingNumber string) error {
    r._trackingNumber = _trackingNumber
    r.Set("tracking_number", _trackingNumber)
    return nil
}

// TrackingNumber Getter
func (r CainiaoGlobalHandoverCancelRequest) GetTrackingNumber() string {
    return r._trackingNumber
}
// HandoverOrderId Setter
// 要取消的交接单id
func (r *CainiaoGlobalHandoverCancelRequest) SetHandoverOrderId(_handoverOrderId int64) error {
    r._handoverOrderId = _handoverOrderId
    r.Set("handover_order_id", _handoverOrderId)
    return nil
}

// HandoverOrderId Getter
func (r CainiaoGlobalHandoverCancelRequest) GetHandoverOrderId() int64 {
    return r._handoverOrderId
}
// HandoverContentId Setter
// 要取消的交接物id，即大包id
func (r *CainiaoGlobalHandoverCancelRequest) SetHandoverContentId(_handoverContentId int64) error {
    r._handoverContentId = _handoverContentId
    r.Set("handover_content_id", _handoverContentId)
    return nil
}

// HandoverContentId Getter
func (r CainiaoGlobalHandoverCancelRequest) GetHandoverContentId() int64 {
    return r._handoverContentId
}
// Client Setter
// ISV名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
func (r *CainiaoGlobalHandoverCancelRequest) SetClient(_client string) error {
    r._client = _client
    r.Set("client", _client)
    return nil
}

// Client Getter
func (r CainiaoGlobalHandoverCancelRequest) GetClient() string {
    return r._client
}
// Locale Setter
// 多语言
func (r *CainiaoGlobalHandoverCancelRequest) SetLocale(_locale string) error {
    r._locale = _locale
    r.Set("locale", _locale)
    return nil
}

// Locale Getter
func (r CainiaoGlobalHandoverCancelRequest) GetLocale() string {
    return r._locale
}