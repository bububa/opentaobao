package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
CP配送物流时效预测 APIRequest
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
    sendCityName   string 

    // 发货区
    sendCountyName   string 

    // 自己输入的详细发货地址
    sendAddr   string 

    // 发货省
    sendProvName   string 

    // 收货城市
    recCityName   string 

    // 自己输入的详细收货地址
    recAddr   string 

    // 收货区
    recCountyName   string 

    // 收货省
    recProvName   string 

    // 第四级，一般是收货街道
    recTownName   string 

    // 物流公司id
    cpId   string 

}

func NewCainiaoDataLogisticsCpDeliveryAgingPredictRequest() *CainiaoDataLogisticsCpDeliveryAgingPredictRequest{
    return &CainiaoDataLogisticsCpDeliveryAgingPredictRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoDataLogisticsCpDeliveryAgingPredictRequest) GetApiMethodName() string {
    return "cainiao.data.logistics.cp.delivery.aging.predict"
}

func (r CainiaoDataLogisticsCpDeliveryAgingPredictRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoDataLogisticsCpDeliveryAgingPredictRequest) SetSendCityName(sendCityName string) error {
    r.sendCityName = sendCityName
    r.Set("send_city_name", sendCityName)
    return nil
}

func (r CainiaoDataLogisticsCpDeliveryAgingPredictRequest) GetSendCityName() string {
    return r.sendCityName
}

func (r *CainiaoDataLogisticsCpDeliveryAgingPredictRequest) SetSendCountyName(sendCountyName string) error {
    r.sendCountyName = sendCountyName
    r.Set("send_county_name", sendCountyName)
    return nil
}

func (r CainiaoDataLogisticsCpDeliveryAgingPredictRequest) GetSendCountyName() string {
    return r.sendCountyName
}

func (r *CainiaoDataLogisticsCpDeliveryAgingPredictRequest) SetSendAddr(sendAddr string) error {
    r.sendAddr = sendAddr
    r.Set("send_addr", sendAddr)
    return nil
}

func (r CainiaoDataLogisticsCpDeliveryAgingPredictRequest) GetSendAddr() string {
    return r.sendAddr
}

func (r *CainiaoDataLogisticsCpDeliveryAgingPredictRequest) SetSendProvName(sendProvName string) error {
    r.sendProvName = sendProvName
    r.Set("send_prov_name", sendProvName)
    return nil
}

func (r CainiaoDataLogisticsCpDeliveryAgingPredictRequest) GetSendProvName() string {
    return r.sendProvName
}

func (r *CainiaoDataLogisticsCpDeliveryAgingPredictRequest) SetRecCityName(recCityName string) error {
    r.recCityName = recCityName
    r.Set("rec_city_name", recCityName)
    return nil
}

func (r CainiaoDataLogisticsCpDeliveryAgingPredictRequest) GetRecCityName() string {
    return r.recCityName
}

func (r *CainiaoDataLogisticsCpDeliveryAgingPredictRequest) SetRecAddr(recAddr string) error {
    r.recAddr = recAddr
    r.Set("rec_addr", recAddr)
    return nil
}

func (r CainiaoDataLogisticsCpDeliveryAgingPredictRequest) GetRecAddr() string {
    return r.recAddr
}

func (r *CainiaoDataLogisticsCpDeliveryAgingPredictRequest) SetRecCountyName(recCountyName string) error {
    r.recCountyName = recCountyName
    r.Set("rec_county_name", recCountyName)
    return nil
}

func (r CainiaoDataLogisticsCpDeliveryAgingPredictRequest) GetRecCountyName() string {
    return r.recCountyName
}

func (r *CainiaoDataLogisticsCpDeliveryAgingPredictRequest) SetRecProvName(recProvName string) error {
    r.recProvName = recProvName
    r.Set("rec_prov_name", recProvName)
    return nil
}

func (r CainiaoDataLogisticsCpDeliveryAgingPredictRequest) GetRecProvName() string {
    return r.recProvName
}

func (r *CainiaoDataLogisticsCpDeliveryAgingPredictRequest) SetRecTownName(recTownName string) error {
    r.recTownName = recTownName
    r.Set("rec_town_name", recTownName)
    return nil
}

func (r CainiaoDataLogisticsCpDeliveryAgingPredictRequest) GetRecTownName() string {
    return r.recTownName
}

func (r *CainiaoDataLogisticsCpDeliveryAgingPredictRequest) SetCpId(cpId string) error {
    r.cpId = cpId
    r.Set("cp_id", cpId)
    return nil
}

func (r CainiaoDataLogisticsCpDeliveryAgingPredictRequest) GetCpId() string {
    return r.cpId
}

