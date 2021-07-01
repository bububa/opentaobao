package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallItemSizemappingTemplateUpdateAPIRequest
更新天猫商品尺码表模板 API请求
tmall.item.sizemapping.template.update

更新天猫商品尺码表模板 */
type TmallItemSizemappingTemplateUpdateAPIRequest struct {
	model.Params
	// 尺码表模板ID
	_templateId int64
	// 尺码表模板名称
	_templateName string
	// 尺码表模板内容，格式为"尺码值:维度名称:数值,尺码值:维度名称:数值"。其中，数值的单位，长度单位为厘米（cm），体重单位为公斤（kg）。尺码值，维度数据不能包含数字，特殊字符。数值为0-999.9的数字，且最多一位小数。
	_templateContent string
}

// New
