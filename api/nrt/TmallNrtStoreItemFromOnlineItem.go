package nrt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// TmallNrtStoreItemFromOnlineItem 基于新模型商品id查询摊位子品id
// tmall.nrt.store.item.from.online.item
//
// 基于新模型商品id查询摊位子品id
func TmallNrtStoreItemFromOnlineItem(clt *core.SDKClient, req *nrt.TmallNrtStoreItemFromOnlineItemAPIRequest, resp *nrt.TmallNrtStoreItemFromOnlineItemAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
