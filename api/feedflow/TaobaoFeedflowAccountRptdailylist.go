package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// TaobaoFeedflowAccountRptdailylist 获取广告主分日数据
// taobao.feedflow.account.rptdailylist
//
// 获取广告主分日数据
func TaobaoFeedflowAccountRptdailylist(clt *core.SDKClient, req *feedflow.TaobaoFeedflowAccountRptdailylistAPIRequest, session string) (*feedflow.TaobaoFeedflowAccountRptdailylistAPIResponse, error) {
	var resp feedflow.TaobaoFeedflowAccountRptdailylistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
