package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtNewcouponSendAPIRequest 券发放接口 API请求
// tmall.nrt.newcoupon.send
//
// 券发放接口
type TmallNrtNewcouponSendAPIRequest struct {
	model.Params
	// 发券dto
	_nrtCouponSendDto *NrtCouponSendDto
}

// NewTmallNrtNewcouponSendRequest 初始化TmallNrtNewcouponSendAPIRequest对象
func NewTmallNrtNewcouponSendRequest() *TmallNrtNewcouponSendAPIRequest {
	return &TmallNrtNewcouponSendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallNrtNewcouponSendAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.newcoupon.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallNrtNewcouponSendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallNrtNewcouponSendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNrtCouponSendDto is NrtCouponSendDto Setter
// 发券dto
func (r *TmallNrtNewcouponSendAPIRequest) SetNrtCouponSendDto(_nrtCouponSendDto *NrtCouponSendDto) error {
	r._nrtCouponSendDto = _nrtCouponSendDto
	r.Set("nrt_coupon_send_dto", _nrtCouponSendDto)
	return nil
}

// GetNrtCouponSendDto NrtCouponSendDto Getter
func (r TmallNrtNewcouponSendAPIRequest) GetNrtCouponSendDto() *NrtCouponSendDto {
	return r._nrtCouponSendDto
}
