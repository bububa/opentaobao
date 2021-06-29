package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
添加/修改分类 API请求
taobao.omniitem.classify.operator

添加/修改分类
*/
type TaobaoOmniitemClassifyOperatorRequest struct {
    model.Params
    // 分类信息
    category   *OmniItemCategoryDetailDto
    // 需要添加的关联关系的商品
    addItemIds   []int64
    // 需要修改的关联关系的商品
    removeItemIds   []int64
    // 操作人信息（暂时不填）
    operator   string
}

// 初始化TaobaoOmniitemClassifyOperatorRequest对象
func NewTaobaoOmniitemClassifyOperatorRequest() *TaobaoOmniitemClassifyOperatorRequest{
    return &TaobaoOmniitemClassifyOperatorRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOmniitemClassifyOperatorRequest) GetApiMethodName() string {
    return "taobao.omniitem.classify.operator"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOmniitemClassifyOperatorRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Category Setter
// 分类信息
func (r *TaobaoOmniitemClassifyOperatorRequest) SetCategory(category *OmniItemCategoryDetailDto) error {
    r.category = category
    r.Set("category", category)
    return nil
}

// Category Getter
func (r TaobaoOmniitemClassifyOperatorRequest) GetCategory() *OmniItemCategoryDetailDto {
    return r.category
}
// AddItemIds Setter
// 需要添加的关联关系的商品
func (r *TaobaoOmniitemClassifyOperatorRequest) SetAddItemIds(addItemIds []int64) error {
    r.addItemIds = addItemIds
    r.Set("add_item_ids", addItemIds)
    return nil
}

// AddItemIds Getter
func (r TaobaoOmniitemClassifyOperatorRequest) GetAddItemIds() []int64 {
    return r.addItemIds
}
// RemoveItemIds Setter
// 需要修改的关联关系的商品
func (r *TaobaoOmniitemClassifyOperatorRequest) SetRemoveItemIds(removeItemIds []int64) error {
    r.removeItemIds = removeItemIds
    r.Set("remove_item_ids", removeItemIds)
    return nil
}

// RemoveItemIds Getter
func (r TaobaoOmniitemClassifyOperatorRequest) GetRemoveItemIds() []int64 {
    return r.removeItemIds
}
// Operator Setter
// 操作人信息（暂时不填）
func (r *TaobaoOmniitemClassifyOperatorRequest) SetOperator(operator string) error {
    r.operator = operator
    r.Set("operator", operator)
    return nil
}

// Operator Getter
func (r TaobaoOmniitemClassifyOperatorRequest) GetOperator() string {
    return r.operator
}
