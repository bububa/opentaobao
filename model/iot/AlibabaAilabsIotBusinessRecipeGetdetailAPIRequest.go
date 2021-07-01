package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取食谱详情 API请求
alibaba.ailabs.iot.business.recipe.getdetail

获取食谱详情接口，获取ISV自己的食谱详情数据
*/
type AlibabaAilabsIotBusinessRecipeGetdetailAPIRequest struct {
    model.Params
    // 行业食谱id
    _businessRecipeId   int64
    // 开放账号id
    _openAccountId   string
}

// 初始化AlibabaAilabsIotBusinessRecipeGetdetailAPIRequest对象
func NewAlibabaAilabsIotBusinessRecipeGetdetailRequest() *AlibabaAilabsIotBusinessRecipeGetdetailAPIRequest{
    return &AlibabaAilabsIotBusinessRecipeGetdetailAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAilabsIotBusinessRecipeGetdetailAPIRequest) GetApiMethodName() string {
    return "alibaba.ailabs.iot.business.recipe.getdetail"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAilabsIotBusinessRecipeGetdetailAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BusinessRecipeId Setter
// 行业食谱id
func (r *AlibabaAilabsIotBusinessRecipeGetdetailAPIRequest) SetBusinessRecipeId(_businessRecipeId int64) error {
    r._businessRecipeId = _businessRecipeId
    r.Set("business_recipe_id", _businessRecipeId)
    return nil
}

// BusinessRecipeId Getter
func (r AlibabaAilabsIotBusinessRecipeGetdetailAPIRequest) GetBusinessRecipeId() int64 {
    return r._businessRecipeId
}
// OpenAccountId Setter
// 开放账号id
func (r *AlibabaAilabsIotBusinessRecipeGetdetailAPIRequest) SetOpenAccountId(_openAccountId string) error {
    r._openAccountId = _openAccountId
    r.Set("open_account_id", _openAccountId)
    return nil
}

// OpenAccountId Getter
func (r AlibabaAilabsIotBusinessRecipeGetdetailAPIRequest) GetOpenAccountId() string {
    return r._openAccountId
}
