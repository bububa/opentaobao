package lstwarehouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstIcStockItemsUpdateAPIRequest
零售通经销商商品库存设置 API请求
alibaba.lst.ic.stock.items.update

零售通经销商商品库存设置 */
type AlibabaLstIcStockItemsUpdateAPIRequest struct {
	model.Params
	// 零售通经销商商品库存
	_query *LstItemStockParam
}

// New
