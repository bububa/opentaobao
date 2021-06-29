package tmallhk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫国际订单清关信息 API请求
tmall.hk.order.clearance.query

天猫国际订单清关信息查询
*/
type TmallHkOrderClearanceQueryRequest struct {
    model.Params
    // 交易主订单号
    bizOrderId   int64
    // 调用方业务身份(由国际侧配置提供给调用方)
    businessSymbol   string
}

// 初始化TmallHkOrderClearanceQueryRequest对象
func NewTmallHkOrderClearanceQueryRequest() *TmallHkOrderClearanceQueryRequest{
    return &TmallHkOrderClearanceQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallHkOrderClearanceQueryRequest) GetApiMethodName() string {
    return "tmall.hk.order.clearance.query"
}

// IRequest interface 方法, 获取API参数
func (r TmallHkOrderClearanceQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizOrderId Setter
// 交易主订单号
func (r *TmallHkOrderClearanceQueryRequest) SetBizOrderId(bizOrderId int64) error {
    r.bizOrderId = bizOrderId
    r.Set("biz_order_id", bizOrderId)
    return nil
}

// BizOrderId Getter
func (r TmallHkOrderClearanceQueryRequest) GetBizOrderId() int64 {
    return r.bizOrderId
}
// BusinessSymbol Setter
// 调用方业务身份(由国际侧配置提供给调用方)
func (r *TmallHkOrderClearanceQueryRequest) SetBusinessSymbol(businessSymbol string) error {
    r.businessSymbol = businessSymbol
    r.Set("business_symbol", businessSymbol)
    return nil
}

// BusinessSymbol Getter
func (r TmallHkOrderClearanceQueryRequest) GetBusinessSymbol() string {
    return r.businessSymbol
}
