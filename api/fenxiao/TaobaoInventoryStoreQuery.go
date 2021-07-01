package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

/* TaobaoInventoryStoreQuery
查询仓库信息
taobao.inventory.store.query

查询商家仓信息 */
func TaobaoInventoryStoreQuery(clt *core.SDKClient, req *fenxiao.TaobaoInventoryStoreQueryAPIRequest, session string) (*fenxiao.TaobaoInventoryStoreQueryAPIResponse, error) {
	var resp fenxiao.TaobaoInventoryStoreQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
