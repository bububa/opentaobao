package product

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallItemSizemappingTemplateGetAPIRequest 获取天猫商品尺码表模板 API请求
// tmall.item.sizemapping.template.get
//
// 获取天猫商品尺码表模板
type TmallItemSizemappingTemplateGetAPIRequest struct {
	model.Params
	// 尺码表模板ID
	_templateId int64
}

// NewTmallItemSizemappingTemplateGetRequest 初始化TmallItemSizemappingTemplateGetAPIRequest对象
func NewTmallItemSizemappingTemplateGetRequest() *TmallItemSizemappingTemplateGetAPIRequest {
	return &TmallItemSizemappingTemplateGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallItemSizemappingTemplateGetAPIRequest) Reset() {
	r._templateId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallItemSizemappingTemplateGetAPIRequest) GetApiMethodName() string {
	return "tmall.item.sizemapping.template.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallItemSizemappingTemplateGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallItemSizemappingTemplateGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTemplateId is TemplateId Setter
// 尺码表模板ID
func (r *TmallItemSizemappingTemplateGetAPIRequest) SetTemplateId(_templateId int64) error {
	r._templateId = _templateId
	r.Set("template_id", _templateId)
	return nil
}

// GetTemplateId TemplateId Getter
func (r TmallItemSizemappingTemplateGetAPIRequest) GetTemplateId() int64 {
	return r._templateId
}

var poolTmallItemSizemappingTemplateGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallItemSizemappingTemplateGetRequest()
	},
}

// GetTmallItemSizemappingTemplateGetRequest 从 sync.Pool 获取 TmallItemSizemappingTemplateGetAPIRequest
func GetTmallItemSizemappingTemplateGetAPIRequest() *TmallItemSizemappingTemplateGetAPIRequest {
	return poolTmallItemSizemappingTemplateGetAPIRequest.Get().(*TmallItemSizemappingTemplateGetAPIRequest)
}

// ReleaseTmallItemSizemappingTemplateGetAPIRequest 将 TmallItemSizemappingTemplateGetAPIRequest 放入 sync.Pool
func ReleaseTmallItemSizemappingTemplateGetAPIRequest(v *TmallItemSizemappingTemplateGetAPIRequest) {
	v.Reset()
	poolTmallItemSizemappingTemplateGetAPIRequest.Put(v)
}
