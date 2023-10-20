package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdAccountLevelGet 查询推广账户等级
// alibaba.scbp.ad.account.level.get
//
// 查询推广账户等级
func AlibabaScbpAdAccountLevelGet(clt *core.SDKClient, req *scbp.AlibabaScbpAdAccountLevelGetAPIRequest, resp *scbp.AlibabaScbpAdAccountLevelGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
