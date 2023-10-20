package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TmallItemSimpleschemaUpdate 天猫简化编辑商品
// tmall.item.simpleschema.update
//
// 国外大商家天猫简化编辑商品
func TmallItemSimpleschemaUpdate(clt *core.SDKClient, req *tbitem.TmallItemSimpleschemaUpdateAPIRequest, resp *tbitem.TmallItemSimpleschemaUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
