package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyactivityaddressaddAPIRequest 星河-营销抽奖奖品邮寄地址添加 API请求
// alitrip.merchant.galaxy.activity.address.add
//
// 营销抽奖活动，奖品邮寄地址填写
type AlitripmerchantgalaxyactivityaddressaddAPIRequest struct {
	model.Params
	// 租户身份信息
	_tenantKey string
	// 用户token
	_token string
	// 收货人地址
	_address string
	// 收货人电话
	_phone string
	// 收货人姓名
	_name string
	// 活动id
	_offerId int64
	// 活动记录ID
	_recordId int64
	// 奖品ID
	_goodsId int64
}

// NewAlitripmerchantgalaxyactivityaddressaddRequest 初始化AlitripmerchantgalaxyactivityaddressaddAPIRequest对象
func NewAlitripmerchantgalaxyactivityaddressaddRequest() *AlitripmerchantgalaxyactivityaddressaddAPIRequest {
	return &AlitripmerchantgalaxyactivityaddressaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxyactivityaddressaddAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.activity.address.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxyactivityaddressaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxyactivityaddressaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户身份信息
func (r *AlitripmerchantgalaxyactivityaddressaddAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxyactivityaddressaddAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户token
func (r *AlitripmerchantgalaxyactivityaddressaddAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripmerchantgalaxyactivityaddressaddAPIRequest) GetToken() string {
	return r._token
}

// SetAddress is Address Setter
// 收货人地址
func (r *AlitripmerchantgalaxyactivityaddressaddAPIRequest) SetAddress(_address string) error {
	r._address = _address
	r.Set("address", _address)
	return nil
}

// GetAddress Address Getter
func (r AlitripmerchantgalaxyactivityaddressaddAPIRequest) GetAddress() string {
	return r._address
}

// SetPhone is Phone Setter
// 收货人电话
func (r *AlitripmerchantgalaxyactivityaddressaddAPIRequest) SetPhone(_phone string) error {
	r._phone = _phone
	r.Set("phone", _phone)
	return nil
}

// GetPhone Phone Getter
func (r AlitripmerchantgalaxyactivityaddressaddAPIRequest) GetPhone() string {
	return r._phone
}

// SetName is Name Setter
// 收货人姓名
func (r *AlitripmerchantgalaxyactivityaddressaddAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r AlitripmerchantgalaxyactivityaddressaddAPIRequest) GetName() string {
	return r._name
}

// SetOfferId is OfferId Setter
// 活动id
func (r *AlitripmerchantgalaxyactivityaddressaddAPIRequest) SetOfferId(_offerId int64) error {
	r._offerId = _offerId
	r.Set("offer_id", _offerId)
	return nil
}

// GetOfferId OfferId Getter
func (r AlitripmerchantgalaxyactivityaddressaddAPIRequest) GetOfferId() int64 {
	return r._offerId
}

// SetRecordId is RecordId Setter
// 活动记录ID
func (r *AlitripmerchantgalaxyactivityaddressaddAPIRequest) SetRecordId(_recordId int64) error {
	r._recordId = _recordId
	r.Set("record_id", _recordId)
	return nil
}

// GetRecordId RecordId Getter
func (r AlitripmerchantgalaxyactivityaddressaddAPIRequest) GetRecordId() int64 {
	return r._recordId
}

// SetGoodsId is GoodsId Setter
// 奖品ID
func (r *AlitripmerchantgalaxyactivityaddressaddAPIRequest) SetGoodsId(_goodsId int64) error {
	r._goodsId = _goodsId
	r.Set("goods_id", _goodsId)
	return nil
}

// GetGoodsId GoodsId Getter
func (r AlitripmerchantgalaxyactivityaddressaddAPIRequest) GetGoodsId() int64 {
	return r._goodsId
}
