package lstwarehouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstTradeSellerWarehouseQueryAPIRequest
供应商-本云商家-仓库查询接口 API请求
alibaba.lst.trade.seller.warehouse.query

查询本地云仓商家的仓库 */
type AlibabaLstTradeSellerWarehouseQueryAPIRequest struct {
	model.Params
	// 入参
	_warehouseQueryParam *WarehouseQueryParam
}

// New
