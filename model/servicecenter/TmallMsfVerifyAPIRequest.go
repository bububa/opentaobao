package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallMsfVerifyAPIRequest
喵师傅核销接口 API请求
tmall.msf.verify

msf服务核销的top接口 */
type TmallMsfVerifyAPIRequest struct {
	model.Params
	// 111
	_shopId string
	// 111
	_bizType string
	// 111
	_code string
}

// NewTmallMsfVerifyRequest 初始化TmallMsfVerifyAPIRequest对象
func NewTmallMsfVerifyRequest() *TmallMsfVerifyAPIRequest {
	return &TmallMsfVerifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallMsfVerifyAPIRequest) GetApiMethodName() string {
	return "tmall.msf.verify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallMsfVerifyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ShopId Setter
// 111
func (r *TmallMsfVerifyAPIRequest) SetShopId(_shopId string) error {
	r._shopId = _shopId
	r.Set("shop_id", _shopId)
	return nil
}

// Get ShopId Getter
func (r TmallMsfVerifyAPIRequest) GetShopId() string {
	return r._shopId
}

// Set is BizType Setter
// 111
func (r *TmallMsfVerifyAPIRequest) SetBizType(_bizType string) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// Get BizType Getter
func (r TmallMsfVerifyAPIRequest) GetBizType() string {
	return r._bizType
}

// Set is Code Setter
// 111
func (r *TmallMsfVerifyAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// Get Code Getter
func (r TmallMsfVerifyAPIRequest) GetCode() string {
	return r._code
}
