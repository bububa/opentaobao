package cainiaohandover

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
实际承运商查询 API请求
cainiao.global.logistics.carrier.querylist

查询出所有的实际承运商
*/
type CainiaoGlobalLogisticsCarrierQuerylistRequest struct {
    model.Params
    // 多语言(暂不支持，保留入参)
    locale   string
}

// 初始化CainiaoGlobalLogisticsCarrierQuerylistRequest对象
func NewCainiaoGlobalLogisticsCarrierQuerylistRequest() *CainiaoGlobalLogisticsCarrierQuerylistRequest{
    return &CainiaoGlobalLogisticsCarrierQuerylistRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoGlobalLogisticsCarrierQuerylistRequest) GetApiMethodName() string {
    return "cainiao.global.logistics.carrier.querylist"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoGlobalLogisticsCarrierQuerylistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Locale Setter
// 多语言(暂不支持，保留入参)
func (r *CainiaoGlobalLogisticsCarrierQuerylistRequest) SetLocale(locale string) error {
    r.locale = locale
    r.Set("locale", locale)
    return nil
}

// Locale Getter
func (r CainiaoGlobalLogisticsCarrierQuerylistRequest) GetLocale() string {
    return r.locale
}
