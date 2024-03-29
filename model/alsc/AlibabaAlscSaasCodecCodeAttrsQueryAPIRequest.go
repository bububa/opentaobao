package alsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscSaasCodecCodeAttrsQueryAPIRequest 码业务属性查询 API请求
// alibaba.alsc.saas.codec.code.attrs.query
//
// 码业务属性查询
type AlibabaAlscSaasCodecCodeAttrsQueryAPIRequest struct {
	model.Params
	// 请求入参
	_queryCodeRequest *QueryCodeBizAttrRequest
}

// NewAlibabaAlscSaasCodecCodeAttrsQueryRequest 初始化AlibabaAlscSaasCodecCodeAttrsQueryAPIRequest对象
func NewAlibabaAlscSaasCodecCodeAttrsQueryRequest() *AlibabaAlscSaasCodecCodeAttrsQueryAPIRequest {
	return &AlibabaAlscSaasCodecCodeAttrsQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscSaasCodecCodeAttrsQueryAPIRequest) Reset() {
	r._queryCodeRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscSaasCodecCodeAttrsQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.saas.codec.code.attrs.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscSaasCodecCodeAttrsQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscSaasCodecCodeAttrsQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryCodeRequest is QueryCodeRequest Setter
// 请求入参
func (r *AlibabaAlscSaasCodecCodeAttrsQueryAPIRequest) SetQueryCodeRequest(_queryCodeRequest *QueryCodeBizAttrRequest) error {
	r._queryCodeRequest = _queryCodeRequest
	r.Set("query_code_request", _queryCodeRequest)
	return nil
}

// GetQueryCodeRequest QueryCodeRequest Getter
func (r AlibabaAlscSaasCodecCodeAttrsQueryAPIRequest) GetQueryCodeRequest() *QueryCodeBizAttrRequest {
	return r._queryCodeRequest
}

var poolAlibabaAlscSaasCodecCodeAttrsQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscSaasCodecCodeAttrsQueryRequest()
	},
}

// GetAlibabaAlscSaasCodecCodeAttrsQueryRequest 从 sync.Pool 获取 AlibabaAlscSaasCodecCodeAttrsQueryAPIRequest
func GetAlibabaAlscSaasCodecCodeAttrsQueryAPIRequest() *AlibabaAlscSaasCodecCodeAttrsQueryAPIRequest {
	return poolAlibabaAlscSaasCodecCodeAttrsQueryAPIRequest.Get().(*AlibabaAlscSaasCodecCodeAttrsQueryAPIRequest)
}

// ReleaseAlibabaAlscSaasCodecCodeAttrsQueryAPIRequest 将 AlibabaAlscSaasCodecCodeAttrsQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscSaasCodecCodeAttrsQueryAPIRequest(v *AlibabaAlscSaasCodecCodeAttrsQueryAPIRequest) {
	v.Reset()
	poolAlibabaAlscSaasCodecCodeAttrsQueryAPIRequest.Put(v)
}
