package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
快递改约api API请求
taobao.logistics.express.modify.appoint

商家通过此api操作修改物流单，交易单的收货人地址、收货人联系方式、预约配送日期
*/
type TaobaoLogisticsExpressModifyAppointRequest struct {
    model.Params
    // 改约请求对象
    _expressModifyAppointTopRequest   *ExpressModifyAppointTopRequestDto
}

// 初始化TaobaoLogisticsExpressModifyAppointRequest对象
func NewTaobaoLogisticsExpressModifyAppointRequest() *TaobaoLogisticsExpressModifyAppointRequest{
    return &TaobaoLogisticsExpressModifyAppointRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsExpressModifyAppointRequest) GetApiMethodName() string {
    return "taobao.logistics.express.modify.appoint"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsExpressModifyAppointRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ExpressModifyAppointTopRequest Setter
// 改约请求对象
func (r *TaobaoLogisticsExpressModifyAppointRequest) SetExpressModifyAppointTopRequest(_expressModifyAppointTopRequest *ExpressModifyAppointTopRequestDto) error {
    r._expressModifyAppointTopRequest = _expressModifyAppointTopRequest
    r.Set("express_modify_appoint_top_request", _expressModifyAppointTopRequest)
    return nil
}

// ExpressModifyAppointTopRequest Getter
func (r TaobaoLogisticsExpressModifyAppointRequest) GetExpressModifyAppointTopRequest() *ExpressModifyAppointTopRequestDto {
    return r._expressModifyAppointTopRequest
}
