package lstmarketing

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据订单查询营销信息 APIRequest
alibaba.lst.marketing.querybyorderid

根据订单查询营销信息
*/
type AlibabaLstMarketingQuerybyorderidRequest struct {
    model.Params

    // 主订单
    mainOrderId   int64 

    // 子订单
    subOrderId   int64 

}

func NewAlibabaLstMarketingQuerybyorderidRequest() *AlibabaLstMarketingQuerybyorderidRequest{
    return &AlibabaLstMarketingQuerybyorderidRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLstMarketingQuerybyorderidRequest) GetApiMethodName() string {
    return "alibaba.lst.marketing.querybyorderid"
}

func (r AlibabaLstMarketingQuerybyorderidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLstMarketingQuerybyorderidRequest) SetMainOrderId(mainOrderId int64) error {
    r.mainOrderId = mainOrderId
    r.Set("main_order_id", mainOrderId)
    return nil
}

func (r AlibabaLstMarketingQuerybyorderidRequest) GetMainOrderId() int64 {
    return r.mainOrderId
}

func (r *AlibabaLstMarketingQuerybyorderidRequest) SetSubOrderId(subOrderId int64) error {
    r.subOrderId = subOrderId
    r.Set("sub_order_id", subOrderId)
    return nil
}

func (r AlibabaLstMarketingQuerybyorderidRequest) GetSubOrderId() int64 {
    return r.subOrderId
}

