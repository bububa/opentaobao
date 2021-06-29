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
    _appId   string
    // 外部订单号
    _partnerOrderCode   string
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
func (r *AlibabaEleFengniaoCarrierdriverLocationRequest) SetAppId(_appId string) error {
    r._appId = _appId
    r.Set("app_id", _appId)
    return nil
}

// AppId Getter
func (r AlibabaEleFengniaoCarrierdriverLocationRequest) GetAppId() string {
    return r._appId
}
// PartnerOrderCode Setter
// 外部订单号
func (r *AlibabaEleFengniaoCarrierdriverLocationRequest) SetPartnerOrderCode(_partnerOrderCode string) error {
    r._partnerOrderCode = _partnerOrderCode
    r.Set("partner_order_code", _partnerOrderCode)
    return nil
}

// PartnerOrderCode Getter
func (r AlibabaEleFengniaoCarrierdriverLocationRequest) GetPartnerOrderCode() string {
    return r._partnerOrderCode
}
