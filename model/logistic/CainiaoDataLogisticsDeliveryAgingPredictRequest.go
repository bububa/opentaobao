package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
配送物流时效预测 API请求
cainiao.data.logistics.delivery.aging.predict

时效和服务预期是商家发货时比较关注的信息，也是选择快递公司的一个重要参考（除去长期合作合同因素）。所以，在商家发货时给商家提供线路时效预估能帮助商家选择更好的快递公司，且对消费者来说也能整体提升体验。

日常，展示具体的预测时效数值

大促期间，展示预测的时效区间
*/
type CainiaoDataLogisticsDeliveryAgingPredictRequest struct {
    model.Params
    // 发货城市
    _sendCityName   string
    // 发货区
    _sendCountyName   string
    // 发货详细地址
    _sendAddr   string
    // 发货省
    _sendProvName   string
    // 收货市
    _recCityName   string
    // 收货详细地址
    _recAddr   string
    // 收货区
    _recCountyName   string
    // 收货省
    _recProvName   string
    // 收货街道
    _recTownName   string
}

// 初始化CainiaoDataLogisticsDeliveryAgingPredictRequest对象
func NewCainiaoDataLogisticsDeliveryAgingPredictRequest() *CainiaoDataLogisticsDeliveryAgingPredictRequest{
    return &CainiaoDataLogisticsDeliveryAgingPredictRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoDataLogisticsDeliveryAgingPredictRequest) GetApiMethodName() string {
    return "cainiao.data.logistics.delivery.aging.predict"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoDataLogisticsDeliveryAgingPredictRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SendCityName Setter
// 发货城市
func (r *CainiaoDataLogisticsDeliveryAgingPredictRequest) SetSendCityName(_sendCityName string) error {
    r._sendCityName = _sendCityName
    r.Set("send_city_name", _sendCityName)
    return nil
}

// SendCityName Getter
func (r CainiaoDataLogisticsDeliveryAgingPredictRequest) GetSendCityName() string {
    return r._sendCityName
}
// SendCountyName Setter
// 发货区
func (r *CainiaoDataLogisticsDeliveryAgingPredictRequest) SetSendCountyName(_sendCountyName string) error {
    r._sendCountyName = _sendCountyName
    r.Set("send_county_name", _sendCountyName)
    return nil
}

// SendCountyName Getter
func (r CainiaoDataLogisticsDeliveryAgingPredictRequest) GetSendCountyName() string {
    return r._sendCountyName
}
// SendAddr Setter
// 发货详细地址
func (r *CainiaoDataLogisticsDeliveryAgingPredictRequest) SetSendAddr(_sendAddr string) error {
    r._sendAddr = _sendAddr
    r.Set("send_addr", _sendAddr)
    return nil
}

// SendAddr Getter
func (r CainiaoDataLogisticsDeliveryAgingPredictRequest) GetSendAddr() string {
    return r._sendAddr
}
// SendProvName Setter
// 发货省
func (r *CainiaoDataLogisticsDeliveryAgingPredictRequest) SetSendProvName(_sendProvName string) error {
    r._sendProvName = _sendProvName
    r.Set("send_prov_name", _sendProvName)
    return nil
}

// SendProvName Getter
func (r CainiaoDataLogisticsDeliveryAgingPredictRequest) GetSendProvName() string {
    return r._sendProvName
}
// RecCityName Setter
// 收货市
func (r *CainiaoDataLogisticsDeliveryAgingPredictRequest) SetRecCityName(_recCityName string) error {
    r._recCityName = _recCityName
    r.Set("rec_city_name", _recCityName)
    return nil
}

// RecCityName Getter
func (r CainiaoDataLogisticsDeliveryAgingPredictRequest) GetRecCityName() string {
    return r._recCityName
}
// RecAddr Setter
// 收货详细地址
func (r *CainiaoDataLogisticsDeliveryAgingPredictRequest) SetRecAddr(_recAddr string) error {
    r._recAddr = _recAddr
    r.Set("rec_addr", _recAddr)
    return nil
}

// RecAddr Getter
func (r CainiaoDataLogisticsDeliveryAgingPredictRequest) GetRecAddr() string {
    return r._recAddr
}
// RecCountyName Setter
// 收货区
func (r *CainiaoDataLogisticsDeliveryAgingPredictRequest) SetRecCountyName(_recCountyName string) error {
    r._recCountyName = _recCountyName
    r.Set("rec_county_name", _recCountyName)
    return nil
}

// RecCountyName Getter
func (r CainiaoDataLogisticsDeliveryAgingPredictRequest) GetRecCountyName() string {
    return r._recCountyName
}
// RecProvName Setter
// 收货省
func (r *CainiaoDataLogisticsDeliveryAgingPredictRequest) SetRecProvName(_recProvName string) error {
    r._recProvName = _recProvName
    r.Set("rec_prov_name", _recProvName)
    return nil
}

// RecProvName Getter
func (r CainiaoDataLogisticsDeliveryAgingPredictRequest) GetRecProvName() string {
    return r._recProvName
}
// RecTownName Setter
// 收货街道
func (r *CainiaoDataLogisticsDeliveryAgingPredictRequest) SetRecTownName(_recTownName string) error {
    r._recTownName = _recTownName
    r.Set("rec_town_name", _recTownName)
    return nil
}

// RecTownName Getter
func (r CainiaoDataLogisticsDeliveryAgingPredictRequest) GetRecTownName() string {
    return r._recTownName
}
