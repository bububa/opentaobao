package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyderbymembervouchercardpurchasablequeryAPIRequest 德比付费会员卡可购查询 API请求
// alitrip.merchant.galaxy.derby.member.voucher.card.purchasable.query
//
// 德比付费会员卡可购查询
type AlitripmerchantgalaxyderbymembervouchercardpurchasablequeryAPIRequest struct {
	model.Params
	// 1
	_tenantKey string
	// 1
	_token string
	// 商品枚举（DerbyVoucherCardOrderProductMinTypeEnum）  APLUSP:p卡 APLUSB：b卡
	_voucherCardCategory string
	// 接口用于查询商品是否下架，需要缓存
	_checkGoodsOffShelf bool
}

// NewAlitripmerchantgalaxyderbymembervouchercardpurchasablequeryRequest 初始化AlitripmerchantgalaxyderbymembervouchercardpurchasablequeryAPIRequest对象
func NewAlitripmerchantgalaxyderbymembervouchercardpurchasablequeryRequest() *AlitripmerchantgalaxyderbymembervouchercardpurchasablequeryAPIRequest {
	return &AlitripmerchantgalaxyderbymembervouchercardpurchasablequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxyderbymembervouchercardpurchasablequeryAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.derby.member.voucher.card.purchasable.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxyderbymembervouchercardpurchasablequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxyderbymembervouchercardpurchasablequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 1
func (r *AlitripmerchantgalaxyderbymembervouchercardpurchasablequeryAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxyderbymembervouchercardpurchasablequeryAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 1
func (r *AlitripmerchantgalaxyderbymembervouchercardpurchasablequeryAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripmerchantgalaxyderbymembervouchercardpurchasablequeryAPIRequest) GetToken() string {
	return r._token
}

// SetVoucherCardCategory is VoucherCardCategory Setter
// 商品枚举（DerbyVoucherCardOrderProductMinTypeEnum）  APLUSP:p卡 APLUSB：b卡
func (r *AlitripmerchantgalaxyderbymembervouchercardpurchasablequeryAPIRequest) SetVoucherCardCategory(_voucherCardCategory string) error {
	r._voucherCardCategory = _voucherCardCategory
	r.Set("voucher_card_category", _voucherCardCategory)
	return nil
}

// GetVoucherCardCategory VoucherCardCategory Getter
func (r AlitripmerchantgalaxyderbymembervouchercardpurchasablequeryAPIRequest) GetVoucherCardCategory() string {
	return r._voucherCardCategory
}

// SetCheckGoodsOffShelf is CheckGoodsOffShelf Setter
// 接口用于查询商品是否下架，需要缓存
func (r *AlitripmerchantgalaxyderbymembervouchercardpurchasablequeryAPIRequest) SetCheckGoodsOffShelf(_checkGoodsOffShelf bool) error {
	r._checkGoodsOffShelf = _checkGoodsOffShelf
	r.Set("check_goods_off_shelf", _checkGoodsOffShelf)
	return nil
}

// GetCheckGoodsOffShelf CheckGoodsOffShelf Getter
func (r AlitripmerchantgalaxyderbymembervouchercardpurchasablequeryAPIRequest) GetCheckGoodsOffShelf() bool {
	return r._checkGoodsOffShelf
}
