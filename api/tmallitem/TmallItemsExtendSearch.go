package tmallitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallitem"
)

// TmallItemsExtendSearch 搜索天猫商品
// tmall.items.extend.search
//
// 提供天猫商品搜索结果，需要调用精选商品，请改为调用：tmall.selected.items.search
func TmallItemsExtendSearch(clt *core.SDKClient, req *tmallitem.TmallItemsExtendSearchAPIRequest, resp *tmallitem.TmallItemsExtendSearchAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
