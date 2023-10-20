package ascpqcc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpQccSampleCancelItemRelationAPIRequest 魅力惠样品解除父子商品关系 API请求
// alibaba.ascp.qcc.sample.cancel.item.relation
//
// 品控中心魅力惠样品解除父子商品关系
type AlibabaAscpQccSampleCancelItemRelationAPIRequest struct {
	model.Params
	// 请求参数对象
	_cancelRequest *CancelSampleRelationRequest
}

// NewAlibabaAscpQccSampleCancelItemRelationRequest 初始化AlibabaAscpQccSampleCancelItemRelationAPIRequest对象
func NewAlibabaAscpQccSampleCancelItemRelationRequest() *AlibabaAscpQccSampleCancelItemRelationAPIRequest {
	return &AlibabaAscpQccSampleCancelItemRelationAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscpQccSampleCancelItemRelationAPIRequest) Reset() {
	r._cancelRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpQccSampleCancelItemRelationAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.qcc.sample.cancel.item.relation"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpQccSampleCancelItemRelationAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpQccSampleCancelItemRelationAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCancelRequest is CancelRequest Setter
// 请求参数对象
func (r *AlibabaAscpQccSampleCancelItemRelationAPIRequest) SetCancelRequest(_cancelRequest *CancelSampleRelationRequest) error {
	r._cancelRequest = _cancelRequest
	r.Set("cancel_request", _cancelRequest)
	return nil
}

// GetCancelRequest CancelRequest Getter
func (r AlibabaAscpQccSampleCancelItemRelationAPIRequest) GetCancelRequest() *CancelSampleRelationRequest {
	return r._cancelRequest
}

var poolAlibabaAscpQccSampleCancelItemRelationAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscpQccSampleCancelItemRelationRequest()
	},
}

// GetAlibabaAscpQccSampleCancelItemRelationRequest 从 sync.Pool 获取 AlibabaAscpQccSampleCancelItemRelationAPIRequest
func GetAlibabaAscpQccSampleCancelItemRelationAPIRequest() *AlibabaAscpQccSampleCancelItemRelationAPIRequest {
	return poolAlibabaAscpQccSampleCancelItemRelationAPIRequest.Get().(*AlibabaAscpQccSampleCancelItemRelationAPIRequest)
}

// ReleaseAlibabaAscpQccSampleCancelItemRelationAPIRequest 将 AlibabaAscpQccSampleCancelItemRelationAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscpQccSampleCancelItemRelationAPIRequest(v *AlibabaAscpQccSampleCancelItemRelationAPIRequest) {
	v.Reset()
	poolAlibabaAscpQccSampleCancelItemRelationAPIRequest.Put(v)
}
