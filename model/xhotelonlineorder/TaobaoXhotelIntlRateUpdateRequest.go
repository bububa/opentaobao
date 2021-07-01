package xhotelonlineorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
不落库商家推送更新酒店rate API请求
taobao.xhotel.intl.rate.update

商家主动推送不落库商品的酒店信息
*/
type TaobaoXhotelIntlRateUpdateAPIRequest struct {
    model.Params
    // rate更新参数
    _updateRateParam   *UpdateRateParam
}

// 初始化TaobaoXhotelIntlRateUpdateAPIRequest对象
func NewTaobaoXhotelIntlRateUpdateRequest() *TaobaoXhotelIntlRateUpdateAPIRequest{
    return &TaobaoXhotelIntlRateUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelIntlRateUpdateAPIRequest) GetApiMethodName() string {
    return "taobao.xhotel.intl.rate.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelIntlRateUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UpdateRateParam Setter
// rate更新参数
func (r *TaobaoXhotelIntlRateUpdateAPIRequest) SetUpdateRateParam(_updateRateParam *UpdateRateParam) error {
    r._updateRateParam = _updateRateParam
    r.Set("update_rate_param", _updateRateParam)
    return nil
}

// UpdateRateParam Getter
func (r TaobaoXhotelIntlRateUpdateAPIRequest) GetUpdateRateParam() *UpdateRateParam {
    return r._updateRateParam
}
