package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
product匹配接口 API请求
tmall.product.schema.match

根据tmall.product.match.schema.get获取到的规则，填充相应地的字段值以及类目，匹配符合条件的产品，返回匹配product结果，注意，有可能返回多个产品ID，以逗号分隔（尤其是图书类目）；
*/
type TmallProductSchemaMatchAPIRequest struct {
    model.Params
    // 商品发布的目标类目，必须是叶子类目
    _categoryId   int64
    // 根据tmall.product.match.schema.get获取到的模板，ISV将需要的字段填充好相应的值结果XML。
    _propvalues   string
}

// 初始化TmallProductSchemaMatchAPIRequest对象
func NewTmallProductSchemaMatchRequest() *TmallProductSchemaMatchAPIRequest{
    return &TmallProductSchemaMatchAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallProductSchemaMatchAPIRequest) GetApiMethodName() string {
    return "tmall.product.schema.match"
}

// IRequest interface 方法, 获取API参数
func (r TmallProductSchemaMatchAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CategoryId Setter
// 商品发布的目标类目，必须是叶子类目
func (r *TmallProductSchemaMatchAPIRequest) SetCategoryId(_categoryId int64) error {
    r._categoryId = _categoryId
    r.Set("category_id", _categoryId)
    return nil
}

// CategoryId Getter
func (r TmallProductSchemaMatchAPIRequest) GetCategoryId() int64 {
    return r._categoryId
}
// Propvalues Setter
// 根据tmall.product.match.schema.get获取到的模板，ISV将需要的字段填充好相应的值结果XML。
func (r *TmallProductSchemaMatchAPIRequest) SetPropvalues(_propvalues string) error {
    r._propvalues = _propvalues
    r.Set("propvalues", _propvalues)
    return nil
}

// Propvalues Getter
func (r TmallProductSchemaMatchAPIRequest) GetPropvalues() string {
    return r._propvalues
}
