package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallmsfreceiveAPIRequest 签收接口 API请求
// tmall.msf.receive
//
// 签收接口
type TmallmsfreceiveAPIRequest struct {
	model.Params
	// 1
	_shopId string
	// 1
	_bizType string
	// 1
	_code string
}

// NewTmallmsfreceiveRequest 初始化TmallmsfreceiveAPIRequest对象
func NewTmallmsfreceiveRequest() *TmallmsfreceiveAPIRequest {
	return &TmallmsfreceiveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallmsfreceiveAPIRequest) GetApiMethodName() string {
	return "tmall.msf.receive"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallmsfreceiveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallmsfreceiveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetShopId is ShopId Setter
// 1
func (r *TmallmsfreceiveAPIRequest) SetShopId(_shopId string) error {
	r._shopId = _shopId
	r.Set("shop_id", _shopId)
	return nil
}

// GetShopId ShopId Getter
func (r TmallmsfreceiveAPIRequest) GetShopId() string {
	return r._shopId
}

// SetBizType is BizType Setter
// 1
func (r *TmallmsfreceiveAPIRequest) SetBizType(_bizType string) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// GetBizType BizType Getter
func (r TmallmsfreceiveAPIRequest) GetBizType() string {
	return r._bizType
}

// SetCode is Code Setter
// 1
func (r *TmallmsfreceiveAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r TmallmsfreceiveAPIRequest) GetCode() string {
	return r._code
}
