package tmallchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通过主采购单号查询采购单 API请求
tmall.channel.trade.order.get

通过主采购单号查询采购单
*/
type TmallChannelTradeOrderGetRequest struct {
    model.Params
    // 主采购单ID
    mainPurchaseOrderNo   int64
    // 是否包含子采购单
    isIncludeSubOrder   bool
    // 是否包含主采购单（针对特殊业务）
    isIncludeMainOrder   bool
    // 是否包含物流信息
    isIncludeLogistics   bool
}

// 初始化TmallChannelTradeOrderGetRequest对象
func NewTmallChannelTradeOrderGetRequest() *TmallChannelTradeOrderGetRequest{
    return &TmallChannelTradeOrderGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallChannelTradeOrderGetRequest) GetApiMethodName() string {
    return "tmall.channel.trade.order.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallChannelTradeOrderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MainPurchaseOrderNo Setter
// 主采购单ID
func (r *TmallChannelTradeOrderGetRequest) SetMainPurchaseOrderNo(mainPurchaseOrderNo int64) error {
    r.mainPurchaseOrderNo = mainPurchaseOrderNo
    r.Set("main_purchase_order_no", mainPurchaseOrderNo)
    return nil
}

// MainPurchaseOrderNo Getter
func (r TmallChannelTradeOrderGetRequest) GetMainPurchaseOrderNo() int64 {
    return r.mainPurchaseOrderNo
}
// IsIncludeSubOrder Setter
// 是否包含子采购单
func (r *TmallChannelTradeOrderGetRequest) SetIsIncludeSubOrder(isIncludeSubOrder bool) error {
    r.isIncludeSubOrder = isIncludeSubOrder
    r.Set("is_include_sub_order", isIncludeSubOrder)
    return nil
}

// IsIncludeSubOrder Getter
func (r TmallChannelTradeOrderGetRequest) GetIsIncludeSubOrder() bool {
    return r.isIncludeSubOrder
}
// IsIncludeMainOrder Setter
// 是否包含主采购单（针对特殊业务）
func (r *TmallChannelTradeOrderGetRequest) SetIsIncludeMainOrder(isIncludeMainOrder bool) error {
    r.isIncludeMainOrder = isIncludeMainOrder
    r.Set("is_include_main_order", isIncludeMainOrder)
    return nil
}

// IsIncludeMainOrder Getter
func (r TmallChannelTradeOrderGetRequest) GetIsIncludeMainOrder() bool {
    return r.isIncludeMainOrder
}
// IsIncludeLogistics Setter
// 是否包含物流信息
func (r *TmallChannelTradeOrderGetRequest) SetIsIncludeLogistics(isIncludeLogistics bool) error {
    r.isIncludeLogistics = isIncludeLogistics
    r.Set("is_include_logistics", isIncludeLogistics)
    return nil
}

// IsIncludeLogistics Getter
func (r TmallChannelTradeOrderGetRequest) GetIsIncludeLogistics() bool {
    return r.isIncludeLogistics
}
