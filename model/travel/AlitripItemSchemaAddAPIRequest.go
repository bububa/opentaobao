package travel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripItemSchemaAddAPIRequest 使用schema模板进行商品发布 API请求
// alitrip.item.schema.add
//
// 飞猪度假商品使用schema模板进行商品发布。目前支持类目：出境自由行(50278002)、境内自由行(50272002)、出境跟团游(50258005)、境内跟团游(50258004)、境外一日游/多日游(50276003)
type AlitripItemSchemaAddAPIRequest struct {
	model.Params
	// 商品数据
	_schemaXmlFields string
	// 类目id
	_catId int64
}

// NewAlitripItemSchemaAddRequest 初始化AlitripItemSchemaAddAPIRequest对象
func NewAlitripItemSchemaAddRequest() *AlitripItemSchemaAddAPIRequest {
	return &AlitripItemSchemaAddAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripItemSchemaAddAPIRequest) Reset() {
	r._schemaXmlFields = ""
	r._catId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripItemSchemaAddAPIRequest) GetApiMethodName() string {
	return "alitrip.item.schema.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripItemSchemaAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripItemSchemaAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSchemaXmlFields is SchemaXmlFields Setter
// 商品数据
func (r *AlitripItemSchemaAddAPIRequest) SetSchemaXmlFields(_schemaXmlFields string) error {
	r._schemaXmlFields = _schemaXmlFields
	r.Set("schema_xml_fields", _schemaXmlFields)
	return nil
}

// GetSchemaXmlFields SchemaXmlFields Getter
func (r AlitripItemSchemaAddAPIRequest) GetSchemaXmlFields() string {
	return r._schemaXmlFields
}

// SetCatId is CatId Setter
// 类目id
func (r *AlitripItemSchemaAddAPIRequest) SetCatId(_catId int64) error {
	r._catId = _catId
	r.Set("cat_id", _catId)
	return nil
}

// GetCatId CatId Getter
func (r AlitripItemSchemaAddAPIRequest) GetCatId() int64 {
	return r._catId
}

var poolAlitripItemSchemaAddAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripItemSchemaAddRequest()
	},
}

// GetAlitripItemSchemaAddRequest 从 sync.Pool 获取 AlitripItemSchemaAddAPIRequest
func GetAlitripItemSchemaAddAPIRequest() *AlitripItemSchemaAddAPIRequest {
	return poolAlitripItemSchemaAddAPIRequest.Get().(*AlitripItemSchemaAddAPIRequest)
}

// ReleaseAlitripItemSchemaAddAPIRequest 将 AlitripItemSchemaAddAPIRequest 放入 sync.Pool
func ReleaseAlitripItemSchemaAddAPIRequest(v *AlitripItemSchemaAddAPIRequest) {
	v.Reset()
	poolAlitripItemSchemaAddAPIRequest.Put(v)
}
