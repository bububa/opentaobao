package nrt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallnrtcouponsendAPIRequest 券发放接口 API请求
// tmall.nrt.coupon.send
//
// 新零售场景，商家自有渠道发放券
type TmallnrtcouponsendAPIRequest struct {
	model.Params
	// 发券dto
	_nrtCouponSendDto *NrtCouponSendDto
}

// NewTmallnrtcouponsendRequest 初始化TmallnrtcouponsendAPIRequest对象
func NewTmallnrtcouponsendRequest() *TmallnrtcouponsendAPIRequest {
	return &TmallnrtcouponsendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallnrtcouponsendAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.coupon.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallnrtcouponsendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallnrtcouponsendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNrtCouponSendDto is NrtCouponSendDto Setter
// 发券dto
func (r *TmallnrtcouponsendAPIRequest) SetNrtCouponSendDto(_nrtCouponSendDto *NrtCouponSendDto) error {
	r._nrtCouponSendDto = _nrtCouponSendDto
	r.Set("nrt_coupon_send_dto", _nrtCouponSendDto)
	return nil
}

// GetNrtCouponSendDto NrtCouponSendDto Getter
func (r TmallnrtcouponsendAPIRequest) GetNrtCouponSendDto() *NrtCouponSendDto {
	return r._nrtCouponSendDto
}
