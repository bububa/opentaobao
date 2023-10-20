package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAccountDaycostGet 查询今日消耗
// alibaba.scbp.account.daycost.get
//
// 查询今日消耗
func AlibabaScbpAccountDaycostGet(clt *core.SDKClient, req *scbp.AlibabaScbpAccountDaycostGetAPIRequest, resp *scbp.AlibabaScbpAccountDaycostGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
