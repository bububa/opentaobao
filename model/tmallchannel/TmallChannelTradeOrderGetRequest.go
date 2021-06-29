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
    _mainPurchaseOrderNo   int64
    // 是否包含子采购单
    _isIncludeSubOrder   bool
    // 是否包含主采购单（针对特殊业务）
    _isIncludeMainOrder   bool
    // 是否包含物流信息
    _isIncludeLogistics   bool
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
func (r *TmallChannelTradeOrderGetRequest) SetMainPurchaseOrderNo(_mainPurchaseOrderNo int64) error {
    r._mainPurchaseOrderNo = _mainPurchaseOrderNo
    r.Set("main_purchase_order_no", _mainPurchaseOrderNo)
    return nil
}

// MainPurchaseOrderNo Getter
func (r TmallChannelTradeOrderGetRequest) GetMainPurchaseOrderNo() int64 {
    return r._mainPurchaseOrderNo
}
// IsIncludeSubOrder Setter
// 是否包含子采购单
func (r *TmallChannelTradeOrderGetRequest) SetIsIncludeSubOrder(_isIncludeSubOrder bool) error {
    r._isIncludeSubOrder = _isIncludeSubOrder
    r.Set("is_include_sub_order", _isIncludeSubOrder)
    return nil
}

// IsIncludeSubOrder Getter
func (r TmallChannelTradeOrderGetRequest) GetIsIncludeSubOrder() bool {
    return r._isIncludeSubOrder
}
// IsIncludeMainOrder Setter
// 是否包含主采购单（针对特殊业务）
func (r *TmallChannelTradeOrderGetRequest) SetIsIncludeMainOrder(_isIncludeMainOrder bool) error {
    r._isIncludeMainOrder = _isIncludeMainOrder
    r.Set("is_include_main_order", _isIncludeMainOrder)
    return nil
}

// IsIncludeMainOrder Getter
func (r TmallChannelTradeOrderGetRequest) GetIsIncludeMainOrder() bool {
    return r._isIncludeMainOrder
}
// IsIncludeLogistics Setter
// 是否包含物流信息
func (r *TmallChannelTradeOrderGetRequest) SetIsIncludeLogistics(_isIncludeLogistics bool) error {
    r._isIncludeLogistics = _isIncludeLogistics
    r.Set("is_include_logistics", _isIncludeLogistics)
    return nil
}

// IsIncludeLogistics Getter
func (r TmallChannelTradeOrderGetRequest) GetIsIncludeLogistics() bool {
    return r._isIncludeLogistics
}
