package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Taobaoregionsalequery 查询商品销售区域
// taobao.region.sale.query
//
// 查询商品销售区域
func Taobaoregionsalequery(clt *core.SDKClient, req *fenxiao.TaobaoregionsalequeryAPIRequest, session string) (*fenxiao.TaobaoregionsalequeryAPIResponse, error) {
	var resp fenxiao.TaobaoregionsalequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
