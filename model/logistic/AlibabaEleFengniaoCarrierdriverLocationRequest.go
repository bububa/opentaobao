package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询骑手当前位置 API请求
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

// 初始化AlibabaEleFengniaoCarrierdriverLocationRequest对象
func NewAlibabaEleFengniaoCarrierdriverLocationRequest() *AlibabaEleFengniaoCarrierdriverLocationRequest{
    return &AlibabaEleFengniaoCarrierdriverLocationRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEleFengniaoCarrierdriverLocationRequest) GetApiMethodName() string {
    return "alibaba.ele.fengniao.carrierdriver.location"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEleFengniaoCarrierdriverLocationRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AppId Setter
// appid
func (r *AlibabaEleFengniaoCarrierdriverLocationRequest) SetAppId(appId string) error {
    r.appId = appId
    r.Set("app_id", appId)
    return nil
}

// AppId Getter
func (r AlibabaEleFengniaoCarrierdriverLocationRequest) GetAppId() string {
    return r.appId
}
// PartnerOrderCode Setter
// 外部订单号
func (r *AlibabaEleFengniaoCarrierdriverLocationRequest) SetPartnerOrderCode(partnerOrderCode string) error {
    r.partnerOrderCode = partnerOrderCode
    r.Set("partner_order_code", partnerOrderCode)
    return nil
}

// PartnerOrderCode Getter
func (r AlibabaEleFengniaoCarrierdriverLocationRequest) GetPartnerOrderCode() string {
    return r.partnerOrderCode
}
