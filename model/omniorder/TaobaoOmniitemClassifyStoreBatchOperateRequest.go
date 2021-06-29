package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量添加/删除门店和分类的关联关系 API请求
taobao.omniitem.classify.store.batch.operate

批量添加/删除门店和分类的关联关系
*/
type TaobaoOmniitemClassifyStoreBatchOperateRequest struct {
    model.Params
    // 商家门店ID
    _storeIds   []int64
    // 需要添加的分类ID
    _addCategoryIds   []int64
    // 需要删除的分类ID
    _removeCategoryIds   []int64
    // 操作信息（暂时不填）
    _operator   string
}

// 初始化TaobaoOmniitemClassifyStoreBatchOperateRequest对象
func NewTaobaoOmniitemClassifyStoreBatchOperateRequest() *TaobaoOmniitemClassifyStoreBatchOperateRequest{
    return &TaobaoOmniitemClassifyStoreBatchOperateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOmniitemClassifyStoreBatchOperateRequest) GetApiMethodName() string {
    return "taobao.omniitem.classify.store.batch.operate"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOmniitemClassifyStoreBatchOperateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreIds Setter
// 商家门店ID
func (r *TaobaoOmniitemClassifyStoreBatchOperateRequest) SetStoreIds(_storeIds []int64) error {
    r._storeIds = _storeIds
    r.Set("store_ids", _storeIds)
    return nil
}

// StoreIds Getter
func (r TaobaoOmniitemClassifyStoreBatchOperateRequest) GetStoreIds() []int64 {
    return r._storeIds
}
// AddCategoryIds Setter
// 需要添加的分类ID
func (r *TaobaoOmniitemClassifyStoreBatchOperateRequest) SetAddCategoryIds(_addCategoryIds []int64) error {
    r._addCategoryIds = _addCategoryIds
    r.Set("add_category_ids", _addCategoryIds)
    return nil
}

// AddCategoryIds Getter
func (r TaobaoOmniitemClassifyStoreBatchOperateRequest) GetAddCategoryIds() []int64 {
    return r._addCategoryIds
}
// RemoveCategoryIds Setter
// 需要删除的分类ID
func (r *TaobaoOmniitemClassifyStoreBatchOperateRequest) SetRemoveCategoryIds(_removeCategoryIds []int64) error {
    r._removeCategoryIds = _removeCategoryIds
    r.Set("remove_category_ids", _removeCategoryIds)
    return nil
}

// RemoveCategoryIds Getter
func (r TaobaoOmniitemClassifyStoreBatchOperateRequest) GetRemoveCategoryIds() []int64 {
    return r._removeCategoryIds
}
// Operator Setter
// 操作信息（暂时不填）
func (r *TaobaoOmniitemClassifyStoreBatchOperateRequest) SetOperator(_operator string) error {
    r._operator = _operator
    r.Set("operator", _operator)
    return nil
}

// Operator Getter
func (r TaobaoOmniitemClassifyStoreBatchOperateRequest) GetOperator() string {
    return r._operator
}
