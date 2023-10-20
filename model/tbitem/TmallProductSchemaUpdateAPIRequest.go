package tbitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallproductschemaupdateAPIRequest 产品更新接口 API请求
// tmall.product.schema.update
//
// 产品更新接口
type TmallproductschemaupdateAPIRequest struct {
	model.Params
	// 根据tmall.product.update.schema.get生成的产品更新规则入参数据
	_xmlData string
	// 产品编号
	_productId int64
}

// NewTmallproductschemaupdateRequest 初始化TmallproductschemaupdateAPIRequest对象
func NewTmallproductschemaupdateRequest() *TmallproductschemaupdateAPIRequest {
	return &TmallproductschemaupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallproductschemaupdateAPIRequest) GetApiMethodName() string {
	return "tmall.product.schema.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallproductschemaupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallproductschemaupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetXmlData is XmlData Setter
// 根据tmall.product.update.schema.get生成的产品更新规则入参数据
func (r *TmallproductschemaupdateAPIRequest) SetXmlData(_xmlData string) error {
	r._xmlData = _xmlData
	r.Set("xml_data", _xmlData)
	return nil
}

// GetXmlData XmlData Getter
func (r TmallproductschemaupdateAPIRequest) GetXmlData() string {
	return r._xmlData
}

// SetProductId is ProductId Setter
// 产品编号
func (r *TmallproductschemaupdateAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TmallproductschemaupdateAPIRequest) GetProductId() int64 {
	return r._productId
}
