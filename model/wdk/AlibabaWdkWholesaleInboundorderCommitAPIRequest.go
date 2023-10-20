package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkwholesaleinboundordercommitAPIRequest 盒马帮退货信息回传接口 API请求
// alibaba.wdk.wholesale.inboundorder.commit
//
// 盒马帮退货信息回传接口
type AlibabawdkwholesaleinboundordercommitAPIRequest struct {
	model.Params
	// 退货信息参数
	_inboundInfoCommitReq *InboundInfoCommitReq
}

// NewAlibabawdkwholesaleinboundordercommitRequest 初始化AlibabawdkwholesaleinboundordercommitAPIRequest对象
func NewAlibabawdkwholesaleinboundordercommitRequest() *AlibabawdkwholesaleinboundordercommitAPIRequest {
	return &AlibabawdkwholesaleinboundordercommitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkwholesaleinboundordercommitAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.wholesale.inboundorder.commit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkwholesaleinboundordercommitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkwholesaleinboundordercommitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInboundInfoCommitReq is InboundInfoCommitReq Setter
// 退货信息参数
func (r *AlibabawdkwholesaleinboundordercommitAPIRequest) SetInboundInfoCommitReq(_inboundInfoCommitReq *InboundInfoCommitReq) error {
	r._inboundInfoCommitReq = _inboundInfoCommitReq
	r.Set("inbound_info_commit_req", _inboundInfoCommitReq)
	return nil
}

// GetInboundInfoCommitReq InboundInfoCommitReq Getter
func (r AlibabawdkwholesaleinboundordercommitAPIRequest) GetInboundInfoCommitReq() *InboundInfoCommitReq {
	return r._inboundInfoCommitReq
}
