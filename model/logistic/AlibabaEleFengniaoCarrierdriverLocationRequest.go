package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询骑手当前位置 APIRequest
alibaba.ele.fengniao.carrierdriver.location

查询骑手当前位置
*/
type AlibabaEleFengniaoCarrierdriverLocationRequest struct {
    model.Params

    // appid
    appId   string 

    // 外部订单号
    partnerOrderCode   string 

}

func NewAlibabaEleFengniaoCarrierdriverLocationRequest() *AlibabaEleFengniaoCarrierdriverLocationRequest{
    return &AlibabaEleFengniaoCarrierdriverLocationRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEleFengniaoCarrierdriverLocationRequest) GetApiMethodName() string {
    return "alibaba.ele.fengniao.carrierdriver.location"
}

func (r AlibabaEleFengniaoCarrierdriverLocationRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEleFengniaoCarrierdriverLocationRequest) SetAppId(appId string) error {
    r.appId = appId
    r.Set("app_id", appId)
    return nil
}

func (r AlibabaEleFengniaoCarrierdriverLocationRequest) GetAppId() string {
    return r.appId
}

func (r *AlibabaEleFengniaoCarrierdriverLocationRequest) SetPartnerOrderCode(partnerOrderCode string) error {
    r.partnerOrderCode = partnerOrderCode
    r.Set("partner_order_code", partnerOrderCode)
    return nil
}

func (r AlibabaEleFengniaoCarrierdriverLocationRequest) GetPartnerOrderCode() string {
    return r.partnerOrderCode
}

