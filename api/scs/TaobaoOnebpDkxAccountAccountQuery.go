package scs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scs"
)

// TaobaoOnebpDkxAccountAccountQuery 万相台账号余额查询
// taobao.onebp.dkx.account.account.query
//
// 万相台账号余额查询
func TaobaoOnebpDkxAccountAccountQuery(clt *core.SDKClient, req *scs.TaobaoOnebpDkxAccountAccountQueryAPIRequest, resp *scs.TaobaoOnebpDkxAccountAccountQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
