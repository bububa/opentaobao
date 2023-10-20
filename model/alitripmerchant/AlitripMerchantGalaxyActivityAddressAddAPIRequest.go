package alitripmerchant

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyActivityAddressAddAPIRequest 星河-营销抽奖奖品邮寄地址添加 API请求
// alitrip.merchant.galaxy.activity.address.add
//
// 营销抽奖活动，奖品邮寄地址填写
type AlitripMerchantGalaxyActivityAddressAddAPIRequest struct {
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

// NewAlitripMerchantGalaxyActivityAddressAddRequest 初始化AlitripMerchantGalaxyActivityAddressAddAPIRequest对象
func NewAlitripMerchantGalaxyActivityAddressAddRequest() *AlitripMerchantGalaxyActivityAddressAddAPIRequest {
	return &AlitripMerchantGalaxyActivityAddressAddAPIRequest{
		Params: model.NewParams(8),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripMerchantGalaxyActivityAddressAddAPIRequest) Reset() {
	r._tenantKey = ""
	r._token = ""
	r._address = ""
	r._phone = ""
	r._name = ""
	r._offerId = 0
	r._recordId = 0
	r._goodsId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyActivityAddressAddAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.activity.address.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyActivityAddressAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyActivityAddressAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户身份信息
func (r *AlitripMerchantGalaxyActivityAddressAddAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyActivityAddressAddAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户token
func (r *AlitripMerchantGalaxyActivityAddressAddAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyActivityAddressAddAPIRequest) GetToken() string {
	return r._token
}

// SetAddress is Address Setter
// 收货人地址
func (r *AlitripMerchantGalaxyActivityAddressAddAPIRequest) SetAddress(_address string) error {
	r._address = _address
	r.Set("address", _address)
	return nil
}

// GetAddress Address Getter
func (r AlitripMerchantGalaxyActivityAddressAddAPIRequest) GetAddress() string {
	return r._address
}

// SetPhone is Phone Setter
// 收货人电话
func (r *AlitripMerchantGalaxyActivityAddressAddAPIRequest) SetPhone(_phone string) error {
	r._phone = _phone
	r.Set("phone", _phone)
	return nil
}

// GetPhone Phone Getter
func (r AlitripMerchantGalaxyActivityAddressAddAPIRequest) GetPhone() string {
	return r._phone
}

// SetName is Name Setter
// 收货人姓名
func (r *AlitripMerchantGalaxyActivityAddressAddAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r AlitripMerchantGalaxyActivityAddressAddAPIRequest) GetName() string {
	return r._name
}

// SetOfferId is OfferId Setter
// 活动id
func (r *AlitripMerchantGalaxyActivityAddressAddAPIRequest) SetOfferId(_offerId int64) error {
	r._offerId = _offerId
	r.Set("offer_id", _offerId)
	return nil
}

// GetOfferId OfferId Getter
func (r AlitripMerchantGalaxyActivityAddressAddAPIRequest) GetOfferId() int64 {
	return r._offerId
}

// SetRecordId is RecordId Setter
// 活动记录ID
func (r *AlitripMerchantGalaxyActivityAddressAddAPIRequest) SetRecordId(_recordId int64) error {
	r._recordId = _recordId
	r.Set("record_id", _recordId)
	return nil
}

// GetRecordId RecordId Getter
func (r AlitripMerchantGalaxyActivityAddressAddAPIRequest) GetRecordId() int64 {
	return r._recordId
}

// SetGoodsId is GoodsId Setter
// 奖品ID
func (r *AlitripMerchantGalaxyActivityAddressAddAPIRequest) SetGoodsId(_goodsId int64) error {
	r._goodsId = _goodsId
	r.Set("goods_id", _goodsId)
	return nil
}

// GetGoodsId GoodsId Getter
func (r AlitripMerchantGalaxyActivityAddressAddAPIRequest) GetGoodsId() int64 {
	return r._goodsId
}

var poolAlitripMerchantGalaxyActivityAddressAddAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripMerchantGalaxyActivityAddressAddRequest()
	},
}

// GetAlitripMerchantGalaxyActivityAddressAddRequest 从 sync.Pool 获取 AlitripMerchantGalaxyActivityAddressAddAPIRequest
func GetAlitripMerchantGalaxyActivityAddressAddAPIRequest() *AlitripMerchantGalaxyActivityAddressAddAPIRequest {
	return poolAlitripMerchantGalaxyActivityAddressAddAPIRequest.Get().(*AlitripMerchantGalaxyActivityAddressAddAPIRequest)
}

// ReleaseAlitripMerchantGalaxyActivityAddressAddAPIRequest 将 AlitripMerchantGalaxyActivityAddressAddAPIRequest 放入 sync.Pool
func ReleaseAlitripMerchantGalaxyActivityAddressAddAPIRequest(v *AlitripMerchantGalaxyActivityAddressAddAPIRequest) {
	v.Reset()
	poolAlitripMerchantGalaxyActivityAddressAddAPIRequest.Put(v)
}
