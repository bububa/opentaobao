package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoDataLogisticsCpDeliveryAgingPredictAPIRequest CP配送物流时效预测 API请求
// cainiao.data.logistics.cp.delivery.aging.predict
//
// 时效和服务预期是商家发货时比较关注的信息，也是选择快递公司的一个重要参考（除去长期合作合同因素）。所以，在商家发货时给商家提供线路时效预估能帮助商家选择更好的快递公司，且对消费者来说也能整体提升体验。
//
// 日常时效是数值字符串
// 大促时效是数值区间字符串
// 方式1:
// 输入：发货省、市、区、详细地址，收货省、市、区、街道、详细地址，快递公司ID
// 输出：预估时效（小时数）
type CainiaoDataLogisticsCpDeliveryAgingPredictAPIRequest struct {
	model.Params
	// 发货城市
	_sendCityName string
	// 发货区
	_sendCountyName string
	// 自己输入的详细发货地址
	_sendAddr string
	// 发货省
	_sendProvName string
	// 收货城市
	_recCityName string
	// 自己输入的详细收货地址
	_recAddr string
	// 收货区
	_recCountyName string
	// 收货省
	_recProvName string
	// 第四级，一般是收货街道
	_recTownName string
	// 物流公司id
	_cpId string
}

// NewCainiaoDataLogisticsCpDeliveryAgingPredictRequest 初始化CainiaoDataLogisticsCpDeliveryAgingPredictAPIRequest对象
func NewCainiaoDataLogisticsCpDeliveryAgingPredictRequest() *CainiaoDataLogisticsCpDeliveryAgingPredictAPIRequest {
	return &CainiaoDataLogisticsCpDeliveryAgingPredictAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoDataLogisticsCpDeliveryAgingPredictAPIRequest) GetApiMethodName() string {
	return "cainiao.data.logistics.cp.delivery.aging.predict"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoDataLogisticsCpDeliveryAgingPredictAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoDataLogisticsCpDeliveryAgingPredictAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSendCityName is SendCityName Setter
// 发货城市
func (r *CainiaoDataLogisticsCpDeliveryAgingPredictAPIRequest) SetSendCityName(_sendCityName string) error {
	r._sendCityName = _sendCityName
	r.Set("send_city_name", _sendCityName)
	return nil
}

// GetSendCityName SendCityName Getter
func (r CainiaoDataLogisticsCpDeliveryAgingPredictAPIRequest) GetSendCityName() string {
	return r._sendCityName
}

// SetSendCountyName is SendCountyName Setter
// 发货区
func (r *CainiaoDataLogisticsCpDeliveryAgingPredictAPIRequest) SetSendCountyName(_sendCountyName string) error {
	r._sendCountyName = _sendCountyName
	r.Set("send_county_name", _sendCountyName)
	return nil
}

// GetSendCountyName SendCountyName Getter
func (r CainiaoDataLogisticsCpDeliveryAgingPredictAPIRequest) GetSendCountyName() string {
	return r._sendCountyName
}

// SetSendAddr is SendAddr Setter
// 自己输入的详细发货地址
func (r *CainiaoDataLogisticsCpDeliveryAgingPredictAPIRequest) SetSendAddr(_sendAddr string) error {
	r._sendAddr = _sendAddr
	r.Set("send_addr", _sendAddr)
	return nil
}

// GetSendAddr SendAddr Getter
func (r CainiaoDataLogisticsCpDeliveryAgingPredictAPIRequest) GetSendAddr() string {
	return r._sendAddr
}

// SetSendProvName is SendProvName Setter
// 发货省
func (r *CainiaoDataLogisticsCpDeliveryAgingPredictAPIRequest) SetSendProvName(_sendProvName string) error {
	r._sendProvName = _sendProvName
	r.Set("send_prov_name", _sendProvName)
	return nil
}

// GetSendProvName SendProvName Getter
func (r CainiaoDataLogisticsCpDeliveryAgingPredictAPIRequest) GetSendProvName() string {
	return r._sendProvName
}

// SetRecCityName is RecCityName Setter
// 收货城市
func (r *CainiaoDataLogisticsCpDeliveryAgingPredictAPIRequest) SetRecCityName(_recCityName string) error {
	r._recCityName = _recCityName
	r.Set("rec_city_name", _recCityName)
	return nil
}

// GetRecCityName RecCityName Getter
func (r CainiaoDataLogisticsCpDeliveryAgingPredictAPIRequest) GetRecCityName() string {
	return r._recCityName
}

// SetRecAddr is RecAddr Setter
// 自己输入的详细收货地址
func (r *CainiaoDataLogisticsCpDeliveryAgingPredictAPIRequest) SetRecAddr(_recAddr string) error {
	r._recAddr = _recAddr
	r.Set("rec_addr", _recAddr)
	return nil
}

// GetRecAddr RecAddr Getter
func (r CainiaoDataLogisticsCpDeliveryAgingPredictAPIRequest) GetRecAddr() string {
	return r._recAddr
}

// SetRecCountyName is RecCountyName Setter
// 收货区
func (r *CainiaoDataLogisticsCpDeliveryAgingPredictAPIRequest) SetRecCountyName(_recCountyName string) error {
	r._recCountyName = _recCountyName
	r.Set("rec_county_name", _recCountyName)
	return nil
}

// GetRecCountyName RecCountyName Getter
func (r CainiaoDataLogisticsCpDeliveryAgingPredictAPIRequest) GetRecCountyName() string {
	return r._recCountyName
}

// SetRecProvName is RecProvName Setter
// 收货省
func (r *CainiaoDataLogisticsCpDeliveryAgingPredictAPIRequest) SetRecProvName(_recProvName string) error {
	r._recProvName = _recProvName
	r.Set("rec_prov_name", _recProvName)
	return nil
}

// GetRecProvName RecProvName Getter
func (r CainiaoDataLogisticsCpDeliveryAgingPredictAPIRequest) GetRecProvName() string {
	return r._recProvName
}

// SetRecTownName is RecTownName Setter
// 第四级，一般是收货街道
func (r *CainiaoDataLogisticsCpDeliveryAgingPredictAPIRequest) SetRecTownName(_recTownName string) error {
	r._recTownName = _recTownName
	r.Set("rec_town_name", _recTownName)
	return nil
}

// GetRecTownName RecTownName Getter
func (r CainiaoDataLogisticsCpDeliveryAgingPredictAPIRequest) GetRecTownName() string {
	return r._recTownName
}

// SetCpId is CpId Setter
// 物流公司id
func (r *CainiaoDataLogisticsCpDeliveryAgingPredictAPIRequest) SetCpId(_cpId string) error {
	r._cpId = _cpId
	r.Set("cp_id", _cpId)
	return nil
}

// GetCpId CpId Getter
func (r CainiaoDataLogisticsCpDeliveryAgingPredictAPIRequest) GetCpId() string {
	return r._cpId
}
