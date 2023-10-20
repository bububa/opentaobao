package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkReverseReversedetailAPIRequest 退款详情 API请求
// alibaba.wdk.reverse.reversedetail
//
// 退款详情
type AlibabaWdkReverseReversedetailAPIRequest struct {
	model.Params
	// 退款单id
	_reverseId string
}

// NewAlibabaWdkReverseReversedetailRequest 初始化AlibabaWdkReverseReversedetailAPIRequest对象
func NewAlibabaWdkReverseReversedetailRequest() *AlibabaWdkReverseReversedetailAPIRequest {
	return &AlibabaWdkReverseReversedetailAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkReverseReversedetailAPIRequest) Reset() {
	r._reverseId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkReverseReversedetailAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.reverse.reversedetail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkReverseReversedetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkReverseReversedetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReverseId is ReverseId Setter
// 退款单id
func (r *AlibabaWdkReverseReversedetailAPIRequest) SetReverseId(_reverseId string) error {
	r._reverseId = _reverseId
	r.Set("reverse_id", _reverseId)
	return nil
}

// GetReverseId ReverseId Getter
func (r AlibabaWdkReverseReversedetailAPIRequest) GetReverseId() string {
	return r._reverseId
}

var poolAlibabaWdkReverseReversedetailAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkReverseReversedetailRequest()
	},
}

// GetAlibabaWdkReverseReversedetailRequest 从 sync.Pool 获取 AlibabaWdkReverseReversedetailAPIRequest
func GetAlibabaWdkReverseReversedetailAPIRequest() *AlibabaWdkReverseReversedetailAPIRequest {
	return poolAlibabaWdkReverseReversedetailAPIRequest.Get().(*AlibabaWdkReverseReversedetailAPIRequest)
}

// ReleaseAlibabaWdkReverseReversedetailAPIRequest 将 AlibabaWdkReverseReversedetailAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkReverseReversedetailAPIRequest(v *AlibabaWdkReverseReversedetailAPIRequest) {
	v.Reset()
	poolAlibabaWdkReverseReversedetailAPIRequest.Put(v)
}
