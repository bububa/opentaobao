package paimai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/paimai"
)

// TaobaoPaimaiNftOrderinfoQuery 查询订单类型
// taobao.paimai.nft.orderinfo.query
//
// 查询订单类型
func TaobaoPaimaiNftOrderinfoQuery(clt *core.SDKClient, req *paimai.TaobaoPaimaiNftOrderinfoQueryAPIRequest, session string) (*paimai.TaobaoPaimaiNftOrderinfoQueryAPIResponse, error) {
	var resp paimai.TaobaoPaimaiNftOrderinfoQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
