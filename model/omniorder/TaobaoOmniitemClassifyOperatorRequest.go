package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
添加/修改分类 APIRequest
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

func NewTaobaoOmniitemClassifyOperatorRequest() *TaobaoOmniitemClassifyOperatorRequest{
    return &TaobaoOmniitemClassifyOperatorRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOmniitemClassifyOperatorRequest) GetApiMethodName() string {
    return "taobao.omniitem.classify.operator"
}

func (r TaobaoOmniitemClassifyOperatorRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOmniitemClassifyOperatorRequest) SetCategory(category *OmniItemCategoryDetailDto) error {
    r.category = category
    r.Set("category", category)
    return nil
}

func (r TaobaoOmniitemClassifyOperatorRequest) GetCategory() *OmniItemCategoryDetailDto {
    return r.category
}

func (r *TaobaoOmniitemClassifyOperatorRequest) SetAddItemIds(addItemIds []int64) error {
    r.addItemIds = addItemIds
    r.Set("add_item_ids", addItemIds)
    return nil
}

func (r TaobaoOmniitemClassifyOperatorRequest) GetAddItemIds() []int64 {
    return r.addItemIds
}

func (r *TaobaoOmniitemClassifyOperatorRequest) SetRemoveItemIds(removeItemIds []int64) error {
    r.removeItemIds = removeItemIds
    r.Set("remove_item_ids", removeItemIds)
    return nil
}

func (r TaobaoOmniitemClassifyOperatorRequest) GetRemoveItemIds() []int64 {
    return r.removeItemIds
}

func (r *TaobaoOmniitemClassifyOperatorRequest) SetOperator(operator string) error {
    r.operator = operator
    r.Set("operator", operator)
    return nil
}

func (r TaobaoOmniitemClassifyOperatorRequest) GetOperator() string {
    return r.operator
}

