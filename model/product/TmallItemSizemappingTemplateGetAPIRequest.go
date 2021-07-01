package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallItemSizemappingTemplateGetAPIRequest
获取天猫商品尺码表模板 API请求
tmall.item.sizemapping.template.get

获取天猫商品尺码表模板 */
type TmallItemSizemappingTemplateGetAPIRequest struct {
	model.Params
	// 尺码表模板ID
	_templateId int64
}

// New
