package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbahourreportaccountget 账户级别小时报表获取
// taobao.simba.hour.report.account.get
//
// 获取账户小时实时报表数据
func Taobaosimbahourreportaccountget(clt *core.SDKClient, req *simba.TaobaosimbahourreportaccountgetAPIRequest, session string) (*simba.TaobaosimbahourreportaccountgetAPIResponse, error) {
	var resp simba.TaobaosimbahourreportaccountgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
