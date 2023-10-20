package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallitemsizemappingtemplatedeleteAPIRequest 删除天猫商品尺码表模板 API请求
// tmall.item.sizemapping.template.delete
//
// 删除天猫商品尺码表模板
type TmallitemsizemappingtemplatedeleteAPIRequest struct {
	model.Params
	// 尺码表模板ID
	_templateId int64
}

// NewTmallitemsizemappingtemplatedeleteRequest 初始化TmallitemsizemappingtemplatedeleteAPIRequest对象
func NewTmallitemsizemappingtemplatedeleteRequest() *TmallitemsizemappingtemplatedeleteAPIRequest {
	return &TmallitemsizemappingtemplatedeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallitemsizemappingtemplatedeleteAPIRequest) GetApiMethodName() string {
	return "tmall.item.sizemapping.template.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallitemsizemappingtemplatedeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallitemsizemappingtemplatedeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTemplateId is TemplateId Setter
// 尺码表模板ID
func (r *TmallitemsizemappingtemplatedeleteAPIRequest) SetTemplateId(_templateId int64) error {
	r._templateId = _templateId
	r.Set("template_id", _templateId)
	return nil
}

// GetTemplateId TemplateId Getter
func (r TmallitemsizemappingtemplatedeleteAPIRequest) GetTemplateId() int64 {
	return r._templateId
}
