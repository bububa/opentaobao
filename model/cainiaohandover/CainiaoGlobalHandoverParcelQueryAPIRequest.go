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
type CainiaoGlobalHandoverParcelQueryAPIRequest struct {
    model.Params
    // 用户信息
    _userInfo   *UserInfoDto
    // 多语言
    _locale   string
    // 小包的物流订单号,和小包的国际运单号参数任选其一即可
    _orderCode   string
    // 小包的国际运单号，和小包的物流订单号参数任选其一即可
    _trackingNumber   string
    // 客户端名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
    _client   string
}

// 初始化CainiaoGlobalHandoverParcelQueryAPIRequest对象
func NewCainiaoGlobalHandoverParcelQueryRequest() *CainiaoGlobalHandoverParcelQueryAPIRequest{
    return &CainiaoGlobalHandoverParcelQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoGlobalHandoverParcelQueryAPIRequest) GetApiMethodName() string {
    return "cainiao.global.handover.parcel.query"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoGlobalHandoverParcelQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserInfo Setter
// 用户信息
func (r *CainiaoGlobalHandoverParcelQueryAPIRequest) SetUserInfo(_userInfo *UserInfoDto) error {
    r._userInfo = _userInfo
    r.Set("user_info", _userInfo)
    return nil
}

// UserInfo Getter
func (r CainiaoGlobalHandoverParcelQueryAPIRequest) GetUserInfo() *UserInfoDto {
    return r._userInfo
}
// Locale Setter
// 多语言
func (r *CainiaoGlobalHandoverParcelQueryAPIRequest) SetLocale(_locale string) error {
    r._locale = _locale
    r.Set("locale", _locale)
    return nil
}

// Locale Getter
func (r CainiaoGlobalHandoverParcelQueryAPIRequest) GetLocale() string {
    return r._locale
}
// OrderCode Setter
// 小包的物流订单号,和小包的国际运单号参数任选其一即可
func (r *CainiaoGlobalHandoverParcelQueryAPIRequest) SetOrderCode(_orderCode string) error {
    r._orderCode = _orderCode
    r.Set("order_code", _orderCode)
    return nil
}

// OrderCode Getter
func (r CainiaoGlobalHandoverParcelQueryAPIRequest) GetOrderCode() string {
    return r._orderCode
}
// TrackingNumber Setter
// 小包的国际运单号，和小包的物流订单号参数任选其一即可
func (r *CainiaoGlobalHandoverParcelQueryAPIRequest) SetTrackingNumber(_trackingNumber string) error {
    r._trackingNumber = _trackingNumber
    r.Set("tracking_number", _trackingNumber)
    return nil
}

// TrackingNumber Getter
func (r CainiaoGlobalHandoverParcelQueryAPIRequest) GetTrackingNumber() string {
    return r._trackingNumber
}
// Client Setter
// 客户端名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
func (r *CainiaoGlobalHandoverParcelQueryAPIRequest) SetClient(_client string) error {
    r._client = _client
    r.Set("client", _client)
    return nil
}

// Client Getter
func (r CainiaoGlobalHandoverParcelQueryAPIRequest) GetClient() string {
    return r._client
}
