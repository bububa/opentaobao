package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallItemSizemappingTemplateDeleteAPIRequest
删除天猫商品尺码表模板 API请求
tmall.item.sizemapping.template.delete

删除天猫商品尺码表模板 */
type TmallItemSizemappingTemplateDeleteAPIRequest struct {
	model.Params
	// 尺码表模板ID
	_templateId int64
}

// New
