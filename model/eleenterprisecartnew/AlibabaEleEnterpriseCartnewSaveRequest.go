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
    phone   string
    // 用户所在纬度
    latitude   string
    // [[{\"id\": 1526467414,\"new_specs\": [{\"name\": \"规格\",\"value\": \"那么大鲜柠特饮(雪碧) 660ml\"}],\"attrs\": [{\"name\": \"可选小食\",\"value\": \"金黄脆薯格\"}],\"quantity\": 2}]]
    food   string
    // 用户所在经度
    longitude   string
    // 餐厅id
    erestaurantId   string
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
func (r *AlibabaEleEnterpriseCartnewSaveRequest) SetPhone(phone string) error {
    r.phone = phone
    r.Set("phone", phone)
    return nil
}

// Phone Getter
func (r AlibabaEleEnterpriseCartnewSaveRequest) GetPhone() string {
    return r.phone
}
// Latitude Setter
// 用户所在纬度
func (r *AlibabaEleEnterpriseCartnewSaveRequest) SetLatitude(latitude string) error {
    r.latitude = latitude
    r.Set("latitude", latitude)
    return nil
}

// Latitude Getter
func (r AlibabaEleEnterpriseCartnewSaveRequest) GetLatitude() string {
    return r.latitude
}
// Food Setter
// [[{\"id\": 1526467414,\"new_specs\": [{\"name\": \"规格\",\"value\": \"那么大鲜柠特饮(雪碧) 660ml\"}],\"attrs\": [{\"name\": \"可选小食\",\"value\": \"金黄脆薯格\"}],\"quantity\": 2}]]
func (r *AlibabaEleEnterpriseCartnewSaveRequest) SetFood(food string) error {
    r.food = food
    r.Set("food", food)
    return nil
}

// Food Getter
func (r AlibabaEleEnterpriseCartnewSaveRequest) GetFood() string {
    return r.food
}
// Longitude Setter
// 用户所在经度
func (r *AlibabaEleEnterpriseCartnewSaveRequest) SetLongitude(longitude string) error {
    r.longitude = longitude
    r.Set("longitude", longitude)
    return nil
}

// Longitude Getter
func (r AlibabaEleEnterpriseCartnewSaveRequest) GetLongitude() string {
    return r.longitude
}
// ErestaurantId Setter
// 餐厅id
func (r *AlibabaEleEnterpriseCartnewSaveRequest) SetErestaurantId(erestaurantId string) error {
    r.erestaurantId = erestaurantId
    r.Set("erestaurant_id", erestaurantId)
    return nil
}

// ErestaurantId Getter
func (r AlibabaEleEnterpriseCartnewSaveRequest) GetErestaurantId() string {
    return r.erestaurantId
}
