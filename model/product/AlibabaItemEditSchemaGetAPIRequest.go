package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaItemEditSchemaGetAPIRequest
商品编辑获取schema信息 API请求
alibaba.item.edit.schema.get

商品编辑时，获取商品规则信息 */
type AlibabaItemEditSchemaGetAPIRequest struct {
	model.Params
	// 业务扩展参数，需与平台约定好
	_bizType string
	// 商品ID
	_itemId int64
	// 制定返回schema中field字段列表，可用于裁剪返回的schema信息。不填则为全部field
	_fields []string
}

// New
