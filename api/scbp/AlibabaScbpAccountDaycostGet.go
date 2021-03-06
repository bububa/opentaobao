package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAccountDaycostGet 查询今日消耗
// alibaba.scbp.account.daycost.get
//
// 查询今日消耗
func AlibabaScbpAccountDaycostGet(clt *core.SDKClient, req *scbp.AlibabaScbpAccountDaycostGetAPIRequest, session string) (*scbp.AlibabaScbpAccountDaycostGetAPIResponse, error) {
	var resp scbp.AlibabaScbpAccountDaycostGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
