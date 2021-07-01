package eleenterprisecartnew

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新版创建购物车 API请求
alibaba.ele.enterprise.cartnew.save

新版创建购物车
*/
type AlibabaEleEnterpriseCartnewSaveAPIRequest struct {
    model.Params
    // 用户11位手机号
    _phone   string
    // 用户所在纬度
    _latitude   string
    // [[{\"id\": 1526467414,\"new_specs\": [{\"name\": \"规格\",\"value\": \"那么大鲜柠特饮(雪碧) 660ml\"}],\"attrs\": [{\"name\": \"可选小食\",\"value\": \"金黄脆薯格\"}],\"quantity\": 2}]]
    _food   string
    // 用户所在经度
    _longitude   string
    // 餐厅id
    _erestaurantId   string
}

// 初始化AlibabaEleEnterpriseCartnewSaveAPIRequest对象
func NewAlibabaEleEnterpriseCartnewSaveRequest() *AlibabaEleEnterpriseCartnewSaveAPIRequest{
    return &AlibabaEleEnterpriseCartnewSaveAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEleEnterpriseCartnewSaveAPIRequest) GetApiMethodName() string {
    return "alibaba.ele.enterprise.cartnew.save"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEleEnterpriseCartnewSaveAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Phone Setter
// 用户11位手机号
func (r *AlibabaEleEnterpriseCartnewSaveAPIRequest) SetPhone(_phone string) error {
    r._phone = _phone
    r.Set("phone", _phone)
    return nil
}

// Phone Getter
func (r AlibabaEleEnterpriseCartnewSaveAPIRequest) GetPhone() string {
    return r._phone
}
// Latitude Setter
// 用户所在纬度
func (r *AlibabaEleEnterpriseCartnewSaveAPIRequest) SetLatitude(_latitude string) error {
    r._latitude = _latitude
    r.Set("latitude", _latitude)
    return nil
}

// Latitude Getter
func (r AlibabaEleEnterpriseCartnewSaveAPIRequest) GetLatitude() string {
    return r._latitude
}
// Food Setter
// [[{\"id\": 1526467414,\"new_specs\": [{\"name\": \"规格\",\"value\": \"那么大鲜柠特饮(雪碧) 660ml\"}],\"attrs\": [{\"name\": \"可选小食\",\"value\": \"金黄脆薯格\"}],\"quantity\": 2}]]
func (r *AlibabaEleEnterpriseCartnewSaveAPIRequest) SetFood(_food string) error {
    r._food = _food
    r.Set("food", _food)
    return nil
}

// Food Getter
func (r AlibabaEleEnterpriseCartnewSaveAPIRequest) GetFood() string {
    return r._food
}
// Longitude Setter
// 用户所在经度
func (r *AlibabaEleEnterpriseCartnewSaveAPIRequest) SetLongitude(_longitude string) error {
    r._longitude = _longitude
    r.Set("longitude", _longitude)
    return nil
}

// Longitude Getter
func (r AlibabaEleEnterpriseCartnewSaveAPIRequest) GetLongitude() string {
    return r._longitude
}
// ErestaurantId Setter
// 餐厅id
func (r *AlibabaEleEnterpriseCartnewSaveAPIRequest) SetErestaurantId(_erestaurantId string) error {
    r._erestaurantId = _erestaurantId
    r.Set("erestaurant_id", _erestaurantId)
    return nil
}

// ErestaurantId Getter
func (r AlibabaEleEnterpriseCartnewSaveAPIRequest) GetErestaurantId() string {
    return r._erestaurantId
}
