package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallItemSizemappingTemplatesListAPIRequest
获取天猫商品尺码表模板列表 API请求
tmall.item.sizemapping.templates.list

获取所有尺码表模板列表。 */
type TmallItemSizemappingTemplatesListAPIRequest struct {
	model.Params
}

// NewTmallItemSizemappingTemplatesListRequest 初始化TmallItemSizemappingTemplatesListAPIRequest对象
func NewTmallItemSizemappingTemplatesListRequest() *TmallItemSizemappingTemplatesListAPIRequest {
	return &TmallItemSizemappingTemplatesListAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallItemSizemappingTemplatesListAPIRequest) GetApiMethodName() string {
	return "tmall.item.sizemapping.templates.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallItemSizemappingTemplatesListAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
