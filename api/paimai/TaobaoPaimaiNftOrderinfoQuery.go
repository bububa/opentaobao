package paimai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/paimai"
)

// Taobaopaimainftorderinfoquery 查询订单类型
// taobao.paimai.nft.orderinfo.query
//
// 查询订单类型
func Taobaopaimainftorderinfoquery(clt *core.SDKClient, req *paimai.TaobaopaimainftorderinfoqueryAPIRequest, session string) (*paimai.TaobaopaimainftorderinfoqueryAPIResponse, error) {
	var resp paimai.TaobaopaimainftorderinfoqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
