package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
快递改约api APIRequest
taobao.logistics.express.modify.appoint

商家通过此api操作修改物流单，交易单的收货人地址、收货人联系方式、预约配送日期
*/
type TaobaoLogisticsExpressModifyAppointRequest struct {
    model.Params

    // 改约请求对象
    expressModifyAppointTopRequest   *ExpressModifyAppointTopRequestDto 

}

func NewTaobaoLogisticsExpressModifyAppointRequest() *TaobaoLogisticsExpressModifyAppointRequest{
    return &TaobaoLogisticsExpressModifyAppointRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoLogisticsExpressModifyAppointRequest) GetApiMethodName() string {
    return "taobao.logistics.express.modify.appoint"
}

func (r TaobaoLogisticsExpressModifyAppointRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoLogisticsExpressModifyAppointRequest) SetExpressModifyAppointTopRequest(expressModifyAppointTopRequest *ExpressModifyAppointTopRequestDto) error {
    r.expressModifyAppointTopRequest = expressModifyAppointTopRequest
    r.Set("express_modify_appoint_top_request", expressModifyAppointTopRequest)
    return nil
}

func (r TaobaoLogisticsExpressModifyAppointRequest) GetExpressModifyAppointTopRequest() *ExpressModifyAppointTopRequestDto {
    return r.expressModifyAppointTopRequest
}

