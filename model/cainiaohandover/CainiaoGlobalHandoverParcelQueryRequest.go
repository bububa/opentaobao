package cainiaohandover

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取交接单小包信息 APIRequest
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

func NewCainiaoGlobalHandoverParcelQueryRequest() *CainiaoGlobalHandoverParcelQueryRequest{
    return &CainiaoGlobalHandoverParcelQueryRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoGlobalHandoverParcelQueryRequest) GetApiMethodName() string {
    return "cainiao.global.handover.parcel.query"
}

func (r CainiaoGlobalHandoverParcelQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoGlobalHandoverParcelQueryRequest) SetUserInfo(userInfo *UserInfoDto) error {
    r.userInfo = userInfo
    r.Set("user_info", userInfo)
    return nil
}

func (r CainiaoGlobalHandoverParcelQueryRequest) GetUserInfo() *UserInfoDto {
    return r.userInfo
}

func (r *CainiaoGlobalHandoverParcelQueryRequest) SetLocale(locale string) error {
    r.locale = locale
    r.Set("locale", locale)
    return nil
}

func (r CainiaoGlobalHandoverParcelQueryRequest) GetLocale() string {
    return r.locale
}

func (r *CainiaoGlobalHandoverParcelQueryRequest) SetOrderCode(orderCode string) error {
    r.orderCode = orderCode
    r.Set("order_code", orderCode)
    return nil
}

func (r CainiaoGlobalHandoverParcelQueryRequest) GetOrderCode() string {
    return r.orderCode
}

func (r *CainiaoGlobalHandoverParcelQueryRequest) SetTrackingNumber(trackingNumber string) error {
    r.trackingNumber = trackingNumber
    r.Set("tracking_number", trackingNumber)
    return nil
}

func (r CainiaoGlobalHandoverParcelQueryRequest) GetTrackingNumber() string {
    return r.trackingNumber
}

func (r *CainiaoGlobalHandoverParcelQueryRequest) SetClient(client string) error {
    r.client = client
    r.Set("client", client)
    return nil
}

func (r CainiaoGlobalHandoverParcelQueryRequest) GetClient() string {
    return r.client
}

