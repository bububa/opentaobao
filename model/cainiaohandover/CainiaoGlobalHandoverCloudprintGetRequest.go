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
    userInfo   *UserInfoDto
    // 大包运单号
    trackingNumber   string
    // ISV名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
    client   string
    // 大包物流单LP号
    orderCode   string
    // 多语言
    locale   string
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
func (r *CainiaoGlobalHandoverCloudprintGetRequest) SetUserInfo(userInfo *UserInfoDto) error {
    r.userInfo = userInfo
    r.Set("user_info", userInfo)
    return nil
}

// UserInfo Getter
func (r CainiaoGlobalHandoverCloudprintGetRequest) GetUserInfo() *UserInfoDto {
    return r.userInfo
}
// TrackingNumber Setter
// 大包运单号
func (r *CainiaoGlobalHandoverCloudprintGetRequest) SetTrackingNumber(trackingNumber string) error {
    r.trackingNumber = trackingNumber
    r.Set("tracking_number", trackingNumber)
    return nil
}

// TrackingNumber Getter
func (r CainiaoGlobalHandoverCloudprintGetRequest) GetTrackingNumber() string {
    return r.trackingNumber
}
// Client Setter
// ISV名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
func (r *CainiaoGlobalHandoverCloudprintGetRequest) SetClient(client string) error {
    r.client = client
    r.Set("client", client)
    return nil
}

// Client Getter
func (r CainiaoGlobalHandoverCloudprintGetRequest) GetClient() string {
    return r.client
}
// OrderCode Setter
// 大包物流单LP号
func (r *CainiaoGlobalHandoverCloudprintGetRequest) SetOrderCode(orderCode string) error {
    r.orderCode = orderCode
    r.Set("order_code", orderCode)
    return nil
}

// OrderCode Getter
func (r CainiaoGlobalHandoverCloudprintGetRequest) GetOrderCode() string {
    return r.orderCode
}
// Locale Setter
// 多语言
func (r *CainiaoGlobalHandoverCloudprintGetRequest) SetLocale(locale string) error {
    r.locale = locale
    r.Set("locale", locale)
    return nil
}

// Locale Getter
func (r CainiaoGlobalHandoverCloudprintGetRequest) GetLocale() string {
    return r.locale
}
