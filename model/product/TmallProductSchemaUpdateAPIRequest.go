package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallProductSchemaUpdateAPIRequest 产品更新接口 API请求
// tmall.product.schema.update
//
// 产品更新接口
type TmallProductSchemaUpdateAPIRequest struct {
	model.Params
	// 根据tmall.product.update.schema.get生成的产品更新规则入参数据
	_xmlData string
	// 产品编号
	_productId int64
}

// NewTmallProductSchemaUpdateRequest 初始化TmallProductSchemaUpdateAPIRequest对象
func NewTmallProductSchemaUpdateRequest() *TmallProductSchemaUpdateAPIRequest {
	return &TmallProductSchemaUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallProductSchemaUpdateAPIRequest) GetApiMethodName() string {
	return "tmall.product.schema.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallProductSchemaUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is XmlData Setter
// 根据tmall.product.update.schema.get生成的产品更新规则入参数据
func (r *TmallProductSchemaUpdateAPIRequest) SetXmlData(_xmlData string) error {
	r._xmlData = _xmlData
	r.Set("xml_data", _xmlData)
	return nil
}

// Get XmlData Getter
func (r TmallProductSchemaUpdateAPIRequest) GetXmlData() string {
	return r._xmlData
}

// Set is ProductId Setter
// 产品编号
func (r *TmallProductSchemaUpdateAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// Get ProductId Getter
func (r TmallProductSchemaUpdateAPIRequest) GetProductId() int64 {
	return r._productId
}
