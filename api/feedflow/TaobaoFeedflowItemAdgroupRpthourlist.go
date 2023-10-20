package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// Taobaofeedflowitemadgrouprpthourlist 超级推荐【商品推广】单元分时报表查询
// taobao.feedflow.item.adgroup.rpthourlist
//
// 广告主推广组分时数据查询，支持广告主查询最近90天内某一天的单元维度分时报表数据
func Taobaofeedflowitemadgrouprpthourlist(clt *core.SDKClient, req *feedflow.TaobaofeedflowitemadgrouprpthourlistAPIRequest, session string) (*feedflow.TaobaofeedflowitemadgrouprpthourlistAPIResponse, error) {
	var resp feedflow.TaobaofeedflowitemadgrouprpthourlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
