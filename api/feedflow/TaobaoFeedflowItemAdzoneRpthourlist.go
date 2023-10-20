package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// Taobaofeedflowitemadzonerpthourlist 超级推荐【商品推广】资源位分时报表查询
// taobao.feedflow.item.adzone.rpthourlist
//
// 广告主资源包分时数据查询，支持广告主查询最近90天内某一天的资源包维度分时报表数据
func Taobaofeedflowitemadzonerpthourlist(clt *core.SDKClient, req *feedflow.TaobaofeedflowitemadzonerpthourlistAPIRequest, session string) (*feedflow.TaobaofeedflowitemadzonerpthourlistAPIResponse, error) {
	var resp feedflow.TaobaofeedflowitemadzonerpthourlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
