package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
增加商品池里面的类目 API请求
alibaba.wdk.marketing.itempool.addcategory

增加商品池里面的类目
*/
type AlibabaWdkMarketingItempoolAddcategoryAPIRequest struct {
    model.Params
    // 类目对象
    _itemPoolActivityCategory   *ItemPoolActivityCategory
    // 活动对象
    _commonActivityParam   *CommonActivityParam
}

// 初始化AlibabaWdkMarketingItempoolAddcategoryAPIRequest对象
func NewAlibabaWdkMarketingItempoolAddcategoryRequest() *AlibabaWdkMarketingItempoolAddcategoryAPIRequest{
    return &AlibabaWdkMarketingItempoolAddcategoryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItempoolAddcategoryAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.itempool.addcategory"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItempoolAddcategoryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemPoolActivityCategory Setter
// 类目对象
func (r *AlibabaWdkMarketingItempoolAddcategoryAPIRequest) SetItemPoolActivityCategory(_itemPoolActivityCategory *ItemPoolActivityCategory) error {
    r._itemPoolActivityCategory = _itemPoolActivityCategory
    r.Set("item_pool_activity_category", _itemPoolActivityCategory)
    return nil
}

// ItemPoolActivityCategory Getter
func (r AlibabaWdkMarketingItempoolAddcategoryAPIRequest) GetItemPoolActivityCategory() *ItemPoolActivityCategory {
    return r._itemPoolActivityCategory
}
// CommonActivityParam Setter
// 活动对象
func (r *AlibabaWdkMarketingItempoolAddcategoryAPIRequest) SetCommonActivityParam(_commonActivityParam *CommonActivityParam) error {
    r._commonActivityParam = _commonActivityParam
    r.Set("common_activity_param", _commonActivityParam)
    return nil
}

// CommonActivityParam Getter
func (r AlibabaWdkMarketingItempoolAddcategoryAPIRequest) GetCommonActivityParam() *CommonActivityParam {
    return r._commonActivityParam
}
