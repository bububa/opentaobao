package cainiaohandover

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询大包详情 APIRequest
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

func NewCainiaoGlobalHandoverContentQueryRequest() *CainiaoGlobalHandoverContentQueryRequest{
    return &CainiaoGlobalHandoverContentQueryRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoGlobalHandoverContentQueryRequest) GetApiMethodName() string {
    return "cainiao.global.handover.content.query"
}

func (r CainiaoGlobalHandoverContentQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoGlobalHandoverContentQueryRequest) SetUserInfo(userInfo *UserInfoDto) error {
    r.userInfo = userInfo
    r.Set("user_info", userInfo)
    return nil
}

func (r CainiaoGlobalHandoverContentQueryRequest) GetUserInfo() *UserInfoDto {
    return r.userInfo
}

func (r *CainiaoGlobalHandoverContentQueryRequest) SetTrackingNumber(trackingNumber string) error {
    r.trackingNumber = trackingNumber
    r.Set("tracking_number", trackingNumber)
    return nil
}

func (r CainiaoGlobalHandoverContentQueryRequest) GetTrackingNumber() string {
    return r.trackingNumber
}

func (r *CainiaoGlobalHandoverContentQueryRequest) SetClient(client string) error {
    r.client = client
    r.Set("client", client)
    return nil
}

func (r CainiaoGlobalHandoverContentQueryRequest) GetClient() string {
    return r.client
}

func (r *CainiaoGlobalHandoverContentQueryRequest) SetOrderCode(orderCode string) error {
    r.orderCode = orderCode
    r.Set("order_code", orderCode)
    return nil
}

func (r CainiaoGlobalHandoverContentQueryRequest) GetOrderCode() string {
    return r.orderCode
}

func (r *CainiaoGlobalHandoverContentQueryRequest) SetLocale(locale string) error {
    r.locale = locale
    r.Set("locale", locale)
    return nil
}

func (r CainiaoGlobalHandoverContentQueryRequest) GetLocale() string {
    return r.locale
}

