package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// Taobaofeedflowaccountget 获取信息流账户详情
// taobao.feedflow.account.get
//
// 获取账户信息接口。
// (1) BP显示余额 (字段 ：banlance ) = 现金余额(字段：cash_balance) + 赠款余额；
// (2) 可用余额(字段：availableBalance) = BP显示余额
// (3) 红包(字段：redPacket)
func Taobaofeedflowaccountget(clt *core.SDKClient, req *feedflow.TaobaofeedflowaccountgetAPIRequest, session string) (*feedflow.TaobaofeedflowaccountgetAPIResponse, error) {
	var resp feedflow.TaobaofeedflowaccountgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
