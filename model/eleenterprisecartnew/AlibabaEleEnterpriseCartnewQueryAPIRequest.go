package eleenterprisecartnew

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeleenterprisecartnewqueryAPIRequest 新版购物车查询 API请求
// alibaba.ele.enterprise.cartnew.query
//
// 新版购物车查询
type AlibabaeleenterprisecartnewqueryAPIRequest struct {
	model.Params
	// 1212
	_phone string
	// 1212
	_latitude string
	// 1212
	_longitude string
	// 餐厅id
	_erestaurantId string
}

// NewAlibabaeleenterprisecartnewqueryRequest 初始化AlibabaeleenterprisecartnewqueryAPIRequest对象
func NewAlibabaeleenterprisecartnewqueryRequest() *AlibabaeleenterprisecartnewqueryAPIRequest {
	return &AlibabaeleenterprisecartnewqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaeleenterprisecartnewqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.enterprise.cartnew.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaeleenterprisecartnewqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaeleenterprisecartnewqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPhone is Phone Setter
// 1212
func (r *AlibabaeleenterprisecartnewqueryAPIRequest) SetPhone(_phone string) error {
	r._phone = _phone
	r.Set("phone", _phone)
	return nil
}

// GetPhone Phone Getter
func (r AlibabaeleenterprisecartnewqueryAPIRequest) GetPhone() string {
	return r._phone
}

// SetLatitude is Latitude Setter
// 1212
func (r *AlibabaeleenterprisecartnewqueryAPIRequest) SetLatitude(_latitude string) error {
	r._latitude = _latitude
	r.Set("latitude", _latitude)
	return nil
}

// GetLatitude Latitude Getter
func (r AlibabaeleenterprisecartnewqueryAPIRequest) GetLatitude() string {
	return r._latitude
}

// SetLongitude is Longitude Setter
// 1212
func (r *AlibabaeleenterprisecartnewqueryAPIRequest) SetLongitude(_longitude string) error {
	r._longitude = _longitude
	r.Set("longitude", _longitude)
	return nil
}

// GetLongitude Longitude Getter
func (r AlibabaeleenterprisecartnewqueryAPIRequest) GetLongitude() string {
	return r._longitude
}

// SetErestaurantId is ErestaurantId Setter
// 餐厅id
func (r *AlibabaeleenterprisecartnewqueryAPIRequest) SetErestaurantId(_erestaurantId string) error {
	r._erestaurantId = _erestaurantId
	r.Set("erestaurant_id", _erestaurantId)
	return nil
}

// GetErestaurantId ErestaurantId Getter
func (r AlibabaeleenterprisecartnewqueryAPIRequest) GetErestaurantId() string {
	return r._erestaurantId
}
