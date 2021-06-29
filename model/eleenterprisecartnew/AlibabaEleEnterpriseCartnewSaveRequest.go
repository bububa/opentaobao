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
type AlibabaEleEnterpriseCartnewSaveRequest struct {
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

// 初始化AlibabaEleEnterpriseCartnewSaveRequest对象
func NewAlibabaEleEnterpriseCartnewSaveRequest() *AlibabaEleEnterpriseCartnewSaveRequest{
    return &AlibabaEleEnterpriseCartnewSaveRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEleEnterpriseCartnewSaveRequest) GetApiMethodName() string {
    return "alibaba.ele.enterprise.cartnew.save"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEleEnterpriseCartnewSaveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Phone Setter
// 用户11位手机号
func (r *AlibabaEleEnterpriseCartnewSaveRequest) SetPhone(_phone string) error {
    r._phone = _phone
    r.Set("phone", _phone)
    return nil
}

// Phone Getter
func (r AlibabaEleEnterpriseCartnewSaveRequest) GetPhone() string {
    return r._phone
}
// Latitude Setter
// 用户所在纬度
func (r *AlibabaEleEnterpriseCartnewSaveRequest) SetLatitude(_latitude string) error {
    r._latitude = _latitude
    r.Set("latitude", _latitude)
    return nil
}

// Latitude Getter
func (r AlibabaEleEnterpriseCartnewSaveRequest) GetLatitude() string {
    return r._latitude
}
// Food Setter
// [[{\"id\": 1526467414,\"new_specs\": [{\"name\": \"规格\",\"value\": \"那么大鲜柠特饮(雪碧) 660ml\"}],\"attrs\": [{\"name\": \"可选小食\",\"value\": \"金黄脆薯格\"}],\"quantity\": 2}]]
func (r *AlibabaEleEnterpriseCartnewSaveRequest) SetFood(_food string) error {
    r._food = _food
    r.Set("food", _food)
    return nil
}

// Food Getter
func (r AlibabaEleEnterpriseCartnewSaveRequest) GetFood() string {
    return r._food
}
// Longitude Setter
// 用户所在经度
func (r *AlibabaEleEnterpriseCartnewSaveRequest) SetLongitude(_longitude string) error {
    r._longitude = _longitude
    r.Set("longitude", _longitude)
    return nil
}

// Longitude Getter
func (r AlibabaEleEnterpriseCartnewSaveRequest) GetLongitude() string {
    return r._longitude
}
// ErestaurantId Setter
// 餐厅id
func (r *AlibabaEleEnterpriseCartnewSaveRequest) SetErestaurantId(_erestaurantId string) error {
    r._erestaurantId = _erestaurantId
    r.Set("erestaurant_id", _erestaurantId)
    return nil
}

// ErestaurantId Getter
func (r AlibabaEleEnterpriseCartnewSaveRequest) GetErestaurantId() string {
    return r._erestaurantId
}
