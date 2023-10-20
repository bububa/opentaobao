package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaicbucategoryschemalevelgetAPIRequest (新)层级属性获取 API请求
// alibaba.icbu.category.schema.level.get
//
// 将表单中层级属性的子属性返回
type AlibabaicbucategoryschemalevelgetAPIRequest struct {
	model.Params
	// 返回的文案的语种，可以输入en_US或者zh
	_language string
	// 层级属性的当前层级属性
	_xml string
	// 类目id
	_catId int64
}

// NewAlibabaicbucategoryschemalevelgetRequest 初始化AlibabaicbucategoryschemalevelgetAPIRequest对象
func NewAlibabaicbucategoryschemalevelgetRequest() *AlibabaicbucategoryschemalevelgetAPIRequest {
	return &AlibabaicbucategoryschemalevelgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaicbucategoryschemalevelgetAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.category.schema.level.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaicbucategoryschemalevelgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaicbucategoryschemalevelgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLanguage is Language Setter
// 返回的文案的语种，可以输入en_US或者zh
func (r *AlibabaicbucategoryschemalevelgetAPIRequest) SetLanguage(_language string) error {
	r._language = _language
	r.Set("language", _language)
	return nil
}

// GetLanguage Language Getter
func (r AlibabaicbucategoryschemalevelgetAPIRequest) GetLanguage() string {
	return r._language
}

// SetXml is Xml Setter
// 层级属性的当前层级属性
func (r *AlibabaicbucategoryschemalevelgetAPIRequest) SetXml(_xml string) error {
	r._xml = _xml
	r.Set("xml", _xml)
	return nil
}

// GetXml Xml Getter
func (r AlibabaicbucategoryschemalevelgetAPIRequest) GetXml() string {
	return r._xml
}

// SetCatId is CatId Setter
// 类目id
func (r *AlibabaicbucategoryschemalevelgetAPIRequest) SetCatId(_catId int64) error {
	r._catId = _catId
	r.Set("cat_id", _catId)
	return nil
}

// GetCatId CatId Getter
func (r AlibabaicbucategoryschemalevelgetAPIRequest) GetCatId() int64 {
	return r._catId
}
