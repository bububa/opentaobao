package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
系外拉新代理商增加客户经理接口 API请求
alibaba.aliqin.offline.customer.add

阿里通信这边维护了代理商和其对应的客户经理的关系，用于业务处理，开放该接口用于代理商将他们系统下的客户经理信息同步给我们
*/
type AlibabaAliqinOfflineCustomerAddRequest struct {
    model.Params
    // 代理商id
    distributeId   string
    // 网点id，如果存在填写，不存在的话，填0即可；注意：如果填写了这个字段，后面的pob_name等会失效；如果为0，下面的网点名称、省份、城市必填
    agentId   string
    // 网点名称
    pobName   string
    // 网点所在省份
    province   string
    // 网点所在城市
    city   string
    // 客户经理名称
    customerName   string
    // 手机号码
    phone   string
    // 客户经理编码，如果没有可以不填
    otherKey   string
    // json类型，Map<String,String>
    ext   string
    // 活动编码
    activityCode   string
    // token，页面获取到的参数
    token   string
}

// 初始化AlibabaAliqinOfflineCustomerAddRequest对象
func NewAlibabaAliqinOfflineCustomerAddRequest() *AlibabaAliqinOfflineCustomerAddRequest{
    return &AlibabaAliqinOfflineCustomerAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAliqinOfflineCustomerAddRequest) GetApiMethodName() string {
    return "alibaba.aliqin.offline.customer.add"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAliqinOfflineCustomerAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DistributeId Setter
// 代理商id
func (r *AlibabaAliqinOfflineCustomerAddRequest) SetDistributeId(distributeId string) error {
    r.distributeId = distributeId
    r.Set("distribute_id", distributeId)
    return nil
}

// DistributeId Getter
func (r AlibabaAliqinOfflineCustomerAddRequest) GetDistributeId() string {
    return r.distributeId
}
// AgentId Setter
// 网点id，如果存在填写，不存在的话，填0即可；注意：如果填写了这个字段，后面的pob_name等会失效；如果为0，下面的网点名称、省份、城市必填
func (r *AlibabaAliqinOfflineCustomerAddRequest) SetAgentId(agentId string) error {
    r.agentId = agentId
    r.Set("agent_id", agentId)
    return nil
}

// AgentId Getter
func (r AlibabaAliqinOfflineCustomerAddRequest) GetAgentId() string {
    return r.agentId
}
// PobName Setter
// 网点名称
func (r *AlibabaAliqinOfflineCustomerAddRequest) SetPobName(pobName string) error {
    r.pobName = pobName
    r.Set("pob_name", pobName)
    return nil
}

// PobName Getter
func (r AlibabaAliqinOfflineCustomerAddRequest) GetPobName() string {
    return r.pobName
}
// Province Setter
// 网点所在省份
func (r *AlibabaAliqinOfflineCustomerAddRequest) SetProvince(province string) error {
    r.province = province
    r.Set("province", province)
    return nil
}

// Province Getter
func (r AlibabaAliqinOfflineCustomerAddRequest) GetProvince() string {
    return r.province
}
// City Setter
// 网点所在城市
func (r *AlibabaAliqinOfflineCustomerAddRequest) SetCity(city string) error {
    r.city = city
    r.Set("city", city)
    return nil
}

// City Getter
func (r AlibabaAliqinOfflineCustomerAddRequest) GetCity() string {
    return r.city
}
// CustomerName Setter
// 客户经理名称
func (r *AlibabaAliqinOfflineCustomerAddRequest) SetCustomerName(customerName string) error {
    r.customerName = customerName
    r.Set("customer_name", customerName)
    return nil
}

// CustomerName Getter
func (r AlibabaAliqinOfflineCustomerAddRequest) GetCustomerName() string {
    return r.customerName
}
// Phone Setter
// 手机号码
func (r *AlibabaAliqinOfflineCustomerAddRequest) SetPhone(phone string) error {
    r.phone = phone
    r.Set("phone", phone)
    return nil
}

// Phone Getter
func (r AlibabaAliqinOfflineCustomerAddRequest) GetPhone() string {
    return r.phone
}
// OtherKey Setter
// 客户经理编码，如果没有可以不填
func (r *AlibabaAliqinOfflineCustomerAddRequest) SetOtherKey(otherKey string) error {
    r.otherKey = otherKey
    r.Set("other_key", otherKey)
    return nil
}

// OtherKey Getter
func (r AlibabaAliqinOfflineCustomerAddRequest) GetOtherKey() string {
    return r.otherKey
}
// Ext Setter
// json类型，Map<String,String>
func (r *AlibabaAliqinOfflineCustomerAddRequest) SetExt(ext string) error {
    r.ext = ext
    r.Set("ext", ext)
    return nil
}

// Ext Getter
func (r AlibabaAliqinOfflineCustomerAddRequest) GetExt() string {
    return r.ext
}
// ActivityCode Setter
// 活动编码
func (r *AlibabaAliqinOfflineCustomerAddRequest) SetActivityCode(activityCode string) error {
    r.activityCode = activityCode
    r.Set("activity_code", activityCode)
    return nil
}

// ActivityCode Getter
func (r AlibabaAliqinOfflineCustomerAddRequest) GetActivityCode() string {
    return r.activityCode
}
// Token Setter
// token，页面获取到的参数
func (r *AlibabaAliqinOfflineCustomerAddRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

// Token Getter
func (r AlibabaAliqinOfflineCustomerAddRequest) GetToken() string {
    return r.token
}
