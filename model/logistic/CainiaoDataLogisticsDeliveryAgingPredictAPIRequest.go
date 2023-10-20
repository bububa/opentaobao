package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoDataLogisticsDeliveryAgingPredictAPIRequest 配送物流时效预测 API请求
// cainiao.data.logistics.delivery.aging.predict
//
// 时效和服务预期是商家发货时比较关注的信息，也是选择快递公司的一个重要参考（除去长期合作合同因素）。所以，在商家发货时给商家提供线路时效预估能帮助商家选择更好的快递公司，且对消费者来说也能整体提升体验。
//
// 日常，展示具体的预测时效数值
//
// 大促期间，展示预测的时效区间
type CainiaoDataLogisticsDeliveryAgingPredictAPIRequest struct {
	model.Params
	// 发货城市
	_sendCityName string
	// 发货区
	_sendCountyName string
	// 发货详细地址
	_sendAddr string
	// 发货省
	_sendProvName string
	// 收货市
	_recCityName string
	// 收货详细地址
	_recAddr string
	// 收货区
	_recCountyName string
	// 收货省
	_recProvName string
	// 收货街道
	_recTownName string
}

// NewCainiaoDataLogisticsDeliveryAgingPredictRequest 初始化CainiaoDataLogisticsDeliveryAgingPredictAPIRequest对象
func NewCainiaoDataLogisticsDeliveryAgingPredictRequest() *CainiaoDataLogisticsDeliveryAgingPredictAPIRequest {
	return &CainiaoDataLogisticsDeliveryAgingPredictAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoDataLogisticsDeliveryAgingPredictAPIRequest) GetApiMethodName() string {
	return "cainiao.data.logistics.delivery.aging.predict"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoDataLogisticsDeliveryAgingPredictAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoDataLogisticsDeliveryAgingPredictAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSendCityName is SendCityName Setter
// 发货城市
func (r *CainiaoDataLogisticsDeliveryAgingPredictAPIRequest) SetSendCityName(_sendCityName string) error {
	r._sendCityName = _sendCityName
	r.Set("send_city_name", _sendCityName)
	return nil
}

// GetSendCityName SendCityName Getter
func (r CainiaoDataLogisticsDeliveryAgingPredictAPIRequest) GetSendCityName() string {
	return r._sendCityName
}

// SetSendCountyName is SendCountyName Setter
// 发货区
func (r *CainiaoDataLogisticsDeliveryAgingPredictAPIRequest) SetSendCountyName(_sendCountyName string) error {
	r._sendCountyName = _sendCountyName
	r.Set("send_county_name", _sendCountyName)
	return nil
}

// GetSendCountyName SendCountyName Getter
func (r CainiaoDataLogisticsDeliveryAgingPredictAPIRequest) GetSendCountyName() string {
	return r._sendCountyName
}

// SetSendAddr is SendAddr Setter
// 发货详细地址
func (r *CainiaoDataLogisticsDeliveryAgingPredictAPIRequest) SetSendAddr(_sendAddr string) error {
	r._sendAddr = _sendAddr
	r.Set("send_addr", _sendAddr)
	return nil
}

// GetSendAddr SendAddr Getter
func (r CainiaoDataLogisticsDeliveryAgingPredictAPIRequest) GetSendAddr() string {
	return r._sendAddr
}

// SetSendProvName is SendProvName Setter
// 发货省
func (r *CainiaoDataLogisticsDeliveryAgingPredictAPIRequest) SetSendProvName(_sendProvName string) error {
	r._sendProvName = _sendProvName
	r.Set("send_prov_name", _sendProvName)
	return nil
}

// GetSendProvName SendProvName Getter
func (r CainiaoDataLogisticsDeliveryAgingPredictAPIRequest) GetSendProvName() string {
	return r._sendProvName
}

// SetRecCityName is RecCityName Setter
// 收货市
func (r *CainiaoDataLogisticsDeliveryAgingPredictAPIRequest) SetRecCityName(_recCityName string) error {
	r._recCityName = _recCityName
	r.Set("rec_city_name", _recCityName)
	return nil
}

// GetRecCityName RecCityName Getter
func (r CainiaoDataLogisticsDeliveryAgingPredictAPIRequest) GetRecCityName() string {
	return r._recCityName
}

// SetRecAddr is RecAddr Setter
// 收货详细地址
func (r *CainiaoDataLogisticsDeliveryAgingPredictAPIRequest) SetRecAddr(_recAddr string) error {
	r._recAddr = _recAddr
	r.Set("rec_addr", _recAddr)
	return nil
}

// GetRecAddr RecAddr Getter
func (r CainiaoDataLogisticsDeliveryAgingPredictAPIRequest) GetRecAddr() string {
	return r._recAddr
}

// SetRecCountyName is RecCountyName Setter
// 收货区
func (r *CainiaoDataLogisticsDeliveryAgingPredictAPIRequest) SetRecCountyName(_recCountyName string) error {
	r._recCountyName = _recCountyName
	r.Set("rec_county_name", _recCountyName)
	return nil
}

// GetRecCountyName RecCountyName Getter
func (r CainiaoDataLogisticsDeliveryAgingPredictAPIRequest) GetRecCountyName() string {
	return r._recCountyName
}

// SetRecProvName is RecProvName Setter
// 收货省
func (r *CainiaoDataLogisticsDeliveryAgingPredictAPIRequest) SetRecProvName(_recProvName string) error {
	r._recProvName = _recProvName
	r.Set("rec_prov_name", _recProvName)
	return nil
}

// GetRecProvName RecProvName Getter
func (r CainiaoDataLogisticsDeliveryAgingPredictAPIRequest) GetRecProvName() string {
	return r._recProvName
}

// SetRecTownName is RecTownName Setter
// 收货街道
func (r *CainiaoDataLogisticsDeliveryAgingPredictAPIRequest) SetRecTownName(_recTownName string) error {
	r._recTownName = _recTownName
	r.Set("rec_town_name", _recTownName)
	return nil
}

// GetRecTownName RecTownName Getter
func (r CainiaoDataLogisticsDeliveryAgingPredictAPIRequest) GetRecTownName() string {
	return r._recTownName
}
