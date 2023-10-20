package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAccountStatusUpdate 修改账户级别关键词推广状态
// alibaba.scbp.account.status.update
//
// 修改账户级别关键词推广状态
func AlibabaScbpAccountStatusUpdate(clt *core.SDKClient, req *scbp.AlibabaScbpAccountStatusUpdateAPIRequest, resp *scbp.AlibabaScbpAccountStatusUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
