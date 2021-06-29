package usergrowth

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
标签上报 APIRequest
taobao.usergrowth.delivery.profile.report

渠道上报标签信息
*/
type TaobaoUsergrowthDeliveryProfileReportRequest struct {
    model.Params

    // 标签参数, 支持一次传多个， 一次最多传20个
    data   string 

    // 渠道标识，找淘宝运营申请
    channel   string 

}

func NewTaobaoUsergrowthDeliveryProfileReportRequest() *TaobaoUsergrowthDeliveryProfileReportRequest{
    return &TaobaoUsergrowthDeliveryProfileReportRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoUsergrowthDeliveryProfileReportRequest) GetApiMethodName() string {
    return "taobao.usergrowth.delivery.profile.report"
}

func (r TaobaoUsergrowthDeliveryProfileReportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoUsergrowthDeliveryProfileReportRequest) SetData(data string) error {
    r.data = data
    r.Set("data", data)
    return nil
}

func (r TaobaoUsergrowthDeliveryProfileReportRequest) GetData() string {
    return r.data
}

func (r *TaobaoUsergrowthDeliveryProfileReportRequest) SetChannel(channel string) error {
    r.channel = channel
    r.Set("channel", channel)
    return nil
}

func (r TaobaoUsergrowthDeliveryProfileReportRequest) GetChannel() string {
    return r.channel
}

