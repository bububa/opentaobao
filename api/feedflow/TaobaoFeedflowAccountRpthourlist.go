package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// TaobaoFeedflowAccountRpthourlist 超级推荐广告主分时报表查询
// taobao.feedflow.account.rpthourlist
//
// 广告主分时报表查询，支持广告主查询最近90天内某一天的账户维度分时报表数据
func TaobaoFeedflowAccountRpthourlist(clt *core.SDKClient, req *feedflow.TaobaoFeedflowAccountRpthourlistAPIRequest, session string) (*feedflow.TaobaoFeedflowAccountRpthourlistAPIResponse, error) {
	var resp feedflow.TaobaoFeedflowAccountRpthourlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
