package eleenterprisecartnew

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleEnterpriseCartnewQueryAPIRequest 新版购物车查询 API请求
// alibaba.ele.enterprise.cartnew.query
//
// 新版购物车查询
type AlibabaEleEnterpriseCartnewQueryAPIRequest struct {
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

// NewAlibabaEleEnterpriseCartnewQueryRequest 初始化AlibabaEleEnterpriseCartnewQueryAPIRequest对象
func NewAlibabaEleEnterpriseCartnewQueryRequest() *AlibabaEleEnterpriseCartnewQueryAPIRequest {
	return &AlibabaEleEnterpriseCartnewQueryAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaEleEnterpriseCartnewQueryAPIRequest) Reset() {
	r._phone = ""
	r._latitude = ""
	r._longitude = ""
	r._erestaurantId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEleEnterpriseCartnewQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.enterprise.cartnew.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEleEnterpriseCartnewQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEleEnterpriseCartnewQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPhone is Phone Setter
// 1212
func (r *AlibabaEleEnterpriseCartnewQueryAPIRequest) SetPhone(_phone string) error {
	r._phone = _phone
	r.Set("phone", _phone)
	return nil
}

// GetPhone Phone Getter
func (r AlibabaEleEnterpriseCartnewQueryAPIRequest) GetPhone() string {
	return r._phone
}

// SetLatitude is Latitude Setter
// 1212
func (r *AlibabaEleEnterpriseCartnewQueryAPIRequest) SetLatitude(_latitude string) error {
	r._latitude = _latitude
	r.Set("latitude", _latitude)
	return nil
}

// GetLatitude Latitude Getter
func (r AlibabaEleEnterpriseCartnewQueryAPIRequest) GetLatitude() string {
	return r._latitude
}

// SetLongitude is Longitude Setter
// 1212
func (r *AlibabaEleEnterpriseCartnewQueryAPIRequest) SetLongitude(_longitude string) error {
	r._longitude = _longitude
	r.Set("longitude", _longitude)
	return nil
}

// GetLongitude Longitude Getter
func (r AlibabaEleEnterpriseCartnewQueryAPIRequest) GetLongitude() string {
	return r._longitude
}

// SetErestaurantId is ErestaurantId Setter
// 餐厅id
func (r *AlibabaEleEnterpriseCartnewQueryAPIRequest) SetErestaurantId(_erestaurantId string) error {
	r._erestaurantId = _erestaurantId
	r.Set("erestaurant_id", _erestaurantId)
	return nil
}

// GetErestaurantId ErestaurantId Getter
func (r AlibabaEleEnterpriseCartnewQueryAPIRequest) GetErestaurantId() string {
	return r._erestaurantId
}

var poolAlibabaEleEnterpriseCartnewQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaEleEnterpriseCartnewQueryRequest()
	},
}

// GetAlibabaEleEnterpriseCartnewQueryRequest 从 sync.Pool 获取 AlibabaEleEnterpriseCartnewQueryAPIRequest
func GetAlibabaEleEnterpriseCartnewQueryAPIRequest() *AlibabaEleEnterpriseCartnewQueryAPIRequest {
	return poolAlibabaEleEnterpriseCartnewQueryAPIRequest.Get().(*AlibabaEleEnterpriseCartnewQueryAPIRequest)
}

// ReleaseAlibabaEleEnterpriseCartnewQueryAPIRequest 将 AlibabaEleEnterpriseCartnewQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaEleEnterpriseCartnewQueryAPIRequest(v *AlibabaEleEnterpriseCartnewQueryAPIRequest) {
	v.Reset()
	poolAlibabaEleEnterpriseCartnewQueryAPIRequest.Put(v)
}
