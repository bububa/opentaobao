package scs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scs"
)

// Taobaoonebpdkxaccountaccountquery 万相台账号余额查询
// taobao.onebp.dkx.account.account.query
//
// 万相台账号余额查询
func Taobaoonebpdkxaccountaccountquery(clt *core.SDKClient, req *scs.TaobaoonebpdkxaccountaccountqueryAPIRequest, session string) (*scs.TaobaoonebpdkxaccountaccountqueryAPIResponse, error) {
	var resp scs.TaobaoonebpdkxaccountaccountqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
