package scs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scs"
)

// TaobaoOnebpDkxAccountAccountQuery 万相台账号余额查询
// taobao.onebp.dkx.account.account.query
//
// 万相台账号余额查询
func TaobaoOnebpDkxAccountAccountQuery(clt *core.SDKClient, req *scs.TaobaoOnebpDkxAccountAccountQueryAPIRequest, session string) (*scs.TaobaoOnebpDkxAccountAccountQueryAPIResponse, error) {
	var resp scs.TaobaoOnebpDkxAccountAccountQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
