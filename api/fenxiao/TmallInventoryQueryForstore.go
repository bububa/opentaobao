package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TmallInventoryQueryForstore 查询后端商品仓库库存
// tmall.inventory.query.forstore
//
// 商家查询后端商品仓库库存
func TmallInventoryQueryForstore(clt *core.SDKClient, req *fenxiao.TmallInventoryQueryForstoreAPIRequest, resp *fenxiao.TmallInventoryQueryForstoreAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
