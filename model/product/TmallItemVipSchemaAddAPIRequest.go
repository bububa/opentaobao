package product

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallItemVipSchemaAddAPIRequest 大商家商品发布接口 API请求
// tmall.item.vip.schema.add
//
// 大商家商品发布接口
type TmallItemVipSchemaAddAPIRequest struct {
	model.Params
	// 商品发布schema参数
	_schemaXmlFields string
}

// NewTmallItemVipSchemaAddRequest 初始化TmallItemVipSchemaAddAPIRequest对象
func NewTmallItemVipSchemaAddRequest() *TmallItemVipSchemaAddAPIRequest {
	return &TmallItemVipSchemaAddAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallItemVipSchemaAddAPIRequest) Reset() {
	r._schemaXmlFields = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallItemVipSchemaAddAPIRequest) GetApiMethodName() string {
	return "tmall.item.vip.schema.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallItemVipSchemaAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallItemVipSchemaAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSchemaXmlFields is SchemaXmlFields Setter
// 商品发布schema参数
func (r *TmallItemVipSchemaAddAPIRequest) SetSchemaXmlFields(_schemaXmlFields string) error {
	r._schemaXmlFields = _schemaXmlFields
	r.Set("schema_xml_fields", _schemaXmlFields)
	return nil
}

// GetSchemaXmlFields SchemaXmlFields Getter
func (r TmallItemVipSchemaAddAPIRequest) GetSchemaXmlFields() string {
	return r._schemaXmlFields
}

var poolTmallItemVipSchemaAddAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallItemVipSchemaAddRequest()
	},
}

// GetTmallItemVipSchemaAddRequest 从 sync.Pool 获取 TmallItemVipSchemaAddAPIRequest
func GetTmallItemVipSchemaAddAPIRequest() *TmallItemVipSchemaAddAPIRequest {
	return poolTmallItemVipSchemaAddAPIRequest.Get().(*TmallItemVipSchemaAddAPIRequest)
}

// ReleaseTmallItemVipSchemaAddAPIRequest 将 TmallItemVipSchemaAddAPIRequest 放入 sync.Pool
func ReleaseTmallItemVipSchemaAddAPIRequest(v *TmallItemVipSchemaAddAPIRequest) {
	v.Reset()
	poolTmallItemVipSchemaAddAPIRequest.Put(v)
}
