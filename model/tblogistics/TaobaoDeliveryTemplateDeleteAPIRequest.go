package tblogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoDeliveryTemplateDeleteAPIRequest 删除运费模板 API请求
// taobao.delivery.template.delete
//
// 根据用户指定的模板ID删除指定的模板
type TaobaoDeliveryTemplateDeleteAPIRequest struct {
	model.Params
	// 运费模板ID
	_templateId int64
}

// NewTaobaoDeliveryTemplateDeleteRequest 初始化TaobaoDeliveryTemplateDeleteAPIRequest对象
func NewTaobaoDeliveryTemplateDeleteRequest() *TaobaoDeliveryTemplateDeleteAPIRequest {
	return &TaobaoDeliveryTemplateDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoDeliveryTemplateDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.delivery.template.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoDeliveryTemplateDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoDeliveryTemplateDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTemplateId is TemplateId Setter
// 运费模板ID
func (r *TaobaoDeliveryTemplateDeleteAPIRequest) SetTemplateId(_templateId int64) error {
	r._templateId = _templateId
	r.Set("template_id", _templateId)
	return nil
}

// GetTemplateId TemplateId Getter
func (r TaobaoDeliveryTemplateDeleteAPIRequest) GetTemplateId() int64 {
	return r._templateId
}
