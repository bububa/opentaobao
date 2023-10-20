package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// Taobaofeedflowitemcrowdrpthourlist 超级推荐【商品推广】定向分时报表查询
// taobao.feedflow.item.crowd.rpthourlist
//
// 广告主定向分时数据查询，支持广告主查询最近90天内某一天的定向维度分时报表数据
func Taobaofeedflowitemcrowdrpthourlist(clt *core.SDKClient, req *feedflow.TaobaofeedflowitemcrowdrpthourlistAPIRequest, session string) (*feedflow.TaobaofeedflowitemcrowdrpthourlistAPIResponse, error) {
	var resp feedflow.TaobaofeedflowitemcrowdrpthourlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
