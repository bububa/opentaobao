package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// TmallNrItemTagOps 区域零售商品打标去标
// tmall.nr.item.tag.ops
//
// 参加区域零售的商品，需要批量打标或去标，方便后续设置商品库存
func TmallNrItemTagOps(clt *core.SDKClient, req *tmallnr.TmallNrItemTagOpsAPIRequest, resp *tmallnr.TmallNrItemTagOpsAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
