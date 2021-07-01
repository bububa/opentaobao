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
type CainiaoGlobalHandoverContentQueryAPIRequest struct {
    model.Params
    // 用户信息
    _userInfo   *UserInfoDto
    // 交接物运单号，和交接物物流订单编码参数任选其一即可
    _trackingNumber   string
    // 客户端名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
    _client   string
    // 交接物物流订单编码,和交接物运单号参数可以任选其一即可
    _orderCode   string
    // 多语言
    _locale   string
}

// 初始化CainiaoGlobalHandoverContentQueryAPIRequest对象
func NewCainiaoGlobalHandoverContentQueryRequest() *CainiaoGlobalHandoverContentQueryAPIRequest{
    return &CainiaoGlobalHandoverContentQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoGlobalHandoverContentQueryAPIRequest) GetApiMethodName() string {
    return "cainiao.global.handover.content.query"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoGlobalHandoverContentQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserInfo Setter
// 用户信息
func (r *CainiaoGlobalHandoverContentQueryAPIRequest) SetUserInfo(_userInfo *UserInfoDto) error {
    r._userInfo = _userInfo
    r.Set("user_info", _userInfo)
    return nil
}

// UserInfo Getter
func (r CainiaoGlobalHandoverContentQueryAPIRequest) GetUserInfo() *UserInfoDto {
    return r._userInfo
}
// TrackingNumber Setter
// 交接物运单号，和交接物物流订单编码参数任选其一即可
func (r *CainiaoGlobalHandoverContentQueryAPIRequest) SetTrackingNumber(_trackingNumber string) error {
    r._trackingNumber = _trackingNumber
    r.Set("tracking_number", _trackingNumber)
    return nil
}

// TrackingNumber Getter
func (r CainiaoGlobalHandoverContentQueryAPIRequest) GetTrackingNumber() string {
    return r._trackingNumber
}
// Client Setter
// 客户端名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
func (r *CainiaoGlobalHandoverContentQueryAPIRequest) SetClient(_client string) error {
    r._client = _client
    r.Set("client", _client)
    return nil
}

// Client Getter
func (r CainiaoGlobalHandoverContentQueryAPIRequest) GetClient() string {
    return r._client
}
// OrderCode Setter
// 交接物物流订单编码,和交接物运单号参数可以任选其一即可
func (r *CainiaoGlobalHandoverContentQueryAPIRequest) SetOrderCode(_orderCode string) error {
    r._orderCode = _orderCode
    r.Set("order_code", _orderCode)
    return nil
}

// OrderCode Getter
func (r CainiaoGlobalHandoverContentQueryAPIRequest) GetOrderCode() string {
    return r._orderCode
}
// Locale Setter
// 多语言
func (r *CainiaoGlobalHandoverContentQueryAPIRequest) SetLocale(_locale string) error {
    r._locale = _locale
    r.Set("locale", _locale)
    return nil
}

// Locale Getter
func (r CainiaoGlobalHandoverContentQueryAPIRequest) GetLocale() string {
    return r._locale
}
