package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAccountIsarrearsGet 查询关键词推广账户是否欠款
// alibaba.scbp.account.isarrears.get
//
// 查询关键词推广账户是否欠款
func AlibabaScbpAccountIsarrearsGet(clt *core.SDKClient, req *scbp.AlibabaScbpAccountIsarrearsGetAPIRequest, resp *scbp.AlibabaScbpAccountIsarrearsGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
