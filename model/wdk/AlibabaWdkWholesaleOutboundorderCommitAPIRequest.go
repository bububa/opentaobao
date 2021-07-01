package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkWholesaleOutboundorderCommitAPIRequest
盒马帮发货信息回传接口 API请求
alibaba.wdk.wholesale.outboundorder.commit

盒马帮发货信息回传接口 */
type AlibabaWdkWholesaleOutboundorderCommitAPIRequest struct {
	model.Params
	// 发货信息参数
	_outboundInfoCommitReq *OutboundInfoCommitReq
}

// New
