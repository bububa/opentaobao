package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkWholesaleInboundorderCommitAPIRequest 盒马帮退货信息回传接口 API请求
// alibaba.wdk.wholesale.inboundorder.commit
//
// 盒马帮退货信息回传接口
type AlibabaWdkWholesaleInboundorderCommitAPIRequest struct {
	model.Params
	// 退货信息参数
	_inboundInfoCommitReq *InboundInfoCommitReq
}

// NewAlibabaWdkWholesaleInboundorderCommitRequest 初始化AlibabaWdkWholesaleInboundorderCommitAPIRequest对象
func NewAlibabaWdkWholesaleInboundorderCommitRequest() *AlibabaWdkWholesaleInboundorderCommitAPIRequest {
	return &AlibabaWdkWholesaleInboundorderCommitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkWholesaleInboundorderCommitAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.wholesale.inboundorder.commit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkWholesaleInboundorderCommitAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is InboundInfoCommitReq Setter
// 退货信息参数
func (r *AlibabaWdkWholesaleInboundorderCommitAPIRequest) SetInboundInfoCommitReq(_inboundInfoCommitReq *InboundInfoCommitReq) error {
	r._inboundInfoCommitReq = _inboundInfoCommitReq
	r.Set("inbound_info_commit_req", _inboundInfoCommitReq)
	return nil
}

// Get InboundInfoCommitReq Getter
func (r AlibabaWdkWholesaleInboundorderCommitAPIRequest) GetInboundInfoCommitReq() *InboundInfoCommitReq {
	return r._inboundInfoCommitReq
}
