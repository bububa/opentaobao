package icbu

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuCategorySchemaLevelGetAPIRequest (新)层级属性获取 API请求
// alibaba.icbu.category.schema.level.get
//
// 将表单中层级属性的子属性返回
type AlibabaIcbuCategorySchemaLevelGetAPIRequest struct {
	model.Params
	// 返回的文案的语种，可以输入en_US或者zh
	_language string
	// 层级属性的当前层级属性
	_xml string
	// 类目id
	_catId int64
}

// NewAlibabaIcbuCategorySchemaLevelGetRequest 初始化AlibabaIcbuCategorySchemaLevelGetAPIRequest对象
func NewAlibabaIcbuCategorySchemaLevelGetRequest() *AlibabaIcbuCategorySchemaLevelGetAPIRequest {
	return &AlibabaIcbuCategorySchemaLevelGetAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIcbuCategorySchemaLevelGetAPIRequest) Reset() {
	r._language = ""
	r._xml = ""
	r._catId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIcbuCategorySchemaLevelGetAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.category.schema.level.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIcbuCategorySchemaLevelGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIcbuCategorySchemaLevelGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLanguage is Language Setter
// 返回的文案的语种，可以输入en_US或者zh
func (r *AlibabaIcbuCategorySchemaLevelGetAPIRequest) SetLanguage(_language string) error {
	r._language = _language
	r.Set("language", _language)
	return nil
}

// GetLanguage Language Getter
func (r AlibabaIcbuCategorySchemaLevelGetAPIRequest) GetLanguage() string {
	return r._language
}

// SetXml is Xml Setter
// 层级属性的当前层级属性
func (r *AlibabaIcbuCategorySchemaLevelGetAPIRequest) SetXml(_xml string) error {
	r._xml = _xml
	r.Set("xml", _xml)
	return nil
}

// GetXml Xml Getter
func (r AlibabaIcbuCategorySchemaLevelGetAPIRequest) GetXml() string {
	return r._xml
}

// SetCatId is CatId Setter
// 类目id
func (r *AlibabaIcbuCategorySchemaLevelGetAPIRequest) SetCatId(_catId int64) error {
	r._catId = _catId
	r.Set("cat_id", _catId)
	return nil
}

// GetCatId CatId Getter
func (r AlibabaIcbuCategorySchemaLevelGetAPIRequest) GetCatId() int64 {
	return r._catId
}

var poolAlibabaIcbuCategorySchemaLevelGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIcbuCategorySchemaLevelGetRequest()
	},
}

// GetAlibabaIcbuCategorySchemaLevelGetRequest 从 sync.Pool 获取 AlibabaIcbuCategorySchemaLevelGetAPIRequest
func GetAlibabaIcbuCategorySchemaLevelGetAPIRequest() *AlibabaIcbuCategorySchemaLevelGetAPIRequest {
	return poolAlibabaIcbuCategorySchemaLevelGetAPIRequest.Get().(*AlibabaIcbuCategorySchemaLevelGetAPIRequest)
}

// ReleaseAlibabaIcbuCategorySchemaLevelGetAPIRequest 将 AlibabaIcbuCategorySchemaLevelGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaIcbuCategorySchemaLevelGetAPIRequest(v *AlibabaIcbuCategorySchemaLevelGetAPIRequest) {
	v.Reset()
	poolAlibabaIcbuCategorySchemaLevelGetAPIRequest.Put(v)
}
