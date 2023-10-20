package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkWholesaleOutboundorderCommitAPIRequest 盒马帮发货信息回传接口 API请求
// alibaba.wdk.wholesale.outboundorder.commit
//
// 盒马帮发货信息回传接口
type AlibabaWdkWholesaleOutboundorderCommitAPIRequest struct {
	model.Params
	// 发货信息参数
	_outboundInfoCommitReq *OutboundInfoCommitReq
}

// NewAlibabaWdkWholesaleOutboundorderCommitRequest 初始化AlibabaWdkWholesaleOutboundorderCommitAPIRequest对象
func NewAlibabaWdkWholesaleOutboundorderCommitRequest() *AlibabaWdkWholesaleOutboundorderCommitAPIRequest {
	return &AlibabaWdkWholesaleOutboundorderCommitAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkWholesaleOutboundorderCommitAPIRequest) Reset() {
	r._outboundInfoCommitReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkWholesaleOutboundorderCommitAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.wholesale.outboundorder.commit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkWholesaleOutboundorderCommitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkWholesaleOutboundorderCommitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOutboundInfoCommitReq is OutboundInfoCommitReq Setter
// 发货信息参数
func (r *AlibabaWdkWholesaleOutboundorderCommitAPIRequest) SetOutboundInfoCommitReq(_outboundInfoCommitReq *OutboundInfoCommitReq) error {
	r._outboundInfoCommitReq = _outboundInfoCommitReq
	r.Set("outbound_info_commit_req", _outboundInfoCommitReq)
	return nil
}

// GetOutboundInfoCommitReq OutboundInfoCommitReq Getter
func (r AlibabaWdkWholesaleOutboundorderCommitAPIRequest) GetOutboundInfoCommitReq() *OutboundInfoCommitReq {
	return r._outboundInfoCommitReq
}

var poolAlibabaWdkWholesaleOutboundorderCommitAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkWholesaleOutboundorderCommitRequest()
	},
}

// GetAlibabaWdkWholesaleOutboundorderCommitRequest 从 sync.Pool 获取 AlibabaWdkWholesaleOutboundorderCommitAPIRequest
func GetAlibabaWdkWholesaleOutboundorderCommitAPIRequest() *AlibabaWdkWholesaleOutboundorderCommitAPIRequest {
	return poolAlibabaWdkWholesaleOutboundorderCommitAPIRequest.Get().(*AlibabaWdkWholesaleOutboundorderCommitAPIRequest)
}

// ReleaseAlibabaWdkWholesaleOutboundorderCommitAPIRequest 将 AlibabaWdkWholesaleOutboundorderCommitAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkWholesaleOutboundorderCommitAPIRequest(v *AlibabaWdkWholesaleOutboundorderCommitAPIRequest) {
	v.Reset()
	poolAlibabaWdkWholesaleOutboundorderCommitAPIRequest.Put(v)
}
