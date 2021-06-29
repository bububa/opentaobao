package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量添加/删除门店和分类的关联关系 APIRequest
taobao.omniitem.classify.store.batch.operate

批量添加/删除门店和分类的关联关系
*/
type TaobaoOmniitemClassifyStoreBatchOperateRequest struct {
    model.Params

    // 商家门店ID
    storeIds   []int64 

    // 需要添加的分类ID
    addCategoryIds   []int64 

    // 需要删除的分类ID
    removeCategoryIds   []int64 

    // 操作信息（暂时不填）
    operator   string 

}

func NewTaobaoOmniitemClassifyStoreBatchOperateRequest() *TaobaoOmniitemClassifyStoreBatchOperateRequest{
    return &TaobaoOmniitemClassifyStoreBatchOperateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOmniitemClassifyStoreBatchOperateRequest) GetApiMethodName() string {
    return "taobao.omniitem.classify.store.batch.operate"
}

func (r TaobaoOmniitemClassifyStoreBatchOperateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOmniitemClassifyStoreBatchOperateRequest) SetStoreIds(storeIds []int64) error {
    r.storeIds = storeIds
    r.Set("store_ids", storeIds)
    return nil
}

func (r TaobaoOmniitemClassifyStoreBatchOperateRequest) GetStoreIds() []int64 {
    return r.storeIds
}

func (r *TaobaoOmniitemClassifyStoreBatchOperateRequest) SetAddCategoryIds(addCategoryIds []int64) error {
    r.addCategoryIds = addCategoryIds
    r.Set("add_category_ids", addCategoryIds)
    return nil
}

func (r TaobaoOmniitemClassifyStoreBatchOperateRequest) GetAddCategoryIds() []int64 {
    return r.addCategoryIds
}

func (r *TaobaoOmniitemClassifyStoreBatchOperateRequest) SetRemoveCategoryIds(removeCategoryIds []int64) error {
    r.removeCategoryIds = removeCategoryIds
    r.Set("remove_category_ids", removeCategoryIds)
    return nil
}

func (r TaobaoOmniitemClassifyStoreBatchOperateRequest) GetRemoveCategoryIds() []int64 {
    return r.removeCategoryIds
}

func (r *TaobaoOmniitemClassifyStoreBatchOperateRequest) SetOperator(operator string) error {
    r.operator = operator
    r.Set("operator", operator)
    return nil
}

func (r TaobaoOmniitemClassifyStoreBatchOperateRequest) GetOperator() string {
    return r.operator
}

