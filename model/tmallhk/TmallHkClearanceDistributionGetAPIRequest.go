package tmallhk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallhkclearancedistributiongetAPIRequest 分销供应商获取清关材料 API请求
// tmall.hk.clearance.distribution.get
//
// 供销体系下，提供供应商可以直接获取其订单身份证信息的接口，以使其完成清关。
type TmallhkclearancedistributiongetAPIRequest struct {
	model.Params
	// 订单号
	_orderId int64
	// 是否需要身份证图片，不需要可以缩短接口响应时间
	_needImage bool
}

// NewTmallhkclearancedistributiongetRequest 初始化TmallhkclearancedistributiongetAPIRequest对象
func NewTmallhkclearancedistributiongetRequest() *TmallhkclearancedistributiongetAPIRequest {
	return &TmallhkclearancedistributiongetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallhkclearancedistributiongetAPIRequest) GetApiMethodName() string {
	return "tmall.hk.clearance.distribution.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallhkclearancedistributiongetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallhkclearancedistributiongetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 订单号
func (r *TmallhkclearancedistributiongetAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TmallhkclearancedistributiongetAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// SetNeedImage is NeedImage Setter
// 是否需要身份证图片，不需要可以缩短接口响应时间
func (r *TmallhkclearancedistributiongetAPIRequest) SetNeedImage(_needImage bool) error {
	r._needImage = _needImage
	r.Set("need_image", _needImage)
	return nil
}

// GetNeedImage NeedImage Getter
func (r TmallhkclearancedistributiongetAPIRequest) GetNeedImage() bool {
	return r._needImage
}
