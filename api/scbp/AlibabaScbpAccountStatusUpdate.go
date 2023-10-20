package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpaccountstatusupdate 修改账户级别关键词推广状态
// alibaba.scbp.account.status.update
//
// 修改账户级别关键词推广状态
func Alibabascbpaccountstatusupdate(clt *core.SDKClient, req *scbp.AlibabascbpaccountstatusupdateAPIRequest, session string) (*scbp.AlibabascbpaccountstatusupdateAPIResponse, error) {
	var resp scbp.AlibabascbpaccountstatusupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
