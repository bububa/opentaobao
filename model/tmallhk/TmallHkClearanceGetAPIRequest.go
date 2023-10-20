package tmallhk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallhkclearancegetAPIRequest 天猫国际-清关材料查询 API请求
// tmall.hk.clearance.get
//
// 提供订单收货人身份信息查询功能。
type TmallhkclearancegetAPIRequest struct {
	model.Params
	// 天猫国际订单号
	_orderId int64
	// 是否需要身份证图片，不需要可以缩短接口响应时间
	_needImage bool
}

// NewTmallhkclearancegetRequest 初始化TmallhkclearancegetAPIRequest对象
func NewTmallhkclearancegetRequest() *TmallhkclearancegetAPIRequest {
	return &TmallhkclearancegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallhkclearancegetAPIRequest) GetApiMethodName() string {
	return "tmall.hk.clearance.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallhkclearancegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallhkclearancegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 天猫国际订单号
func (r *TmallhkclearancegetAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TmallhkclearancegetAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// SetNeedImage is NeedImage Setter
// 是否需要身份证图片，不需要可以缩短接口响应时间
func (r *TmallhkclearancegetAPIRequest) SetNeedImage(_needImage bool) error {
	r._needImage = _needImage
	r.Set("need_image", _needImage)
	return nil
}

// GetNeedImage NeedImage Getter
func (r TmallhkclearancegetAPIRequest) GetNeedImage() bool {
	return r._needImage
}
