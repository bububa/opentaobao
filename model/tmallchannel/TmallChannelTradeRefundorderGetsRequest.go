package tmallchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商查询退款单 APIRequest
tmall.channel.trade.refundorder.gets

供应商分页查询退款单
*/
type TmallChannelTradeRefundorderGetsRequest struct {
    model.Params

    // 退款单号
    refundId   int64 

    // 采购单号
    mainChannelOrderNo   string 

    // 每页数据条数
    pageSize   int64 

    // 页码
    pageNumber   int64 

}

func NewTmallChannelTradeRefundorderGetsRequest() *TmallChannelTradeRefundorderGetsRequest{
    return &TmallChannelTradeRefundorderGetsRequest{
        Params: model.NewParams(),
    }
}

func (r TmallChannelTradeRefundorderGetsRequest) GetApiMethodName() string {
    return "tmall.channel.trade.refundorder.gets"
}

func (r TmallChannelTradeRefundorderGetsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallChannelTradeRefundorderGetsRequest) SetRefundId(refundId int64) error {
    r.refundId = refundId
    r.Set("refund_id", refundId)
    return nil
}

func (r TmallChannelTradeRefundorderGetsRequest) GetRefundId() int64 {
    return r.refundId
}

func (r *TmallChannelTradeRefundorderGetsRequest) SetMainChannelOrderNo(mainChannelOrderNo string) error {
    r.mainChannelOrderNo = mainChannelOrderNo
    r.Set("main_channel_order_no", mainChannelOrderNo)
    return nil
}

func (r TmallChannelTradeRefundorderGetsRequest) GetMainChannelOrderNo() string {
    return r.mainChannelOrderNo
}

func (r *TmallChannelTradeRefundorderGetsRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TmallChannelTradeRefundorderGetsRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TmallChannelTradeRefundorderGetsRequest) SetPageNumber(pageNumber int64) error {
    r.pageNumber = pageNumber
    r.Set("page_number", pageNumber)
    return nil
}

func (r TmallChannelTradeRefundorderGetsRequest) GetPageNumber() int64 {
    return r.pageNumber
}

