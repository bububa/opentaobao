package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取食谱详情 APIRequest
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

func NewAlibabaAilabsIotBusinessRecipeGetdetailRequest() *AlibabaAilabsIotBusinessRecipeGetdetailRequest{
    return &AlibabaAilabsIotBusinessRecipeGetdetailRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAilabsIotBusinessRecipeGetdetailRequest) GetApiMethodName() string {
    return "alibaba.ailabs.iot.business.recipe.getdetail"
}

func (r AlibabaAilabsIotBusinessRecipeGetdetailRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAilabsIotBusinessRecipeGetdetailRequest) SetBusinessRecipeId(businessRecipeId int64) error {
    r.businessRecipeId = businessRecipeId
    r.Set("business_recipe_id", businessRecipeId)
    return nil
}

func (r AlibabaAilabsIotBusinessRecipeGetdetailRequest) GetBusinessRecipeId() int64 {
    return r.businessRecipeId
}

func (r *AlibabaAilabsIotBusinessRecipeGetdetailRequest) SetOpenAccountId(openAccountId string) error {
    r.openAccountId = openAccountId
    r.Set("open_account_id", openAccountId)
    return nil
}

func (r AlibabaAilabsIotBusinessRecipeGetdetailRequest) GetOpenAccountId() string {
    return r.openAccountId
}

