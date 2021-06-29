package cainiaohandover

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取面单云打印数据 APIRequest
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

func NewCainiaoGlobalHandoverCloudprintGetRequest() *CainiaoGlobalHandoverCloudprintGetRequest{
    return &CainiaoGlobalHandoverCloudprintGetRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoGlobalHandoverCloudprintGetRequest) GetApiMethodName() string {
    return "cainiao.global.handover.cloudprint.get"
}

func (r CainiaoGlobalHandoverCloudprintGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoGlobalHandoverCloudprintGetRequest) SetUserInfo(userInfo *UserInfoDto) error {
    r.userInfo = userInfo
    r.Set("user_info", userInfo)
    return nil
}

func (r CainiaoGlobalHandoverCloudprintGetRequest) GetUserInfo() *UserInfoDto {
    return r.userInfo
}

func (r *CainiaoGlobalHandoverCloudprintGetRequest) SetTrackingNumber(trackingNumber string) error {
    r.trackingNumber = trackingNumber
    r.Set("tracking_number", trackingNumber)
    return nil
}

func (r CainiaoGlobalHandoverCloudprintGetRequest) GetTrackingNumber() string {
    return r.trackingNumber
}

func (r *CainiaoGlobalHandoverCloudprintGetRequest) SetClient(client string) error {
    r.client = client
    r.Set("client", client)
    return nil
}

func (r CainiaoGlobalHandoverCloudprintGetRequest) GetClient() string {
    return r.client
}

func (r *CainiaoGlobalHandoverCloudprintGetRequest) SetOrderCode(orderCode string) error {
    r.orderCode = orderCode
    r.Set("order_code", orderCode)
    return nil
}

func (r CainiaoGlobalHandoverCloudprintGetRequest) GetOrderCode() string {
    return r.orderCode
}

func (r *CainiaoGlobalHandoverCloudprintGetRequest) SetLocale(locale string) error {
    r.locale = locale
    r.Set("locale", locale)
    return nil
}

func (r CainiaoGlobalHandoverCloudprintGetRequest) GetLocale() string {
    return r.locale
}

