package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallitemsizemappingtemplategetAPIRequest 获取天猫商品尺码表模板 API请求
// tmall.item.sizemapping.template.get
//
// 获取天猫商品尺码表模板
type TmallitemsizemappingtemplategetAPIRequest struct {
	model.Params
	// 尺码表模板ID
	_templateId int64
}

// NewTmallitemsizemappingtemplategetRequest 初始化TmallitemsizemappingtemplategetAPIRequest对象
func NewTmallitemsizemappingtemplategetRequest() *TmallitemsizemappingtemplategetAPIRequest {
	return &TmallitemsizemappingtemplategetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallitemsizemappingtemplategetAPIRequest) GetApiMethodName() string {
	return "tmall.item.sizemapping.template.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallitemsizemappingtemplategetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallitemsizemappingtemplategetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTemplateId is TemplateId Setter
// 尺码表模板ID
func (r *TmallitemsizemappingtemplategetAPIRequest) SetTemplateId(_templateId int64) error {
	r._templateId = _templateId
	r.Set("template_id", _templateId)
	return nil
}

// GetTemplateId TemplateId Getter
func (r TmallitemsizemappingtemplategetAPIRequest) GetTemplateId() int64 {
	return r._templateId
}
