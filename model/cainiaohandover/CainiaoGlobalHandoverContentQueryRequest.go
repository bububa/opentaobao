package cainiaohandover

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询大包详情 API请求
cainiao.global.handover.content.query

查询大包详情
*/
type CainiaoGlobalHandoverContentQueryRequest struct {
    model.Params
    // 用户信息
    userInfo   *UserInfoDto
    // 交接物运单号，和交接物物流订单编码参数任选其一即可
    trackingNumber   string
    // 客户端名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
    client   string
    // 交接物物流订单编码,和交接物运单号参数可以任选其一即可
    orderCode   string
    // 多语言
    locale   string
}

// 初始化CainiaoGlobalHandoverContentQueryRequest对象
func NewCainiaoGlobalHandoverContentQueryRequest() *CainiaoGlobalHandoverContentQueryRequest{
    return &CainiaoGlobalHandoverContentQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoGlobalHandoverContentQueryRequest) GetApiMethodName() string {
    return "cainiao.global.handover.content.query"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoGlobalHandoverContentQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserInfo Setter
// 用户信息
func (r *CainiaoGlobalHandoverContentQueryRequest) SetUserInfo(userInfo *UserInfoDto) error {
    r.userInfo = userInfo
    r.Set("user_info", userInfo)
    return nil
}

// UserInfo Getter
func (r CainiaoGlobalHandoverContentQueryRequest) GetUserInfo() *UserInfoDto {
    return r.userInfo
}
// TrackingNumber Setter
// 交接物运单号，和交接物物流订单编码参数任选其一即可
func (r *CainiaoGlobalHandoverContentQueryRequest) SetTrackingNumber(trackingNumber string) error {
    r.trackingNumber = trackingNumber
    r.Set("tracking_number", trackingNumber)
    return nil
}

// TrackingNumber Getter
func (r CainiaoGlobalHandoverContentQueryRequest) GetTrackingNumber() string {
    return r.trackingNumber
}
// Client Setter
// 客户端名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
func (r *CainiaoGlobalHandoverContentQueryRequest) SetClient(client string) error {
    r.client = client
    r.Set("client", client)
    return nil
}

// Client Getter
func (r CainiaoGlobalHandoverContentQueryRequest) GetClient() string {
    return r.client
}
// OrderCode Setter
// 交接物物流订单编码,和交接物运单号参数可以任选其一即可
func (r *CainiaoGlobalHandoverContentQueryRequest) SetOrderCode(orderCode string) error {
    r.orderCode = orderCode
    r.Set("order_code", orderCode)
    return nil
}

// OrderCode Getter
func (r CainiaoGlobalHandoverContentQueryRequest) GetOrderCode() string {
    return r.orderCode
}
// Locale Setter
// 多语言
func (r *CainiaoGlobalHandoverContentQueryRequest) SetLocale(locale string) error {
    r.locale = locale
    r.Set("locale", locale)
    return nil
}

// Locale Getter
func (r CainiaoGlobalHandoverContentQueryRequest) GetLocale() string {
    return r.locale
}
