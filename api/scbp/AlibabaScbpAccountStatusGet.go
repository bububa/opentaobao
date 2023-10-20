package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpaccountstatusget 查询账户级别关键词推广状态
// alibaba.scbp.account.status.get
//
// 查询账户级别关键词推广状态
func Alibabascbpaccountstatusget(clt *core.SDKClient, req *scbp.AlibabascbpaccountstatusgetAPIRequest, session string) (*scbp.AlibabascbpaccountstatusgetAPIResponse, error) {
	var resp scbp.AlibabascbpaccountstatusgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
