package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// Taobaofeedflowaccountrptdailylist 获取广告主分日数据
// taobao.feedflow.account.rptdailylist
//
// 获取广告主分日数据
func Taobaofeedflowaccountrptdailylist(clt *core.SDKClient, req *feedflow.TaobaofeedflowaccountrptdailylistAPIRequest, session string) (*feedflow.TaobaofeedflowaccountrptdailylistAPIResponse, error) {
	var resp feedflow.TaobaofeedflowaccountrptdailylistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
