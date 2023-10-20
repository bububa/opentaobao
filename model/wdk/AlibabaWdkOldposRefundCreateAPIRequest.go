package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkOldposRefundCreateAPIRequest 五道口外部商户老pos机产生的退款单同步进盒马 API请求
// alibaba.wdk.oldpos.refund.create
//
// 淘鲜达外部商户老pos机产生的退款单同步进淘鲜达
type AlibabaWdkOldposRefundCreateAPIRequest struct {
	model.Params
	// 入参
	_posRefundCreateRequest *PosRefundCreateRequest
}

// NewAlibabaWdkOldposRefundCreateRequest 初始化AlibabaWdkOldposRefundCreateAPIRequest对象
func NewAlibabaWdkOldposRefundCreateRequest() *AlibabaWdkOldposRefundCreateAPIRequest {
	return &AlibabaWdkOldposRefundCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkOldposRefundCreateAPIRequest) Reset() {
	r._posRefundCreateRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkOldposRefundCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.oldpos.refund.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkOldposRefundCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkOldposRefundCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPosRefundCreateRequest is PosRefundCreateRequest Setter
// 入参
func (r *AlibabaWdkOldposRefundCreateAPIRequest) SetPosRefundCreateRequest(_posRefundCreateRequest *PosRefundCreateRequest) error {
	r._posRefundCreateRequest = _posRefundCreateRequest
	r.Set("pos_refund_create_request", _posRefundCreateRequest)
	return nil
}

// GetPosRefundCreateRequest PosRefundCreateRequest Getter
func (r AlibabaWdkOldposRefundCreateAPIRequest) GetPosRefundCreateRequest() *PosRefundCreateRequest {
	return r._posRefundCreateRequest
}

var poolAlibabaWdkOldposRefundCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkOldposRefundCreateRequest()
	},
}

// GetAlibabaWdkOldposRefundCreateRequest 从 sync.Pool 获取 AlibabaWdkOldposRefundCreateAPIRequest
func GetAlibabaWdkOldposRefundCreateAPIRequest() *AlibabaWdkOldposRefundCreateAPIRequest {
	return poolAlibabaWdkOldposRefundCreateAPIRequest.Get().(*AlibabaWdkOldposRefundCreateAPIRequest)
}

// ReleaseAlibabaWdkOldposRefundCreateAPIRequest 将 AlibabaWdkOldposRefundCreateAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkOldposRefundCreateAPIRequest(v *AlibabaWdkOldposRefundCreateAPIRequest) {
	v.Reset()
	poolAlibabaWdkOldposRefundCreateAPIRequest.Put(v)
}
