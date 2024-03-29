package tmallhk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallHkClearanceGetAPIRequest 天猫国际-清关材料查询 API请求
// tmall.hk.clearance.get
//
// 提供订单收货人身份信息查询功能。
type TmallHkClearanceGetAPIRequest struct {
	model.Params
	// 天猫国际订单号
	_orderId int64
	// 是否需要身份证图片，不需要可以缩短接口响应时间
	_needImage bool
}

// NewTmallHkClearanceGetRequest 初始化TmallHkClearanceGetAPIRequest对象
func NewTmallHkClearanceGetRequest() *TmallHkClearanceGetAPIRequest {
	return &TmallHkClearanceGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallHkClearanceGetAPIRequest) Reset() {
	r._orderId = 0
	r._needImage = false
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallHkClearanceGetAPIRequest) GetApiMethodName() string {
	return "tmall.hk.clearance.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallHkClearanceGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallHkClearanceGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 天猫国际订单号
func (r *TmallHkClearanceGetAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TmallHkClearanceGetAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// SetNeedImage is NeedImage Setter
// 是否需要身份证图片，不需要可以缩短接口响应时间
func (r *TmallHkClearanceGetAPIRequest) SetNeedImage(_needImage bool) error {
	r._needImage = _needImage
	r.Set("need_image", _needImage)
	return nil
}

// GetNeedImage NeedImage Getter
func (r TmallHkClearanceGetAPIRequest) GetNeedImage() bool {
	return r._needImage
}

var poolTmallHkClearanceGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallHkClearanceGetRequest()
	},
}

// GetTmallHkClearanceGetRequest 从 sync.Pool 获取 TmallHkClearanceGetAPIRequest
func GetTmallHkClearanceGetAPIRequest() *TmallHkClearanceGetAPIRequest {
	return poolTmallHkClearanceGetAPIRequest.Get().(*TmallHkClearanceGetAPIRequest)
}

// ReleaseTmallHkClearanceGetAPIRequest 将 TmallHkClearanceGetAPIRequest 放入 sync.Pool
func ReleaseTmallHkClearanceGetAPIRequest(v *TmallHkClearanceGetAPIRequest) {
	v.Reset()
	poolTmallHkClearanceGetAPIRequest.Put(v)
}
