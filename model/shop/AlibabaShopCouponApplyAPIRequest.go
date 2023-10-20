package shop

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaShopCouponApplyAPIRequest 通用店铺券领券接口 API请求
// alibaba.shop.coupon.apply
//
// 店铺小部件和模块开发的isv通用店铺券领券接口
type AlibabaShopCouponApplyAPIRequest struct {
	model.Params
	// 券的uuid
	_uuid string
	// 买家的openId
	_openId string
}

// NewAlibabaShopCouponApplyRequest 初始化AlibabaShopCouponApplyAPIRequest对象
func NewAlibabaShopCouponApplyRequest() *AlibabaShopCouponApplyAPIRequest {
	return &AlibabaShopCouponApplyAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaShopCouponApplyAPIRequest) Reset() {
	r._uuid = ""
	r._openId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaShopCouponApplyAPIRequest) GetApiMethodName() string {
	return "alibaba.shop.coupon.apply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaShopCouponApplyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaShopCouponApplyAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaShopCouponApplyAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaShopCouponApplyRequest()
	},
}

// GetAlibabaShopCouponApplyRequest 从 sync.Pool 获取 AlibabaShopCouponApplyAPIRequest
func GetAlibabaShopCouponApplyAPIRequest() *AlibabaShopCouponApplyAPIRequest {
	return poolAlibabaShopCouponApplyAPIRequest.Get().(*AlibabaShopCouponApplyAPIRequest)
}

// ReleaseAlibabaShopCouponApplyAPIRequest 将 AlibabaShopCouponApplyAPIRequest 放入 sync.Pool
func ReleaseAlibabaShopCouponApplyAPIRequest(v *AlibabaShopCouponApplyAPIRequest) {
	v.Reset()
	poolAlibabaShopCouponApplyAPIRequest.Put(v)
}
