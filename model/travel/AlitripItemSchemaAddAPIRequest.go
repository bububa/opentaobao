package travel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripitemschemaaddAPIRequest 使用schema模板进行商品发布 API请求
// alitrip.item.schema.add
//
// 飞猪度假商品使用schema模板进行商品发布。目前支持类目：出境自由行(50278002)、境内自由行(50272002)、出境跟团游(50258005)、境内跟团游(50258004)、境外一日游/多日游(50276003)
type AlitripitemschemaaddAPIRequest struct {
	model.Params
	// 商品数据
	_schemaXmlFields string
	// 类目id
	_catId int64
}

// NewAlitripitemschemaaddRequest 初始化AlitripitemschemaaddAPIRequest对象
func NewAlitripitemschemaaddRequest() *AlitripitemschemaaddAPIRequest {
	return &AlitripitemschemaaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripitemschemaaddAPIRequest) GetApiMethodName() string {
	return "alitrip.item.schema.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripitemschemaaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripitemschemaaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSchemaXmlFields is SchemaXmlFields Setter
// 商品数据
func (r *AlitripitemschemaaddAPIRequest) SetSchemaXmlFields(_schemaXmlFields string) error {
	r._schemaXmlFields = _schemaXmlFields
	r.Set("schema_xml_fields", _schemaXmlFields)
	return nil
}

// GetSchemaXmlFields SchemaXmlFields Getter
func (r AlitripitemschemaaddAPIRequest) GetSchemaXmlFields() string {
	return r._schemaXmlFields
}

// SetCatId is CatId Setter
// 类目id
func (r *AlitripitemschemaaddAPIRequest) SetCatId(_catId int64) error {
	r._catId = _catId
	r.Set("cat_id", _catId)
	return nil
}

// GetCatId CatId Getter
func (r AlitripitemschemaaddAPIRequest) GetCatId() int64 {
	return r._catId
}
