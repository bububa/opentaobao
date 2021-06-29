package cainiaohandover

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取交接单小包信息 API请求
cainiao.global.handover.parcel.query

提供给ISV通过该接口查询小包信息
*/
type CainiaoGlobalHandoverParcelQueryRequest struct {
    model.Params
    // 用户信息
    userInfo   *UserInfoDto
    // 多语言
    locale   string
    // 小包的物流订单号,和小包的国际运单号参数任选其一即可
    orderCode   string
    // 小包的国际运单号，和小包的物流订单号参数任选其一即可
    trackingNumber   string
    // 客户端名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
    client   string
}

// 初始化CainiaoGlobalHandoverParcelQueryRequest对象
func NewCainiaoGlobalHandoverParcelQueryRequest() *CainiaoGlobalHandoverParcelQueryRequest{
    return &CainiaoGlobalHandoverParcelQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoGlobalHandoverParcelQueryRequest) GetApiMethodName() string {
    return "cainiao.global.handover.parcel.query"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoGlobalHandoverParcelQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserInfo Setter
// 用户信息
func (r *CainiaoGlobalHandoverParcelQueryRequest) SetUserInfo(userInfo *UserInfoDto) error {
    r.userInfo = userInfo
    r.Set("user_info", userInfo)
    return nil
}

// UserInfo Getter
func (r CainiaoGlobalHandoverParcelQueryRequest) GetUserInfo() *UserInfoDto {
    return r.userInfo
}
// Locale Setter
// 多语言
func (r *CainiaoGlobalHandoverParcelQueryRequest) SetLocale(locale string) error {
    r.locale = locale
    r.Set("locale", locale)
    return nil
}

// Locale Getter
func (r CainiaoGlobalHandoverParcelQueryRequest) GetLocale() string {
    return r.locale
}
// OrderCode Setter
// 小包的物流订单号,和小包的国际运单号参数任选其一即可
func (r *CainiaoGlobalHandoverParcelQueryRequest) SetOrderCode(orderCode string) error {
    r.orderCode = orderCode
    r.Set("order_code", orderCode)
    return nil
}

// OrderCode Getter
func (r CainiaoGlobalHandoverParcelQueryRequest) GetOrderCode() string {
    return r.orderCode
}
// TrackingNumber Setter
// 小包的国际运单号，和小包的物流订单号参数任选其一即可
func (r *CainiaoGlobalHandoverParcelQueryRequest) SetTrackingNumber(trackingNumber string) error {
    r.trackingNumber = trackingNumber
    r.Set("tracking_number", trackingNumber)
    return nil
}

// TrackingNumber Getter
func (r CainiaoGlobalHandoverParcelQueryRequest) GetTrackingNumber() string {
    return r.trackingNumber
}
// Client Setter
// 客户端名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
func (r *CainiaoGlobalHandoverParcelQueryRequest) SetClient(client string) error {
    r.client = client
    r.Set("client", client)
    return nil
}

// Client Getter
func (r CainiaoGlobalHandoverParcelQueryRequest) GetClient() string {
    return r.client
}
