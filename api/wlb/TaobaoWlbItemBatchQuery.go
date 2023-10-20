package wlb

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

// TaobaoWlbItemBatchQuery 批次库存查询接口
// taobao.wlb.item.batch.query
//
// 根据用户id，item id list和store code来查询商品库存信息和批次信息
func TaobaoWlbItemBatchQuery(clt *core.SDKClient, req *wlb.TaobaoWlbItemBatchQueryAPIRequest, resp *wlb.TaobaoWlbItemBatchQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
