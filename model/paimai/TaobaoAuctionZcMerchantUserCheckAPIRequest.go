package paimai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoauctionzcmerchantusercheckAPIRequest 通过手机号确认阿里资产商家 API请求
// taobao.auction.zc.merchant.user.check
//
// 通过手机号确认阿里资产商家
type TaobaoauctionzcmerchantusercheckAPIRequest struct {
	model.Params
	// 手机号
	_mobile string
}

// NewTaobaoauctionzcmerchantusercheckRequest 初始化TaobaoauctionzcmerchantusercheckAPIRequest对象
func NewTaobaoauctionzcmerchantusercheckRequest() *TaobaoauctionzcmerchantusercheckAPIRequest {
	return &TaobaoauctionzcmerchantusercheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoauctionzcmerchantusercheckAPIRequest) GetApiMethodName() string {
	return "taobao.auction.zc.merchant.user.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoauctionzcmerchantusercheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoauctionzcmerchantusercheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMobile is Mobile Setter
// 手机号
func (r *TaobaoauctionzcmerchantusercheckAPIRequest) SetMobile(_mobile string) error {
	r._mobile = _mobile
	r.Set("mobile", _mobile)
	return nil
}

// GetMobile Mobile Getter
func (r TaobaoauctionzcmerchantusercheckAPIRequest) GetMobile() string {
	return r._mobile
}
