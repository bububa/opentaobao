package lstmarketing

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据订单查询营销信息 API请求
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

// 初始化AlibabaLstMarketingQuerybyorderidRequest对象
func NewAlibabaLstMarketingQuerybyorderidRequest() *AlibabaLstMarketingQuerybyorderidRequest{
    return &AlibabaLstMarketingQuerybyorderidRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLstMarketingQuerybyorderidRequest) GetApiMethodName() string {
    return "alibaba.lst.marketing.querybyorderid"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLstMarketingQuerybyorderidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MainOrderId Setter
// 主订单
func (r *AlibabaLstMarketingQuerybyorderidRequest) SetMainOrderId(mainOrderId int64) error {
    r.mainOrderId = mainOrderId
    r.Set("main_order_id", mainOrderId)
    return nil
}

// MainOrderId Getter
func (r AlibabaLstMarketingQuerybyorderidRequest) GetMainOrderId() int64 {
    return r.mainOrderId
}
// SubOrderId Setter
// 子订单
func (r *AlibabaLstMarketingQuerybyorderidRequest) SetSubOrderId(subOrderId int64) error {
    r.subOrderId = subOrderId
    r.Set("sub_order_id", subOrderId)
    return nil
}

// SubOrderId Getter
func (r AlibabaLstMarketingQuerybyorderidRequest) GetSubOrderId() int64 {
    return r.subOrderId
}
