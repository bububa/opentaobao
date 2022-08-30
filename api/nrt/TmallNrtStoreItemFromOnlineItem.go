package nrt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// TmallNrtStoreItemFromOnlineItem 基于新模型商品id查询摊位子品id
// tmall.nrt.store.item.from.online.item
//
// 基于新模型商品id查询摊位子品id
func TmallNrtStoreItemFromOnlineItem(clt *core.SDKClient, req *nrt.TmallNrtStoreItemFromOnlineItemAPIRequest, session string) (*nrt.TmallNrtStoreItemFromOnlineItemAPIResponse, error) {
	var resp nrt.TmallNrtStoreItemFromOnlineItemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
