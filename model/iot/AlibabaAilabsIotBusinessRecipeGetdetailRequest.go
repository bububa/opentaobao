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
type AlibabaAilabsIotBusinessRecipeGetdetailRequest struct {
    model.Params
    // 行业食谱id
    businessRecipeId   int64
    // 开放账号id
    openAccountId   string
}

// 初始化AlibabaAilabsIotBusinessRecipeGetdetailRequest对象
func NewAlibabaAilabsIotBusinessRecipeGetdetailRequest() *AlibabaAilabsIotBusinessRecipeGetdetailRequest{
    return &AlibabaAilabsIotBusinessRecipeGetdetailRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAilabsIotBusinessRecipeGetdetailRequest) GetApiMethodName() string {
    return "alibaba.ailabs.iot.business.recipe.getdetail"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAilabsIotBusinessRecipeGetdetailRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BusinessRecipeId Setter
// 行业食谱id
func (r *AlibabaAilabsIotBusinessRecipeGetdetailRequest) SetBusinessRecipeId(businessRecipeId int64) error {
    r.businessRecipeId = businessRecipeId
    r.Set("business_recipe_id", businessRecipeId)
    return nil
}

// BusinessRecipeId Getter
func (r AlibabaAilabsIotBusinessRecipeGetdetailRequest) GetBusinessRecipeId() int64 {
    return r.businessRecipeId
}
// OpenAccountId Setter
// 开放账号id
func (r *AlibabaAilabsIotBusinessRecipeGetdetailRequest) SetOpenAccountId(openAccountId string) error {
    r.openAccountId = openAccountId
    r.Set("open_account_id", openAccountId)
    return nil
}

// OpenAccountId Getter
func (r AlibabaAilabsIotBusinessRecipeGetdetailRequest) GetOpenAccountId() string {
    return r.openAccountId
}
