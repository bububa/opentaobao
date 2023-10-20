package tblogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaodeliverytemplatedeleteAPIRequest 删除运费模板 API请求
// taobao.delivery.template.delete
//
// 根据用户指定的模板ID删除指定的模板
type TaobaodeliverytemplatedeleteAPIRequest struct {
	model.Params
	// 运费模板ID
	_templateId int64
}

// NewTaobaodeliverytemplatedeleteRequest 初始化TaobaodeliverytemplatedeleteAPIRequest对象
func NewTaobaodeliverytemplatedeleteRequest() *TaobaodeliverytemplatedeleteAPIRequest {
	return &TaobaodeliverytemplatedeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaodeliverytemplatedeleteAPIRequest) GetApiMethodName() string {
	return "taobao.delivery.template.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaodeliverytemplatedeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaodeliverytemplatedeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTemplateId is TemplateId Setter
// 运费模板ID
func (r *TaobaodeliverytemplatedeleteAPIRequest) SetTemplateId(_templateId int64) error {
	r._templateId = _templateId
	r.Set("template_id", _templateId)
	return nil
}

// GetTemplateId TemplateId Getter
func (r TaobaodeliverytemplatedeleteAPIRequest) GetTemplateId() int64 {
	return r._templateId
}
