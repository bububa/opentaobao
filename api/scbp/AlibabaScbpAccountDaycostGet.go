package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpaccountdaycostget 查询今日消耗
// alibaba.scbp.account.daycost.get
//
// 查询今日消耗
func Alibabascbpaccountdaycostget(clt *core.SDKClient, req *scbp.AlibabascbpaccountdaycostgetAPIRequest, session string) (*scbp.AlibabascbpaccountdaycostgetAPIResponse, error) {
	var resp scbp.AlibabascbpaccountdaycostgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
