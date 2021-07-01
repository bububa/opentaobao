package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoDeliveryTemplateDeleteAPIRequest
删除运费模板 API请求
taobao.delivery.template.delete

根据用户指定的模板ID删除指定的模板 */
type TaobaoDeliveryTemplateDeleteAPIRequest struct {
	model.Params
	// 运费模板ID
	_templateId int64
}

// New
