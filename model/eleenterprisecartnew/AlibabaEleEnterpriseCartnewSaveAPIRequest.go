package eleenterprisecartnew

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeleenterprisecartnewsaveAPIRequest 新版创建购物车 API请求
// alibaba.ele.enterprise.cartnew.save
//
// 新版创建购物车
type AlibabaeleenterprisecartnewsaveAPIRequest struct {
	model.Params
	// 用户11位手机号
	_phone string
	// 用户所在纬度
	_latitude string
	// [[{\"id\": 1526467414,\"new_specs\": [{\"name\": \"规格\",\"value\": \"那么大鲜柠特饮(雪碧) 660ml\"}],\"attrs\": [{\"name\": \"可选小食\",\"value\": \"金黄脆薯格\"}],\"quantity\": 2}]]
	_food string
	// 用户所在经度
	_longitude string
	// 餐厅id
	_erestaurantId string
}

// NewAlibabaeleenterprisecartnewsaveRequest 初始化AlibabaeleenterprisecartnewsaveAPIRequest对象
func NewAlibabaeleenterprisecartnewsaveRequest() *AlibabaeleenterprisecartnewsaveAPIRequest {
	return &AlibabaeleenterprisecartnewsaveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaeleenterprisecartnewsaveAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.enterprise.cartnew.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaeleenterprisecartnewsaveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaeleenterprisecartnewsaveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPhone is Phone Setter
// 用户11位手机号
func (r *AlibabaeleenterprisecartnewsaveAPIRequest) SetPhone(_phone string) error {
	r._phone = _phone
	r.Set("phone", _phone)
	return nil
}

// GetPhone Phone Getter
func (r AlibabaeleenterprisecartnewsaveAPIRequest) GetPhone() string {
	return r._phone
}

// SetLatitude is Latitude Setter
// 用户所在纬度
func (r *AlibabaeleenterprisecartnewsaveAPIRequest) SetLatitude(_latitude string) error {
	r._latitude = _latitude
	r.Set("latitude", _latitude)
	return nil
}

// GetLatitude Latitude Getter
func (r AlibabaeleenterprisecartnewsaveAPIRequest) GetLatitude() string {
	return r._latitude
}

// SetFood is Food Setter
// [[{\&#34;id\&#34;: 1526467414,\&#34;new_specs\&#34;: [{\&#34;name\&#34;: \&#34;规格\&#34;,\&#34;value\&#34;: \&#34;那么大鲜柠特饮(雪碧) 660ml\&#34;}],\&#34;attrs\&#34;: [{\&#34;name\&#34;: \&#34;可选小食\&#34;,\&#34;value\&#34;: \&#34;金黄脆薯格\&#34;}],\&#34;quantity\&#34;: 2}]]
func (r *AlibabaeleenterprisecartnewsaveAPIRequest) SetFood(_food string) error {
	r._food = _food
	r.Set("food", _food)
	return nil
}

// GetFood Food Getter
func (r AlibabaeleenterprisecartnewsaveAPIRequest) GetFood() string {
	return r._food
}

// SetLongitude is Longitude Setter
// 用户所在经度
func (r *AlibabaeleenterprisecartnewsaveAPIRequest) SetLongitude(_longitude string) error {
	r._longitude = _longitude
	r.Set("longitude", _longitude)
	return nil
}

// GetLongitude Longitude Getter
func (r AlibabaeleenterprisecartnewsaveAPIRequest) GetLongitude() string {
	return r._longitude
}

// SetErestaurantId is ErestaurantId Setter
// 餐厅id
func (r *AlibabaeleenterprisecartnewsaveAPIRequest) SetErestaurantId(_erestaurantId string) error {
	r._erestaurantId = _erestaurantId
	r.Set("erestaurant_id", _erestaurantId)
	return nil
}

// GetErestaurantId ErestaurantId Getter
func (r AlibabaeleenterprisecartnewsaveAPIRequest) GetErestaurantId() string {
	return r._erestaurantId
}
