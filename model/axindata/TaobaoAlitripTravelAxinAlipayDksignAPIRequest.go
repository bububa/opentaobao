package axindata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelAxinAlipayDksignAPIRequest 支付宝代扣签约 API请求
// taobao.alitrip.travel.axin.alipay.dksign
//
// 提供支付宝代扣签约服务
type TaobaoAlitripTravelAxinAlipayDksignAPIRequest struct {
	model.Params
	// 登录用户id
	_externalLoginId string
}

// NewTaobaoAlitripTravelAxinAlipayDksignRequest 初始化TaobaoAlitripTravelAxinAlipayDksignAPIRequest对象
func NewTaobaoAlitripTravelAxinAlipayDksignRequest() *TaobaoAlitripTravelAxinAlipayDksignAPIRequest {
	return &TaobaoAlitripTravelAxinAlipayDksignAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelAxinAlipayDksignAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.axin.alipay.dksign"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelAxinAlipayDksignAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetExternalLoginId is ExternalLoginId Setter
// 登录用户id
func (r *TaobaoAlitripTravelAxinAlipayDksignAPIRequest) SetExternalLoginId(_externalLoginId string) error {
	r._externalLoginId = _externalLoginId
	r.Set("external_login_id", _externalLoginId)
	return nil
}

// GetExternalLoginId ExternalLoginId Getter
func (r TaobaoAlitripTravelAxinAlipayDksignAPIRequest) GetExternalLoginId() string {
	return r._externalLoginId
}
