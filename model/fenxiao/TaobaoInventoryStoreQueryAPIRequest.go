package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoInventoryStoreQueryAPIRequest
查询仓库信息 API请求
taobao.inventory.store.query

查询商家仓信息 */
type TaobaoInventoryStoreQueryAPIRequest struct {
	model.Params
	// 商家的仓库编码
	_storeCode string
}

// New
