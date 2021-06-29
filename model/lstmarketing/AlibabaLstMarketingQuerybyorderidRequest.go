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
    _mainOrderId   int64
    // 子订单
    _subOrderId   int64
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
func (r *AlibabaLstMarketingQuerybyorderidRequest) SetMainOrderId(_mainOrderId int64) error {
    r._mainOrderId = _mainOrderId
    r.Set("main_order_id", _mainOrderId)
    return nil
}

// MainOrderId Getter
func (r AlibabaLstMarketingQuerybyorderidRequest) GetMainOrderId() int64 {
    return r._mainOrderId
}
// SubOrderId Setter
// 子订单
func (r *AlibabaLstMarketingQuerybyorderidRequest) SetSubOrderId(_subOrderId int64) error {
    r._subOrderId = _subOrderId
    r.Set("sub_order_id", _subOrderId)
    return nil
}

// SubOrderId Getter
func (r AlibabaLstMarketingQuerybyorderidRequest) GetSubOrderId() int64 {
    return r._subOrderId
}
