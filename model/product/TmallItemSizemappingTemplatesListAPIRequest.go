package product

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallItemSizemappingTemplatesListAPIRequest 获取天猫商品尺码表模板列表 API请求
// tmall.item.sizemapping.templates.list
//
// 获取所有尺码表模板列表。
type TmallItemSizemappingTemplatesListAPIRequest struct {
	model.Params
}

// NewTmallItemSizemappingTemplatesListRequest 初始化TmallItemSizemappingTemplatesListAPIRequest对象
func NewTmallItemSizemappingTemplatesListRequest() *TmallItemSizemappingTemplatesListAPIRequest {
	return &TmallItemSizemappingTemplatesListAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallItemSizemappingTemplatesListAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallItemSizemappingTemplatesListAPIRequest) GetApiMethodName() string {
	return "tmall.item.sizemapping.templates.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallItemSizemappingTemplatesListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallItemSizemappingTemplatesListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolTmallItemSizemappingTemplatesListAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallItemSizemappingTemplatesListRequest()
	},
}

// GetTmallItemSizemappingTemplatesListRequest 从 sync.Pool 获取 TmallItemSizemappingTemplatesListAPIRequest
func GetTmallItemSizemappingTemplatesListAPIRequest() *TmallItemSizemappingTemplatesListAPIRequest {
	return poolTmallItemSizemappingTemplatesListAPIRequest.Get().(*TmallItemSizemappingTemplatesListAPIRequest)
}

// ReleaseTmallItemSizemappingTemplatesListAPIRequest 将 TmallItemSizemappingTemplatesListAPIRequest 放入 sync.Pool
func ReleaseTmallItemSizemappingTemplatesListAPIRequest(v *TmallItemSizemappingTemplatesListAPIRequest) {
	v.Reset()
	poolTmallItemSizemappingTemplatesListAPIRequest.Put(v)
}
