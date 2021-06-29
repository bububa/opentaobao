package tmallchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商查询退款单 API请求
tmall.channel.trade.refundorder.gets

供应商分页查询退款单
*/
type TmallChannelTradeRefundorderGetsRequest struct {
    model.Params
    // 退款单号
    _refundId   int64
    // 采购单号
    _mainChannelOrderNo   string
    // 每页数据条数
    _pageSize   int64
    // 页码
    _pageNumber   int64
}

// 初始化TmallChannelTradeRefundorderGetsRequest对象
func NewTmallChannelTradeRefundorderGetsRequest() *TmallChannelTradeRefundorderGetsRequest{
    return &TmallChannelTradeRefundorderGetsRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallChannelTradeRefundorderGetsRequest) GetApiMethodName() string {
    return "tmall.channel.trade.refundorder.gets"
}

// IRequest interface 方法, 获取API参数
func (r TmallChannelTradeRefundorderGetsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefundId Setter
// 退款单号
func (r *TmallChannelTradeRefundorderGetsRequest) SetRefundId(_refundId int64) error {
    r._refundId = _refundId
    r.Set("refund_id", _refundId)
    return nil
}

// RefundId Getter
func (r TmallChannelTradeRefundorderGetsRequest) GetRefundId() int64 {
    return r._refundId
}
// MainChannelOrderNo Setter
// 采购单号
func (r *TmallChannelTradeRefundorderGetsRequest) SetMainChannelOrderNo(_mainChannelOrderNo string) error {
    r._mainChannelOrderNo = _mainChannelOrderNo
    r.Set("main_channel_order_no", _mainChannelOrderNo)
    return nil
}

// MainChannelOrderNo Getter
func (r TmallChannelTradeRefundorderGetsRequest) GetMainChannelOrderNo() string {
    return r._mainChannelOrderNo
}
// PageSize Setter
// 每页数据条数
func (r *TmallChannelTradeRefundorderGetsRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TmallChannelTradeRefundorderGetsRequest) GetPageSize() int64 {
    return r._pageSize
}
// PageNumber Setter
// 页码
func (r *TmallChannelTradeRefundorderGetsRequest) SetPageNumber(_pageNumber int64) error {
    r._pageNumber = _pageNumber
    r.Set("page_number", _pageNumber)
    return nil
}

// PageNumber Getter
func (r TmallChannelTradeRefundorderGetsRequest) GetPageNumber() int64 {
    return r._pageNumber
}
