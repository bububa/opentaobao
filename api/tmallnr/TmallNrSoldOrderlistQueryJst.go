package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// Tmallnrsoldorderlistqueryjst 已入塔商家查询订单列表
// tmall.nr.sold.orderlist.query.jst
//
// 该服务用于已入聚石塔的商家，获取订单列表信息；
func Tmallnrsoldorderlistqueryjst(clt *core.SDKClient, req *tmallnr.TmallnrsoldorderlistqueryjstAPIRequest, session string) (*tmallnr.TmallnrsoldorderlistqueryjstAPIResponse, error) {
	var resp tmallnr.TmallnrsoldorderlistqueryjstAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
