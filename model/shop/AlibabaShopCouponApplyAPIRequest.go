package shop

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabashopcouponapplyAPIRequest 通用店铺券领券接口 API请求
// alibaba.shop.coupon.apply
//
// 店铺小部件和模块开发的isv通用店铺券领券接口
type AlibabashopcouponapplyAPIRequest struct {
	model.Params
	// 券的uuid
	_uuid string
	// 买家的openId
	_openId string
}

// NewAlibabashopcouponapplyRequest 初始化AlibabashopcouponapplyAPIRequest对象
func NewAlibabashopcouponapplyRequest() *AlibabashopcouponapplyAPIRequest {
	return &AlibabashopcouponapplyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabashopcouponapplyAPIRequest) GetApiMethodName() string {
	return "alibaba.shop.coupon.apply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabashopcouponapplyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabashopcouponapplyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUuid is Uuid Setter
// 券的uuid
func (r *AlibabashopcouponapplyAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r AlibabashopcouponapplyAPIRequest) GetUuid() string {
	return r._uuid
}

// SetOpenId is OpenId Setter
// 买家的openId
func (r *AlibabashopcouponapplyAPIRequest) SetOpenId(_openId string) error {
	r._openId = _openId
	r.Set("open_id", _openId)
	return nil
}

// GetOpenId OpenId Getter
func (r AlibabashopcouponapplyAPIRequest) GetOpenId() string {
	return r._openId
}
