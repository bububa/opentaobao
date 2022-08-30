package shop

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaShopCouponApplyAPIRequest 通用店铺券领券接口 API请求
// alibaba.shop.coupon.apply
//
// 店铺小部件和模块开发的isv通用店铺券领券接口
type AlibabaShopCouponApplyAPIRequest struct {
	model.Params
	// 买家的openId
	_openId string
	// 券的uuid
	_uuid string
}

// NewAlibabaShopCouponApplyRequest 初始化AlibabaShopCouponApplyAPIRequest对象
func NewAlibabaShopCouponApplyRequest() *AlibabaShopCouponApplyAPIRequest {
	return &AlibabaShopCouponApplyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaShopCouponApplyAPIRequest) GetApiMethodName() string {
	return "alibaba.shop.coupon.apply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaShopCouponApplyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOpenId is OpenId Setter
// 买家的openId
func (r *AlibabaShopCouponApplyAPIRequest) SetOpenId(_openId string) error {
	r._openId = _openId
	r.Set("open_id", _openId)
	return nil
}

// GetOpenId OpenId Getter
func (r AlibabaShopCouponApplyAPIRequest) GetOpenId() string {
	return r._openId
}

// SetUuid is Uuid Setter
// 券的uuid
func (r *AlibabaShopCouponApplyAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r AlibabaShopCouponApplyAPIRequest) GetUuid() string {
	return r._uuid
}
