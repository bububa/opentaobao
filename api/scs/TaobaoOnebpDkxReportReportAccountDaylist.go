package scs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scs"
)

// Taobaoonebpdkxreportreportaccountdaylist 获取账户分日报表
// taobao.onebp.dkx.report.report.account.daylist
//
// 获取账户分日报表
func Taobaoonebpdkxreportreportaccountdaylist(clt *core.SDKClient, req *scs.TaobaoonebpdkxreportreportaccountdaylistAPIRequest, session string) (*scs.TaobaoonebpdkxreportreportaccountdaylistAPIResponse, error) {
	var resp scs.TaobaoonebpdkxreportreportaccountdaylistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
