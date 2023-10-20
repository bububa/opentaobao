package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkReverseRefundAPIRequest 退款打款 API请求
// alibaba.wdk.reverse.refund
//
// 五道口退款打款开放能力接口
type AlibabaWdkReverseRefundAPIRequest struct {
	model.Params
	// 退款打款请求
	_openRefundReqDTO *OpenRefundReqDto
}

// NewAlibabaWdkReverseRefundRequest 初始化AlibabaWdkReverseRefundAPIRequest对象
func NewAlibabaWdkReverseRefundRequest() *AlibabaWdkReverseRefundAPIRequest {
	return &AlibabaWdkReverseRefundAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkReverseRefundAPIRequest) Reset() {
	r._openRefundReqDTO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkReverseRefundAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.reverse.refund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkReverseRefundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkReverseRefundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOpenRefundReqDTO is OpenRefundReqDTO Setter
// 退款打款请求
func (r *AlibabaWdkReverseRefundAPIRequest) SetOpenRefundReqDTO(_openRefundReqDTO *OpenRefundReqDto) error {
	r._openRefundReqDTO = _openRefundReqDTO
	r.Set("open_refund_req_d_t_o", _openRefundReqDTO)
	return nil
}

// GetOpenRefundReqDTO OpenRefundReqDTO Getter
func (r AlibabaWdkReverseRefundAPIRequest) GetOpenRefundReqDTO() *OpenRefundReqDto {
	return r._openRefundReqDTO
}

var poolAlibabaWdkReverseRefundAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkReverseRefundRequest()
	},
}

// GetAlibabaWdkReverseRefundRequest 从 sync.Pool 获取 AlibabaWdkReverseRefundAPIRequest
func GetAlibabaWdkReverseRefundAPIRequest() *AlibabaWdkReverseRefundAPIRequest {
	return poolAlibabaWdkReverseRefundAPIRequest.Get().(*AlibabaWdkReverseRefundAPIRequest)
}

// ReleaseAlibabaWdkReverseRefundAPIRequest 将 AlibabaWdkReverseRefundAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkReverseRefundAPIRequest(v *AlibabaWdkReverseRefundAPIRequest) {
	v.Reset()
	poolAlibabaWdkReverseRefundAPIRequest.Put(v)
}
