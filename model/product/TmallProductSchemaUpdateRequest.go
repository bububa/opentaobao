package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
产品更新接口 API请求
tmall.product.schema.update

产品更新接口
*/
type TmallProductSchemaUpdateRequest struct {
    model.Params
    // 根据tmall.product.update.schema.get生成的产品更新规则入参数据
    _xmlData   string
    // 产品编号
    _productId   int64
}

// 初始化TmallProductSchemaUpdateRequest对象
func NewTmallProductSchemaUpdateRequest() *TmallProductSchemaUpdateRequest{
    return &TmallProductSchemaUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallProductSchemaUpdateRequest) GetApiMethodName() string {
    return "tmall.product.schema.update"
}

// IRequest interface 方法, 获取API参数
func (r TmallProductSchemaUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// XmlData Setter
// 根据tmall.product.update.schema.get生成的产品更新规则入参数据
func (r *TmallProductSchemaUpdateRequest) SetXmlData(_xmlData string) error {
    r._xmlData = _xmlData
    r.Set("xml_data", _xmlData)
    return nil
}

// XmlData Getter
func (r TmallProductSchemaUpdateRequest) GetXmlData() string {
    return r._xmlData
}
// ProductId Setter
// 产品编号
func (r *TmallProductSchemaUpdateRequest) SetProductId(_productId int64) error {
    r._productId = _productId
    r.Set("product_id", _productId)
    return nil
}

// ProductId Getter
func (r TmallProductSchemaUpdateRequest) GetProductId() int64 {
    return r._productId
}
