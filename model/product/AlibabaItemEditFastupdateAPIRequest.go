package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaItemEditFastupdateAPIRequest
商品编辑增量更新 API请求
alibaba.item.edit.fastupdate

商品编辑增量更新;
<br/>该接口编辑sku，只能更新价格、库存等信息，不能新增sku;
<br/>新增sku用全量接口alibaba.item.edit.submit，先设置销售属性; */
type AlibabaItemEditFastupdateAPIRequest struct {
	model.Params
	// 商品类目ID。若不需要修改商品类目，则不用填写
	_catId int64
	// 产品ID，若不需要修改关联的产品信息，则不需要填写
	_spuId int64
	// 商品ID
	_itemId int64
	// 编辑后的schema信息(增量更新，只填写需要更新的字段)
	_schema string
}

// New
