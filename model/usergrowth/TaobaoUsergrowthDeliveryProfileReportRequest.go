package usergrowth

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
标签上报 API请求
taobao.usergrowth.delivery.profile.report

渠道上报标签信息
*/
type TaobaoUsergrowthDeliveryProfileReportRequest struct {
    model.Params
    // 标签参数, 支持一次传多个， 一次最多传20个
    _data   string
    // 渠道标识，找淘宝运营申请
    _channel   string
}

// 初始化TaobaoUsergrowthDeliveryProfileReportRequest对象
func NewTaobaoUsergrowthDeliveryProfileReportRequest() *TaobaoUsergrowthDeliveryProfileReportRequest{
    return &TaobaoUsergrowthDeliveryProfileReportRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoUsergrowthDeliveryProfileReportRequest) GetApiMethodName() string {
    return "taobao.usergrowth.delivery.profile.report"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoUsergrowthDeliveryProfileReportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Data Setter
// 标签参数, 支持一次传多个， 一次最多传20个
func (r *TaobaoUsergrowthDeliveryProfileReportRequest) SetData(_data string) error {
    r._data = _data
    r.Set("data", _data)
    return nil
}

// Data Getter
func (r TaobaoUsergrowthDeliveryProfileReportRequest) GetData() string {
    return r._data
}
// Channel Setter
// 渠道标识，找淘宝运营申请
func (r *TaobaoUsergrowthDeliveryProfileReportRequest) SetChannel(_channel string) error {
    r._channel = _channel
    r.Set("channel", _channel)
    return nil
}

// Channel Getter
func (r TaobaoUsergrowthDeliveryProfileReportRequest) GetChannel() string {
    return r._channel
}
