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
type TaobaoUsergrowthDeliveryProfileReportAPIRequest struct {
    model.Params
    // 标签参数, 支持一次传多个， 一次最多传20个
    _data   string
    // 渠道标识，找淘宝运营申请
    _channel   string
}

// 初始化TaobaoUsergrowthDeliveryProfileReportAPIRequest对象
func NewTaobaoUsergrowthDeliveryProfileReportRequest() *TaobaoUsergrowthDeliveryProfileReportAPIRequest{
    return &TaobaoUsergrowthDeliveryProfileReportAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoUsergrowthDeliveryProfileReportAPIRequest) GetApiMethodName() string {
    return "taobao.usergrowth.delivery.profile.report"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoUsergrowthDeliveryProfileReportAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Data Setter
// 标签参数, 支持一次传多个， 一次最多传20个
func (r *TaobaoUsergrowthDeliveryProfileReportAPIRequest) SetData(_data string) error {
    r._data = _data
    r.Set("data", _data)
    return nil
}

// Data Getter
func (r TaobaoUsergrowthDeliveryProfileReportAPIRequest) GetData() string {
    return r._data
}
// Channel Setter
// 渠道标识，找淘宝运营申请
func (r *TaobaoUsergrowthDeliveryProfileReportAPIRequest) SetChannel(_channel string) error {
    r._channel = _channel
    r.Set("channel", _channel)
    return nil
}

// Channel Getter
func (r TaobaoUsergrowthDeliveryProfileReportAPIRequest) GetChannel() string {
    return r._channel
}
