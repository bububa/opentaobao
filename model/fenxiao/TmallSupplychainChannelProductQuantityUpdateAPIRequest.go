package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallSupplychainChannelProductQuantityUpdateAPIRequest
渠道无仓库存更新接口 API请求
tmall.supplychain.channel.product.quantity.update

渠道无仓库存更新接口 */
type TmallSupplychainChannelProductQuantityUpdateAPIRequest struct {
	model.Params
	// 库存修改值，必选。当全量更新库存时，quantity必须为大于等于0的正整数；当增量更新库存时，quantity为整数，可小于等于0。若增量更新时传入的库存为负数，则负数与实际库存之和不能小于0。比如当前实际库存为1，传入增量更新quantity=-1，库存改为0
	_quantity int64
	// 产品数字ID
	_productId int64
	// 产品SKU ID
	_skuId int64
	// 库存更新方式，可选。1为全量更新，2为增量更新。如果不填，默认为全量更新
	_updateType int64
}

// New
