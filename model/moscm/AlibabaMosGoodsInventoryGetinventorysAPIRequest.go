package moscm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMosGoodsInventoryGetinventorysAPIRequest
可售库存查询 API请求
alibaba.mos.goods.inventory.getinventorys

查询商品的可售、在库和占库数量 */
type AlibabaMosGoodsInventoryGetinventorysAPIRequest struct {
	model.Params
	// 查询对象
	_paramVirtualInventoryQueryDto *VirtualInventoryQueryDto
}

// New
