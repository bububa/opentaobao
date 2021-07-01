package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkWholesaleInboundorderCommitAPIRequest
盒马帮退货信息回传接口 API请求
alibaba.wdk.wholesale.inboundorder.commit

盒马帮退货信息回传接口 */
type AlibabaWdkWholesaleInboundorderCommitAPIRequest struct {
	model.Params
	// 退货信息参数
	_inboundInfoCommitReq *InboundInfoCommitReq
}

// New
