package shenjing

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIbShenjingVisitorPadFetchcodeverifyAPIRequest 访客通过PAD提交访客码 API请求
// alibaba.ib.shenjing.visitor.pad.fetchcodeverify
//
// 访客通过PAD提交访客码，录脸进入园区。
type AlibabaIbShenjingVisitorPadFetchcodeverifyAPIRequest struct {
	model.Params
	// 终端ID
	_termId string
	// 访客码
	_visitorCode int64
}

// NewAlibabaIbShenjingVisitorPadFetchcodeverifyRequest 初始化AlibabaIbShenjingVisitorPadFetchcodeverifyAPIRequest对象
func NewAlibabaIbShenjingVisitorPadFetchcodeverifyRequest() *AlibabaIbShenjingVisitorPadFetchcodeverifyAPIRequest {
	return &AlibabaIbShenjingVisitorPadFetchcodeverifyAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIbShenjingVisitorPadFetchcodeverifyAPIRequest) Reset() {
	r._termId = ""
	r._visitorCode = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIbShenjingVisitorPadFetchcodeverifyAPIRequest) GetApiMethodName() string {
	return "alibaba.ib.shenjing.visitor.pad.fetchcodeverify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIbShenjingVisitorPadFetchcodeverifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIbShenjingVisitorPadFetchcodeverifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTermId is TermId Setter
// 终端ID
func (r *AlibabaIbShenjingVisitorPadFetchcodeverifyAPIRequest) SetTermId(_termId string) error {
	r._termId = _termId
	r.Set("term_id", _termId)
	return nil
}

// GetTermId TermId Getter
func (r AlibabaIbShenjingVisitorPadFetchcodeverifyAPIRequest) GetTermId() string {
	return r._termId
}

// SetVisitorCode is VisitorCode Setter
// 访客码
func (r *AlibabaIbShenjingVisitorPadFetchcodeverifyAPIRequest) SetVisitorCode(_visitorCode int64) error {
	r._visitorCode = _visitorCode
	r.Set("visitor_code", _visitorCode)
	return nil
}

// GetVisitorCode VisitorCode Getter
func (r AlibabaIbShenjingVisitorPadFetchcodeverifyAPIRequest) GetVisitorCode() int64 {
	return r._visitorCode
}

var poolAlibabaIbShenjingVisitorPadFetchcodeverifyAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIbShenjingVisitorPadFetchcodeverifyRequest()
	},
}

// GetAlibabaIbShenjingVisitorPadFetchcodeverifyRequest 从 sync.Pool 获取 AlibabaIbShenjingVisitorPadFetchcodeverifyAPIRequest
func GetAlibabaIbShenjingVisitorPadFetchcodeverifyAPIRequest() *AlibabaIbShenjingVisitorPadFetchcodeverifyAPIRequest {
	return poolAlibabaIbShenjingVisitorPadFetchcodeverifyAPIRequest.Get().(*AlibabaIbShenjingVisitorPadFetchcodeverifyAPIRequest)
}

// ReleaseAlibabaIbShenjingVisitorPadFetchcodeverifyAPIRequest 将 AlibabaIbShenjingVisitorPadFetchcodeverifyAPIRequest 放入 sync.Pool
func ReleaseAlibabaIbShenjingVisitorPadFetchcodeverifyAPIRequest(v *AlibabaIbShenjingVisitorPadFetchcodeverifyAPIRequest) {
	v.Reset()
	poolAlibabaIbShenjingVisitorPadFetchcodeverifyAPIRequest.Put(v)
}
