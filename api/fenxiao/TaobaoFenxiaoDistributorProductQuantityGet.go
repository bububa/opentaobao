package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoDistributorProductQuantityGet 分销商查询产品库存
// taobao.fenxiao.distributor.product.quantity.get
//
// 分销商查询产品库存
func TaobaoFenxiaoDistributorProductQuantityGet(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoDistributorProductQuantityGetAPIRequest, resp *fenxiao.TaobaoFenxiaoDistributorProductQuantityGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
