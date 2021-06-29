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
    sendCityName   string
    // 发货区
    sendCountyName   string
    // 发货详细地址
    sendAddr   string
    // 发货省
    sendProvName   string
    // 收货市
    recCityName   string
    // 收货详细地址
    recAddr   string
    // 收货区
    recCountyName   string
    // 收货省
    recProvName   string
    // 收货街道
    recTownName   string
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
func (r *CainiaoDataLogisticsDeliveryAgingPredictRequest) SetSendCityName(sendCityName string) error {
    r.sendCityName = sendCityName
    r.Set("send_city_name", sendCityName)
    return nil
}

// SendCityName Getter
func (r CainiaoDataLogisticsDeliveryAgingPredictRequest) GetSendCityName() string {
    return r.sendCityName
}
// SendCountyName Setter
// 发货区
func (r *CainiaoDataLogisticsDeliveryAgingPredictRequest) SetSendCountyName(sendCountyName string) error {
    r.sendCountyName = sendCountyName
    r.Set("send_county_name", sendCountyName)
    return nil
}

// SendCountyName Getter
func (r CainiaoDataLogisticsDeliveryAgingPredictRequest) GetSendCountyName() string {
    return r.sendCountyName
}
// SendAddr Setter
// 发货详细地址
func (r *CainiaoDataLogisticsDeliveryAgingPredictRequest) SetSendAddr(sendAddr string) error {
    r.sendAddr = sendAddr
    r.Set("send_addr", sendAddr)
    return nil
}

// SendAddr Getter
func (r CainiaoDataLogisticsDeliveryAgingPredictRequest) GetSendAddr() string {
    return r.sendAddr
}
// SendProvName Setter
// 发货省
func (r *CainiaoDataLogisticsDeliveryAgingPredictRequest) SetSendProvName(sendProvName string) error {
    r.sendProvName = sendProvName
    r.Set("send_prov_name", sendProvName)
    return nil
}

// SendProvName Getter
func (r CainiaoDataLogisticsDeliveryAgingPredictRequest) GetSendProvName() string {
    return r.sendProvName
}
// RecCityName Setter
// 收货市
func (r *CainiaoDataLogisticsDeliveryAgingPredictRequest) SetRecCityName(recCityName string) error {
    r.recCityName = recCityName
    r.Set("rec_city_name", recCityName)
    return nil
}

// RecCityName Getter
func (r CainiaoDataLogisticsDeliveryAgingPredictRequest) GetRecCityName() string {
    return r.recCityName
}
// RecAddr Setter
// 收货详细地址
func (r *CainiaoDataLogisticsDeliveryAgingPredictRequest) SetRecAddr(recAddr string) error {
    r.recAddr = recAddr
    r.Set("rec_addr", recAddr)
    return nil
}

// RecAddr Getter
func (r CainiaoDataLogisticsDeliveryAgingPredictRequest) GetRecAddr() string {
    return r.recAddr
}
// RecCountyName Setter
// 收货区
func (r *CainiaoDataLogisticsDeliveryAgingPredictRequest) SetRecCountyName(recCountyName string) error {
    r.recCountyName = recCountyName
    r.Set("rec_county_name", recCountyName)
    return nil
}

// RecCountyName Getter
func (r CainiaoDataLogisticsDeliveryAgingPredictRequest) GetRecCountyName() string {
    return r.recCountyName
}
// RecProvName Setter
// 收货省
func (r *CainiaoDataLogisticsDeliveryAgingPredictRequest) SetRecProvName(recProvName string) error {
    r.recProvName = recProvName
    r.Set("rec_prov_name", recProvName)
    return nil
}

// RecProvName Getter
func (r CainiaoDataLogisticsDeliveryAgingPredictRequest) GetRecProvName() string {
    return r.recProvName
}
// RecTownName Setter
// 收货街道
func (r *CainiaoDataLogisticsDeliveryAgingPredictRequest) SetRecTownName(recTownName string) error {
    r.recTownName = recTownName
    r.Set("rec_town_name", recTownName)
    return nil
}

// RecTownName Getter
func (r CainiaoDataLogisticsDeliveryAgingPredictRequest) GetRecTownName() string {
    return r.recTownName
}
