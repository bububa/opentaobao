package tmallhk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallHkClearanceDistributionGetAPIRequest 分销供应商获取清关材料 API请求
// tmall.hk.clearance.distribution.get
//
// 供销体系下，提供供应商可以直接获取其订单身份证信息的接口，以使其完成清关。
type TmallHkClearanceDistributionGetAPIRequest struct {
	model.Params
	// 订单号
	_orderId int64
	// 是否需要身份证图片，不需要可以缩短接口响应时间
	_needImage bool
}

// NewTmallHkClearanceDistributionGetRequest 初始化TmallHkClearanceDistributionGetAPIRequest对象
func NewTmallHkClearanceDistributionGetRequest() *TmallHkClearanceDistributionGetAPIRequest {
	return &TmallHkClearanceDistributionGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallHkClearanceDistributionGetAPIRequest) Reset() {
	r._orderId = 0
	r._needImage = false
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallHkClearanceDistributionGetAPIRequest) GetApiMethodName() string {
	return "tmall.hk.clearance.distribution.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallHkClearanceDistributionGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallHkClearanceDistributionGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 订单号
func (r *TmallHkClearanceDistributionGetAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TmallHkClearanceDistributionGetAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// SetNeedImage is NeedImage Setter
// 是否需要身份证图片，不需要可以缩短接口响应时间
func (r *TmallHkClearanceDistributionGetAPIRequest) SetNeedImage(_needImage bool) error {
	r._needImage = _needImage
	r.Set("need_image", _needImage)
	return nil
}

// GetNeedImage NeedImage Getter
func (r TmallHkClearanceDistributionGetAPIRequest) GetNeedImage() bool {
	return r._needImage
}

var poolTmallHkClearanceDistributionGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallHkClearanceDistributionGetRequest()
	},
}

// GetTmallHkClearanceDistributionGetRequest 从 sync.Pool 获取 TmallHkClearanceDistributionGetAPIRequest
func GetTmallHkClearanceDistributionGetAPIRequest() *TmallHkClearanceDistributionGetAPIRequest {
	return poolTmallHkClearanceDistributionGetAPIRequest.Get().(*TmallHkClearanceDistributionGetAPIRequest)
}

// ReleaseTmallHkClearanceDistributionGetAPIRequest 将 TmallHkClearanceDistributionGetAPIRequest 放入 sync.Pool
func ReleaseTmallHkClearanceDistributionGetAPIRequest(v *TmallHkClearanceDistributionGetAPIRequest) {
	v.Reset()
	poolTmallHkClearanceDistributionGetAPIRequest.Put(v)
}
