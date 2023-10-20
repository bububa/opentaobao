package tmallitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallitem"
)

// Tmallitemsextendsearch 搜索天猫商品
// tmall.items.extend.search
//
// 提供天猫商品搜索结果，需要调用精选商品，请改为调用：tmall.selected.items.search
func Tmallitemsextendsearch(clt *core.SDKClient, req *tmallitem.TmallitemsextendsearchAPIRequest, session string) (*tmallitem.TmallitemsextendsearchAPIResponse, error) {
	var resp tmallitem.TmallitemsextendsearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
