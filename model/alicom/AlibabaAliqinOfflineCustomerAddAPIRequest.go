package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinOfflineCustomerAddAPIRequest 系外拉新代理商增加客户经理接口 API请求
// alibaba.aliqin.offline.customer.add
//
// 阿里通信这边维护了代理商和其对应的客户经理的关系，用于业务处理，开放该接口用于代理商将他们系统下的客户经理信息同步给我们
type AlibabaAliqinOfflineCustomerAddAPIRequest struct {
	model.Params
	// 代理商id
	_distributeId string
	// 网点id，如果存在填写，不存在的话，填0即可；注意：如果填写了这个字段，后面的pob_name等会失效；如果为0，下面的网点名称、省份、城市必填
	_agentId string
	// 网点名称
	_pobName string
	// 网点所在省份
	_province string
	// 网点所在城市
	_city string
	// 客户经理名称
	_customerName string
	// 手机号码
	_phone string
	// 客户经理编码，如果没有可以不填
	_otherKey string
	// json类型，Map<String,String>
	_ext string
	// 活动编码
	_activityCode string
	// token，页面获取到的参数
	_token string
}

// NewAlibabaAliqinOfflineCustomerAddRequest 初始化AlibabaAliqinOfflineCustomerAddAPIRequest对象
func NewAlibabaAliqinOfflineCustomerAddRequest() *AlibabaAliqinOfflineCustomerAddAPIRequest {
	return &AlibabaAliqinOfflineCustomerAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAliqinOfflineCustomerAddAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.offline.customer.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAliqinOfflineCustomerAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetDistributeId is DistributeId Setter
// 代理商id
func (r *AlibabaAliqinOfflineCustomerAddAPIRequest) SetDistributeId(_distributeId string) error {
	r._distributeId = _distributeId
	r.Set("distribute_id", _distributeId)
	return nil
}

// GetDistributeId DistributeId Getter
func (r AlibabaAliqinOfflineCustomerAddAPIRequest) GetDistributeId() string {
	return r._distributeId
}

// SetAgentId is AgentId Setter
// 网点id，如果存在填写，不存在的话，填0即可；注意：如果填写了这个字段，后面的pob_name等会失效；如果为0，下面的网点名称、省份、城市必填
func (r *AlibabaAliqinOfflineCustomerAddAPIRequest) SetAgentId(_agentId string) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// GetAgentId AgentId Getter
func (r AlibabaAliqinOfflineCustomerAddAPIRequest) GetAgentId() string {
	return r._agentId
}

// SetPobName is PobName Setter
// 网点名称
func (r *AlibabaAliqinOfflineCustomerAddAPIRequest) SetPobName(_pobName string) error {
	r._pobName = _pobName
	r.Set("pob_name", _pobName)
	return nil
}

// GetPobName PobName Getter
func (r AlibabaAliqinOfflineCustomerAddAPIRequest) GetPobName() string {
	return r._pobName
}

// SetProvince is Province Setter
// 网点所在省份
func (r *AlibabaAliqinOfflineCustomerAddAPIRequest) SetProvince(_province string) error {
	r._province = _province
	r.Set("province", _province)
	return nil
}

// GetProvince Province Getter
func (r AlibabaAliqinOfflineCustomerAddAPIRequest) GetProvince() string {
	return r._province
}

// SetCity is City Setter
// 网点所在城市
func (r *AlibabaAliqinOfflineCustomerAddAPIRequest) SetCity(_city string) error {
	r._city = _city
	r.Set("city", _city)
	return nil
}

// GetCity City Getter
func (r AlibabaAliqinOfflineCustomerAddAPIRequest) GetCity() string {
	return r._city
}

// SetCustomerName is CustomerName Setter
// 客户经理名称
func (r *AlibabaAliqinOfflineCustomerAddAPIRequest) SetCustomerName(_customerName string) error {
	r._customerName = _customerName
	r.Set("customer_name", _customerName)
	return nil
}

// GetCustomerName CustomerName Getter
func (r AlibabaAliqinOfflineCustomerAddAPIRequest) GetCustomerName() string {
	return r._customerName
}

// SetPhone is Phone Setter
// 手机号码
func (r *AlibabaAliqinOfflineCustomerAddAPIRequest) SetPhone(_phone string) error {
	r._phone = _phone
	r.Set("phone", _phone)
	return nil
}

// GetPhone Phone Getter
func (r AlibabaAliqinOfflineCustomerAddAPIRequest) GetPhone() string {
	return r._phone
}

// SetOtherKey is OtherKey Setter
// 客户经理编码，如果没有可以不填
func (r *AlibabaAliqinOfflineCustomerAddAPIRequest) SetOtherKey(_otherKey string) error {
	r._otherKey = _otherKey
	r.Set("other_key", _otherKey)
	return nil
}

// GetOtherKey OtherKey Getter
func (r AlibabaAliqinOfflineCustomerAddAPIRequest) GetOtherKey() string {
	return r._otherKey
}

// SetExt is Ext Setter
// json类型，Map<String,String>
func (r *AlibabaAliqinOfflineCustomerAddAPIRequest) SetExt(_ext string) error {
	r._ext = _ext
	r.Set("ext", _ext)
	return nil
}

// GetExt Ext Getter
func (r AlibabaAliqinOfflineCustomerAddAPIRequest) GetExt() string {
	return r._ext
}

// SetActivityCode is ActivityCode Setter
// 活动编码
func (r *AlibabaAliqinOfflineCustomerAddAPIRequest) SetActivityCode(_activityCode string) error {
	r._activityCode = _activityCode
	r.Set("activity_code", _activityCode)
	return nil
}

// GetActivityCode ActivityCode Getter
func (r AlibabaAliqinOfflineCustomerAddAPIRequest) GetActivityCode() string {
	return r._activityCode
}

// SetToken is Token Setter
// token，页面获取到的参数
func (r *AlibabaAliqinOfflineCustomerAddAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlibabaAliqinOfflineCustomerAddAPIRequest) GetToken() string {
	return r._token
}
