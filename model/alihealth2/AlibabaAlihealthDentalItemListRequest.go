package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV获取口腔标品列表 APIRequest
alibaba.alihealth.dental.item.list

ISV获取口腔标品列表
*/
type AlibabaAlihealthDentalItemListRequest struct {
    model.Params

    // 是否需要测试商品
    needTestItem   bool 

}

func NewAlibabaAlihealthDentalItemListRequest() *AlibabaAlihealthDentalItemListRequest{
    return &AlibabaAlihealthDentalItemListRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDentalItemListRequest) GetApiMethodName() string {
    return "alibaba.alihealth.dental.item.list"
}

func (r AlibabaAlihealthDentalItemListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDentalItemListRequest) SetNeedTestItem(needTestItem bool) error {
    r.needTestItem = needTestItem
    r.Set("need_test_item", needTestItem)
    return nil
}

func (r AlibabaAlihealthDentalItemListRequest) GetNeedTestItem() bool {
    return r.needTestItem
}

