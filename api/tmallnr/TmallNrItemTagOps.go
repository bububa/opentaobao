package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// Tmallnritemtagops 区域零售商品打标去标
// tmall.nr.item.tag.ops
//
// 参加区域零售的商品，需要批量打标或去标，方便后续设置商品库存
func Tmallnritemtagops(clt *core.SDKClient, req *tmallnr.TmallnritemtagopsAPIRequest, session string) (*tmallnr.TmallnritemtagopsAPIResponse, error) {
	var resp tmallnr.TmallnritemtagopsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
