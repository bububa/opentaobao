package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallItemStoreUpdateSchemaGetAPIRequest
天猫门店商品修改规则获取 API请求
tmall.item.store.update.schema.get

天猫门店商品修改规则获取 */
type TmallItemStoreUpdateSchemaGetAPIRequest struct {
	model.Params
	// 主商品ID
	_mainItemId int64
	// 门店ID
	_storeId int64
}

// New
