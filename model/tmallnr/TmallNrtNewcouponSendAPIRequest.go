package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallnrtnewcouponsendAPIRequest 券发放接口 API请求
// tmall.nrt.newcoupon.send
//
// 券发放接口
type TmallnrtnewcouponsendAPIRequest struct {
	model.Params
	// 发券dto
	_nrtCouponSendDto *NrtCouponSendDto
}

// NewTmallnrtnewcouponsendRequest 初始化TmallnrtnewcouponsendAPIRequest对象
func NewTmallnrtnewcouponsendRequest() *TmallnrtnewcouponsendAPIRequest {
	return &TmallnrtnewcouponsendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallnrtnewcouponsendAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.newcoupon.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallnrtnewcouponsendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallnrtnewcouponsendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNrtCouponSendDto is NrtCouponSendDto Setter
// 发券dto
func (r *TmallnrtnewcouponsendAPIRequest) SetNrtCouponSendDto(_nrtCouponSendDto *NrtCouponSendDto) error {
	r._nrtCouponSendDto = _nrtCouponSendDto
	r.Set("nrt_coupon_send_dto", _nrtCouponSendDto)
	return nil
}

// GetNrtCouponSendDto NrtCouponSendDto Getter
func (r TmallnrtnewcouponsendAPIRequest) GetNrtCouponSendDto() *NrtCouponSendDto {
	return r._nrtCouponSendDto
}
