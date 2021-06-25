package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
增加商品池里面的类目 APIRequest
alibaba.wdk.marketing.itempool.addcategory

增加商品池里面的类目
*/
type AlibabaWdkMarketingItempoolAddcategoryRequest struct {
    model.Params

    // 类目对象
    itemPoolActivityCategory   *ItemPoolActivityCategory 

    // 活动对象
    commonActivityParam   *CommonActivityParam 

}

func NewAlibabaWdkMarketingItempoolAddcategoryRequest() *AlibabaWdkMarketingItempoolAddcategoryRequest{
    return &AlibabaWdkMarketingItempoolAddcategoryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkMarketingItempoolAddcategoryRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.itempool.addcategory"
}

func (r AlibabaWdkMarketingItempoolAddcategoryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkMarketingItempoolAddcategoryRequest) SetItemPoolActivityCategory(itemPoolActivityCategory *ItemPoolActivityCategory) error {
    r.itemPoolActivityCategory = itemPoolActivityCategory
    r.Set("item_pool_activity_category", itemPoolActivityCategory)
    return nil
}

func (r AlibabaWdkMarketingItempoolAddcategoryRequest) GetItemPoolActivityCategory() *ItemPoolActivityCategory {
    return r.itemPoolActivityCategory
}

func (r *AlibabaWdkMarketingItempoolAddcategoryRequest) SetCommonActivityParam(commonActivityParam *CommonActivityParam) error {
    r.commonActivityParam = commonActivityParam
    r.Set("common_activity_param", commonActivityParam)
    return nil
}

func (r AlibabaWdkMarketingItempoolAddcategoryRequest) GetCommonActivityParam() *CommonActivityParam {
    return r.commonActivityParam
}

