package tbitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallproductschemaaddAPIRequest 使用Schema文件发布一个产品 API请求
// tmall.product.schema.add
//
// Schema体系发布一个产品
type TmallproductschemaaddAPIRequest struct {
	model.Params
	// 根据tmall.product.add.schema.get生成的产品发布规则入参数据
	_xmlData string
	// 商品发布的目标类目，必须是叶子类目
	_categoryId int64
	// 品牌ID
	_brandId int64
}

// NewTmallproductschemaaddRequest 初始化TmallproductschemaaddAPIRequest对象
func NewTmallproductschemaaddRequest() *TmallproductschemaaddAPIRequest {
	return &TmallproductschemaaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallproductschemaaddAPIRequest) GetApiMethodName() string {
	return "tmall.product.schema.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallproductschemaaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallproductschemaaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetXmlData is XmlData Setter
// 根据tmall.product.add.schema.get生成的产品发布规则入参数据
func (r *TmallproductschemaaddAPIRequest) SetXmlData(_xmlData string) error {
	r._xmlData = _xmlData
	r.Set("xml_data", _xmlData)
	return nil
}

// GetXmlData XmlData Getter
func (r TmallproductschemaaddAPIRequest) GetXmlData() string {
	return r._xmlData
}

// SetCategoryId is CategoryId Setter
// 商品发布的目标类目，必须是叶子类目
func (r *TmallproductschemaaddAPIRequest) SetCategoryId(_categoryId int64) error {
	r._categoryId = _categoryId
	r.Set("category_id", _categoryId)
	return nil
}

// GetCategoryId CategoryId Getter
func (r TmallproductschemaaddAPIRequest) GetCategoryId() int64 {
	return r._categoryId
}

// SetBrandId is BrandId Setter
// 品牌ID
func (r *TmallproductschemaaddAPIRequest) SetBrandId(_brandId int64) error {
	r._brandId = _brandId
	r.Set("brand_id", _brandId)
	return nil
}

// GetBrandId BrandId Getter
func (r TmallproductschemaaddAPIRequest) GetBrandId() int64 {
	return r._brandId
}
