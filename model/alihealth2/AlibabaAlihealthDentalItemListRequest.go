package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV获取口腔标品列表 API请求
alibaba.alihealth.dental.item.list

ISV获取口腔标品列表
*/
type AlibabaAlihealthDentalItemListRequest struct {
    model.Params
    // 是否需要测试商品
    needTestItem   bool
}

// 初始化AlibabaAlihealthDentalItemListRequest对象
func NewAlibabaAlihealthDentalItemListRequest() *AlibabaAlihealthDentalItemListRequest{
    return &AlibabaAlihealthDentalItemListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDentalItemListRequest) GetApiMethodName() string {
    return "alibaba.alihealth.dental.item.list"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDentalItemListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// NeedTestItem Setter
// 是否需要测试商品
func (r *AlibabaAlihealthDentalItemListRequest) SetNeedTestItem(needTestItem bool) error {
    r.needTestItem = needTestItem
    r.Set("need_test_item", needTestItem)
    return nil
}

// NeedTestItem Getter
func (r AlibabaAlihealthDentalItemListRequest) GetNeedTestItem() bool {
    return r.needTestItem
}
