package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpaccountisarrearsget 查询关键词推广账户是否欠款
// alibaba.scbp.account.isarrears.get
//
// 查询关键词推广账户是否欠款
func Alibabascbpaccountisarrearsget(clt *core.SDKClient, req *scbp.AlibabascbpaccountisarrearsgetAPIRequest, session string) (*scbp.AlibabascbpaccountisarrearsgetAPIResponse, error) {
	var resp scbp.AlibabascbpaccountisarrearsgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
