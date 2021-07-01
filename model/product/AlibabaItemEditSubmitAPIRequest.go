package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaItemEditSubmitAPIRequest
商品编辑提交schema信息 API请求
alibaba.item.edit.submit

商品编辑提交schema信息 */
type AlibabaItemEditSubmitAPIRequest struct {
	model.Params
	// 业务扩展参数，需与平台约定好
	_bizType string
	// 商品类目ID。若不需要修改商品类目，则不用填写
	_catId int64
	// 产品ID，若不需要修改关联的产品信息，则不需要填写
	_spuId int64
	// 商品ID
	_itemId int64
	// 编辑后的schema信息，通过alibaba.item.edit.schema.get获取
	_schema string
}

// New
