package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoRegionSaleQuery 查询商品销售区域
// taobao.region.sale.query
//
// 查询商品销售区域
func TaobaoRegionSaleQuery(clt *core.SDKClient, req *fenxiao.TaobaoRegionSaleQueryAPIRequest, resp *fenxiao.TaobaoRegionSaleQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
