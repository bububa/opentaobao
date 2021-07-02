package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuCategorySchemaLevelGetAPIRequest (新)层级属性获取 API请求
// alibaba.icbu.category.schema.level.get
//
// 将表单中层级属性的子属性返回
type AlibabaIcbuCategorySchemaLevelGetAPIRequest struct {
	model.Params
	// 类目id
	_catId int64
	// 返回的文案的语种，可以输入en_US或者zh
	_language string
	// 层级属性的当前层级属性
	_xml string
}

// NewAlibabaIcbuCategorySchemaLevelGetRequest 初始化AlibabaIcbuCategorySchemaLevelGetAPIRequest对象
func NewAlibabaIcbuCategorySchemaLevelGetRequest() *AlibabaIcbuCategorySchemaLevelGetAPIRequest {
	return &AlibabaIcbuCategorySchemaLevelGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIcbuCategorySchemaLevelGetAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.category.schema.level.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIcbuCategorySchemaLevelGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CatId Setter
// 类目id
func (r *AlibabaIcbuCategorySchemaLevelGetAPIRequest) SetCatId(_catId int64) error {
	r._catId = _catId
	r.Set("cat_id", _catId)
	return nil
}

// Get CatId Getter
func (r AlibabaIcbuCategorySchemaLevelGetAPIRequest) GetCatId() int64 {
	return r._catId
}

// Set is Language Setter
// 返回的文案的语种，可以输入en_US或者zh
func (r *AlibabaIcbuCategorySchemaLevelGetAPIRequest) SetLanguage(_language string) error {
	r._language = _language
	r.Set("language", _language)
	return nil
}

// Get Language Getter
func (r AlibabaIcbuCategorySchemaLevelGetAPIRequest) GetLanguage() string {
	return r._language
}

// Set is Xml Setter
// 层级属性的当前层级属性
func (r *AlibabaIcbuCategorySchemaLevelGetAPIRequest) SetXml(_xml string) error {
	r._xml = _xml
	r.Set("xml", _xml)
	return nil
}

// Get Xml Getter
func (r AlibabaIcbuCategorySchemaLevelGetAPIRequest) GetXml() string {
	return r._xml
}
