package nrt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallNrtCouponSendAPIRequest
券发放接口 API请求
tmall.nrt.coupon.send

新零售场景，商家自有渠道发放券 */
type TmallNrtCouponSendAPIRequest struct {
	model.Params
	// 发券dto
	_nrtCouponSendDto *NrtCouponSendDto
}

// NewTmallNrtCouponSendRequest 初始化TmallNrtCouponSendAPIRequest对象
func NewTmallNrtCouponSendRequest() *TmallNrtCouponSendAPIRequest {
	return &TmallNrtCouponSendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallNrtCouponSendAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.coupon.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallNrtCouponSendAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is NrtCouponSendDto Setter
// 发券dto
func (r *TmallNrtCouponSendAPIRequest) SetNrtCouponSendDto(_nrtCouponSendDto *NrtCouponSendDto) error {
	r._nrtCouponSendDto = _nrtCouponSendDto
	r.Set("nrt_coupon_send_dto", _nrtCouponSendDto)
	return nil
}

// Get NrtCouponSendDto Getter
func (r TmallNrtCouponSendAPIRequest) GetNrtCouponSendDto() *NrtCouponSendDto {
	return r._nrtCouponSendDto
}
