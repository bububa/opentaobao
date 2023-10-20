package alitripmerchant

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryAPIRequest 德比付费会员卡可购查询 API请求
// alitrip.merchant.galaxy.derby.member.voucher.card.purchasable.query
//
// 德比付费会员卡可购查询
type AlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryAPIRequest struct {
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

// NewAlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryRequest 初始化AlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryAPIRequest对象
func NewAlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryRequest() *AlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryAPIRequest {
	return &AlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryAPIRequest) Reset() {
	r._tenantKey = ""
	r._token = ""
	r._voucherCardCategory = ""
	r._checkGoodsOffShelf = false
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.derby.member.voucher.card.purchasable.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 1
func (r *AlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 1
func (r *AlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryAPIRequest) GetToken() string {
	return r._token
}

// SetVoucherCardCategory is VoucherCardCategory Setter
// 商品枚举（DerbyVoucherCardOrderProductMinTypeEnum）  APLUSP:p卡 APLUSB：b卡
func (r *AlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryAPIRequest) SetVoucherCardCategory(_voucherCardCategory string) error {
	r._voucherCardCategory = _voucherCardCategory
	r.Set("voucher_card_category", _voucherCardCategory)
	return nil
}

// GetVoucherCardCategory VoucherCardCategory Getter
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryAPIRequest) GetVoucherCardCategory() string {
	return r._voucherCardCategory
}

// SetCheckGoodsOffShelf is CheckGoodsOffShelf Setter
// 接口用于查询商品是否下架，需要缓存
func (r *AlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryAPIRequest) SetCheckGoodsOffShelf(_checkGoodsOffShelf bool) error {
	r._checkGoodsOffShelf = _checkGoodsOffShelf
	r.Set("check_goods_off_shelf", _checkGoodsOffShelf)
	return nil
}

// GetCheckGoodsOffShelf CheckGoodsOffShelf Getter
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryAPIRequest) GetCheckGoodsOffShelf() bool {
	return r._checkGoodsOffShelf
}

var poolAlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryRequest()
	},
}

// GetAlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryRequest 从 sync.Pool 获取 AlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryAPIRequest
func GetAlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryAPIRequest() *AlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryAPIRequest {
	return poolAlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryAPIRequest.Get().(*AlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryAPIRequest)
}

// ReleaseAlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryAPIRequest 将 AlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryAPIRequest 放入 sync.Pool
func ReleaseAlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryAPIRequest(v *AlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryAPIRequest) {
	v.Reset()
	poolAlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryAPIRequest.Put(v)
}
