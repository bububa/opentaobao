package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
CP配送物流时效预测 API请求
cainiao.data.logistics.cp.delivery.aging.predict

时效和服务预期是商家发货时比较关注的信息，也是选择快递公司的一个重要参考（除去长期合作合同因素）。所以，在商家发货时给商家提供线路时效预估能帮助商家选择更好的快递公司，且对消费者来说也能整体提升体验。
 
日常时效是数值字符串
大促时效是数值区间字符串
方式1:
输入：发货省、市、区、详细地址，收货省、市、区、街道、详细地址，快递公司ID
输出：预估时效（小时数）
*/
type CainiaoDataLogisticsCpDeliveryAgingPredictRequest struct {
    model.Params
    // 发货城市
    _sendCityName   string
    // 发货区
    _sendCountyName   string
    // 自己输入的详细发货地址
    _sendAddr   string
    // 发货省
    _sendProvName   string
    // 收货城市
    _recCityName   string
    // 自己输入的详细收货地址
    _recAddr   string
    // 收货区
    _recCountyName   string
    // 收货省
    _recProvName   string
    // 第四级，一般是收货街道
    _recTownName   string
    // 物流公司id
    _cpId   string
}

// 初始化CainiaoDataLogisticsCpDeliveryAgingPredictRequest对象
func NewCainiaoDataLogisticsCpDeliveryAgingPredictRequest() *CainiaoDataLogisticsCpDeliveryAgingPredictRequest{
    return &CainiaoDataLogisticsCpDeliveryAgingPredictRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoDataLogisticsCpDeliveryAgingPredictRequest) GetApiMethodName() string {
    return "cainiao.data.logistics.cp.delivery.aging.predict"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoDataLogisticsCpDeliveryAgingPredictRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SendCityName Setter
// 发货城市
func (r *CainiaoDataLogisticsCpDeliveryAgingPredictRequest) SetSendCityName(_sendCityName string) error {
    r._sendCityName = _sendCityName
    r.Set("send_city_name", _sendCityName)
    return nil
}

// SendCityName Getter
func (r CainiaoDataLogisticsCpDeliveryAgingPredictRequest) GetSendCityName() string {
    return r._sendCityName
}
// SendCountyName Setter
// 发货区
func (r *CainiaoDataLogisticsCpDeliveryAgingPredictRequest) SetSendCountyName(_sendCountyName string) error {
    r._sendCountyName = _sendCountyName
    r.Set("send_county_name", _sendCountyName)
    return nil
}

// SendCountyName Getter
func (r CainiaoDataLogisticsCpDeliveryAgingPredictRequest) GetSendCountyName() string {
    return r._sendCountyName
}
// SendAddr Setter
// 自己输入的详细发货地址
func (r *CainiaoDataLogisticsCpDeliveryAgingPredictRequest) SetSendAddr(_sendAddr string) error {
    r._sendAddr = _sendAddr
    r.Set("send_addr", _sendAddr)
    return nil
}

// SendAddr Getter
func (r CainiaoDataLogisticsCpDeliveryAgingPredictRequest) GetSendAddr() string {
    return r._sendAddr
}
// SendProvName Setter
// 发货省
func (r *CainiaoDataLogisticsCpDeliveryAgingPredictRequest) SetSendProvName(_sendProvName string) error {
    r._sendProvName = _sendProvName
    r.Set("send_prov_name", _sendProvName)
    return nil
}

// SendProvName Getter
func (r CainiaoDataLogisticsCpDeliveryAgingPredictRequest) GetSendProvName() string {
    return r._sendProvName
}
// RecCityName Setter
// 收货城市
func (r *CainiaoDataLogisticsCpDeliveryAgingPredictRequest) SetRecCityName(_recCityName string) error {
    r._recCityName = _recCityName
    r.Set("rec_city_name", _recCityName)
    return nil
}

// RecCityName Getter
func (r CainiaoDataLogisticsCpDeliveryAgingPredictRequest) GetRecCityName() string {
    return r._recCityName
}
// RecAddr Setter
// 自己输入的详细收货地址
func (r *CainiaoDataLogisticsCpDeliveryAgingPredictRequest) SetRecAddr(_recAddr string) error {
    r._recAddr = _recAddr
    r.Set("rec_addr", _recAddr)
    return nil
}

// RecAddr Getter
func (r CainiaoDataLogisticsCpDeliveryAgingPredictRequest) GetRecAddr() string {
    return r._recAddr
}
// RecCountyName Setter
// 收货区
func (r *CainiaoDataLogisticsCpDeliveryAgingPredictRequest) SetRecCountyName(_recCountyName string) error {
    r._recCountyName = _recCountyName
    r.Set("rec_county_name", _recCountyName)
    return nil
}

// RecCountyName Getter
func (r CainiaoDataLogisticsCpDeliveryAgingPredictRequest) GetRecCountyName() string {
    return r._recCountyName
}
// RecProvName Setter
// 收货省
func (r *CainiaoDataLogisticsCpDeliveryAgingPredictRequest) SetRecProvName(_recProvName string) error {
    r._recProvName = _recProvName
    r.Set("rec_prov_name", _recProvName)
    return nil
}

// RecProvName Getter
func (r CainiaoDataLogisticsCpDeliveryAgingPredictRequest) GetRecProvName() string {
    return r._recProvName
}
// RecTownName Setter
// 第四级，一般是收货街道
func (r *CainiaoDataLogisticsCpDeliveryAgingPredictRequest) SetRecTownName(_recTownName string) error {
    r._recTownName = _recTownName
    r.Set("rec_town_name", _recTownName)
    return nil
}

// RecTownName Getter
func (r CainiaoDataLogisticsCpDeliveryAgingPredictRequest) GetRecTownName() string {
    return r._recTownName
}
// CpId Setter
// 物流公司id
func (r *CainiaoDataLogisticsCpDeliveryAgingPredictRequest) SetCpId(_cpId string) error {
    r._cpId = _cpId
    r.Set("cp_id", _cpId)
    return nil
}

// CpId Getter
func (r CainiaoDataLogisticsCpDeliveryAgingPredictRequest) GetCpId() string {
    return r._cpId
}
