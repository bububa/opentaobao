package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoRegionWarehouseQueryAPIRequest
查询仓库覆盖范围 API请求
taobao.region.warehouse.query

查询仓库覆盖范围 */
type TaobaoRegionWarehouseQueryAPIRequest struct {
	model.Params
	// 仓库编码
	_storeCode string
}

// New
