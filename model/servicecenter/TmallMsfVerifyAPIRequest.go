package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallmsfverifyAPIRequest 喵师傅核销接口 API请求
// tmall.msf.verify
//
// msf服务核销的top接口
type TmallmsfverifyAPIRequest struct {
	model.Params
	// 111
	_shopId string
	// 111
	_bizType string
	// 111
	_code string
}

// NewTmallmsfverifyRequest 初始化TmallmsfverifyAPIRequest对象
func NewTmallmsfverifyRequest() *TmallmsfverifyAPIRequest {
	return &TmallmsfverifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallmsfverifyAPIRequest) GetApiMethodName() string {
	return "tmall.msf.verify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallmsfverifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallmsfverifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetShopId is ShopId Setter
// 111
func (r *TmallmsfverifyAPIRequest) SetShopId(_shopId string) error {
	r._shopId = _shopId
	r.Set("shop_id", _shopId)
	return nil
}

// GetShopId ShopId Getter
func (r TmallmsfverifyAPIRequest) GetShopId() string {
	return r._shopId
}

// SetBizType is BizType Setter
// 111
func (r *TmallmsfverifyAPIRequest) SetBizType(_bizType string) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// GetBizType BizType Getter
func (r TmallmsfverifyAPIRequest) GetBizType() string {
	return r._bizType
}

// SetCode is Code Setter
// 111
func (r *TmallmsfverifyAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r TmallmsfverifyAPIRequest) GetCode() string {
	return r._code
}
