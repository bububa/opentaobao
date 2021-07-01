package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoRegionWarehouseManageAPIRequest
编辑仓库覆盖范围 API请求
taobao.region.warehouse.manage

编辑仓库覆盖范围 */
type TaobaoRegionWarehouseManageAPIRequest struct {
	model.Params
	// 仓库编码
	_storeCode string
	// 可映射三级地址,例: 广东省
	_regions []string
}

// New
