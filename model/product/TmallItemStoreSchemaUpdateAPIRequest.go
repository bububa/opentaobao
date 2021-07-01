package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallItemStoreSchemaUpdateAPIRequest
天猫门店商品编辑 API请求
tmall.item.store.schema.update

天猫门店商品编辑 */
type TmallItemStoreSchemaUpdateAPIRequest struct {
	model.Params
	// 主商品ID
	_mainItemId int64
	// 门店ID
	_storeId int64
	// 商品的schema xml
	_xml string
}

// New
