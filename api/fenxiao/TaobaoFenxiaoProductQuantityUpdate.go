package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoProductQuantityUpdate 产品库存修改
// taobao.fenxiao.product.quantity.update
//
// 修改产品库存信息，支持全量修改以及增量修改两种方式
func TaobaoFenxiaoProductQuantityUpdate(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoProductQuantityUpdateAPIRequest, resp *fenxiao.TaobaoFenxiaoProductQuantityUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
