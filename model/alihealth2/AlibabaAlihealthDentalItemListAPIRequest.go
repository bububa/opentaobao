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
type AlibabaAlihealthDentalItemListAPIRequest struct {
    model.Params
    // 是否需要测试商品
    _needTestItem   bool
}

// 初始化AlibabaAlihealthDentalItemListAPIRequest对象
func NewAlibabaAlihealthDentalItemListRequest() *AlibabaAlihealthDentalItemListAPIRequest{
    return &AlibabaAlihealthDentalItemListAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDentalItemListAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.dental.item.list"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDentalItemListAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// NeedTestItem Setter
// 是否需要测试商品
func (r *AlibabaAlihealthDentalItemListAPIRequest) SetNeedTestItem(_needTestItem bool) error {
    r._needTestItem = _needTestItem
    r.Set("need_test_item", _needTestItem)
    return nil
}

// NeedTestItem Getter
func (r AlibabaAlihealthDentalItemListAPIRequest) GetNeedTestItem() bool {
    return r._needTestItem
}
