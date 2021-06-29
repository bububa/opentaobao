package xhotelonlineorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
不落库商家推送更新酒店rate APIRequest
taobao.xhotel.intl.rate.update

商家主动推送不落库商品的酒店信息
*/
type TaobaoXhotelIntlRateUpdateRequest struct {
    model.Params

    // rate更新参数
    updateRateParam   *UpdateRateParam 

}

func NewTaobaoXhotelIntlRateUpdateRequest() *TaobaoXhotelIntlRateUpdateRequest{
    return &TaobaoXhotelIntlRateUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoXhotelIntlRateUpdateRequest) GetApiMethodName() string {
    return "taobao.xhotel.intl.rate.update"
}

func (r TaobaoXhotelIntlRateUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoXhotelIntlRateUpdateRequest) SetUpdateRateParam(updateRateParam *UpdateRateParam) error {
    r.updateRateParam = updateRateParam
    r.Set("update_rate_param", updateRateParam)
    return nil
}

func (r TaobaoXhotelIntlRateUpdateRequest) GetUpdateRateParam() *UpdateRateParam {
    return r.updateRateParam
}

