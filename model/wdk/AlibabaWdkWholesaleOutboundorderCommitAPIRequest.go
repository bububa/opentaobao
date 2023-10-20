package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkwholesaleoutboundordercommitAPIRequest 盒马帮发货信息回传接口 API请求
// alibaba.wdk.wholesale.outboundorder.commit
//
// 盒马帮发货信息回传接口
type AlibabawdkwholesaleoutboundordercommitAPIRequest struct {
	model.Params
	// 发货信息参数
	_outboundInfoCommitReq *OutboundInfoCommitReq
}

// NewAlibabawdkwholesaleoutboundordercommitRequest 初始化AlibabawdkwholesaleoutboundordercommitAPIRequest对象
func NewAlibabawdkwholesaleoutboundordercommitRequest() *AlibabawdkwholesaleoutboundordercommitAPIRequest {
	return &AlibabawdkwholesaleoutboundordercommitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkwholesaleoutboundordercommitAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.wholesale.outboundorder.commit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkwholesaleoutboundordercommitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkwholesaleoutboundordercommitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOutboundInfoCommitReq is OutboundInfoCommitReq Setter
// 发货信息参数
func (r *AlibabawdkwholesaleoutboundordercommitAPIRequest) SetOutboundInfoCommitReq(_outboundInfoCommitReq *OutboundInfoCommitReq) error {
	r._outboundInfoCommitReq = _outboundInfoCommitReq
	r.Set("outbound_info_commit_req", _outboundInfoCommitReq)
	return nil
}

// GetOutboundInfoCommitReq OutboundInfoCommitReq Getter
func (r AlibabawdkwholesaleoutboundordercommitAPIRequest) GetOutboundInfoCommitReq() *OutboundInfoCommitReq {
	return r._outboundInfoCommitReq
}
