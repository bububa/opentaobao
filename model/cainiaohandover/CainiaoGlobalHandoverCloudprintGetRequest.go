package cainiaohandover

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取面单云打印数据 API请求
cainiao.global.handover.cloudprint.get

提供给ISV通过该接口获取面单云打印数据
*/
type CainiaoGlobalHandoverCloudprintGetRequest struct {
    model.Params
    // 用户信息
    _userInfo   *UserInfoDTO
    // 大包运单号
    _trackingNumber   string
    // ISV名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
    _client   string
    // 大包物流单LP号
    _orderCode   string
    // 多语言
    _locale   string
}

// 初始化CainiaoGlobalHandoverCloudprintGetRequest对象
func NewCainiaoGlobalHandoverCloudprintGetRequest() *CainiaoGlobalHandoverCloudprintGetRequest{
    return &CainiaoGlobalHandoverCloudprintGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoGlobalHandoverCloudprintGetRequest) GetApiMethodName() string {
    return "cainiao.global.handover.cloudprint.get"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoGlobalHandoverCloudprintGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserInfo Setter
// 用户信息
func (r *CainiaoGlobalHandoverCloudprintGetRequest) SetUserInfo(_userInfo *UserInfoDTO) error {
    r._userInfo = _userInfo
    r.Set("user_info", _userInfo)
    return nil
}

// UserInfo Getter
func (r CainiaoGlobalHandoverCloudprintGetRequest) GetUserInfo() *UserInfoDTO {
    return r._userInfo
}
// TrackingNumber Setter
// 大包运单号
func (r *CainiaoGlobalHandoverCloudprintGetRequest) SetTrackingNumber(_trackingNumber string) error {
    r._trackingNumber = _trackingNumber
    r.Set("tracking_number", _trackingNumber)
    return nil
}

// TrackingNumber Getter
func (r CainiaoGlobalHandoverCloudprintGetRequest) GetTrackingNumber() string {
    return r._trackingNumber
}
// Client Setter
// ISV名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
func (r *CainiaoGlobalHandoverCloudprintGetRequest) SetClient(_client string) error {
    r._client = _client
    r.Set("client", _client)
    return nil
}

// Client Getter
func (r CainiaoGlobalHandoverCloudprintGetRequest) GetClient() string {
    return r._client
}
// OrderCode Setter
// 大包物流单LP号
func (r *CainiaoGlobalHandoverCloudprintGetRequest) SetOrderCode(_orderCode string) error {
    r._orderCode = _orderCode
    r.Set("order_code", _orderCode)
    return nil
}

// OrderCode Getter
func (r CainiaoGlobalHandoverCloudprintGetRequest) GetOrderCode() string {
    return r._orderCode
}
// Locale Setter
// 多语言
func (r *CainiaoGlobalHandoverCloudprintGetRequest) SetLocale(_locale string) error {
    r._locale = _locale
    r.Set("locale", _locale)
    return nil
}

// Locale Getter
func (r CainiaoGlobalHandoverCloudprintGetRequest) GetLocale() string {
    return r._locale
}
