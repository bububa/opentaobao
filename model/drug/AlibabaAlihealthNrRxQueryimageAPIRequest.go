package drug

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthNrRxQueryimageAPIRequest o2o查看处方图片 API请求
// alibaba.alihealth.nr.rx.queryimage
//
// o2o商家查看处方图片，包括电子图片与纸质图片
type AlibabaAlihealthNrRxQueryimageAPIRequest struct {
	model.Params
	// 订单编号
	_orderId int64
}

// NewAlibabaAlihealthNrRxQueryimageRequest 初始化AlibabaAlihealthNrRxQueryimageAPIRequest对象
func NewAlibabaAlihealthNrRxQueryimageRequest() *AlibabaAlihealthNrRxQueryimageAPIRequest {
	return &AlibabaAlihealthNrRxQueryimageAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthNrRxQueryimageAPIRequest) Reset() {
	r._orderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthNrRxQueryimageAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.nr.rx.queryimage"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthNrRxQueryimageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthNrRxQueryimageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 订单编号
func (r *AlibabaAlihealthNrRxQueryimageAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabaAlihealthNrRxQueryimageAPIRequest) GetOrderId() int64 {
	return r._orderId
}

var poolAlibabaAlihealthNrRxQueryimageAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthNrRxQueryimageRequest()
	},
}

// GetAlibabaAlihealthNrRxQueryimageRequest 从 sync.Pool 获取 AlibabaAlihealthNrRxQueryimageAPIRequest
func GetAlibabaAlihealthNrRxQueryimageAPIRequest() *AlibabaAlihealthNrRxQueryimageAPIRequest {
	return poolAlibabaAlihealthNrRxQueryimageAPIRequest.Get().(*AlibabaAlihealthNrRxQueryimageAPIRequest)
}

// ReleaseAlibabaAlihealthNrRxQueryimageAPIRequest 将 AlibabaAlihealthNrRxQueryimageAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthNrRxQueryimageAPIRequest(v *AlibabaAlihealthNrRxQueryimageAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthNrRxQueryimageAPIRequest.Put(v)
}
