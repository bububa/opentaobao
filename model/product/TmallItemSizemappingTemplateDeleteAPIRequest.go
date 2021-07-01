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

// NewTmallItemSizemappingTemplateDeleteRequest 初始化TmallItemSizemappingTemplateDeleteAPIRequest对象
func NewTmallItemSizemappingTemplateDeleteRequest() *TmallItemSizemappingTemplateDeleteAPIRequest {
	return &TmallItemSizemappingTemplateDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallItemSizemappingTemplateDeleteAPIRequest) GetApiMethodName() string {
	return "tmall.item.sizemapping.template.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallItemSizemappingTemplateDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is TemplateId Setter
// 尺码表模板ID
func (r *TmallItemSizemappingTemplateDeleteAPIRequest) SetTemplateId(_templateId int64) error {
	r._templateId = _templateId
	r.Set("template_id", _templateId)
	return nil
}

// Get TemplateId Getter
func (r TmallItemSizemappingTemplateDeleteAPIRequest) GetTemplateId() int64 {
	return r._templateId
}
