package tmallchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询发货单列表 APIRequest
tmall.channel.trade.deliverorder.gets

查询发货单列表
*/
type TmallChannelTradeDeliverorderGetsRequest struct {
    model.Params

    // 发货单单号
    mainDeliverOrderNo   int64 

    // 发货单状态列表
    orderStatusList   []int64 

    // 是否包括子发货单
    isIncludeSubOrder   bool 

    // 每页显示数量
    pageSize   int64 

    // 查询第几页
    pageNumber   int64 

    // 是否分页查询
    needPagination   bool 

    // 渠道
    channel   int64 

}

func NewTmallChannelTradeDeliverorderGetsRequest() *TmallChannelTradeDeliverorderGetsRequest{
    return &TmallChannelTradeDeliverorderGetsRequest{
        Params: model.NewParams(),
    }
}

func (r TmallChannelTradeDeliverorderGetsRequest) GetApiMethodName() string {
    return "tmall.channel.trade.deliverorder.gets"
}

func (r TmallChannelTradeDeliverorderGetsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallChannelTradeDeliverorderGetsRequest) SetMainDeliverOrderNo(mainDeliverOrderNo int64) error {
    r.mainDeliverOrderNo = mainDeliverOrderNo
    r.Set("main_deliver_order_no", mainDeliverOrderNo)
    return nil
}

func (r TmallChannelTradeDeliverorderGetsRequest) GetMainDeliverOrderNo() int64 {
    return r.mainDeliverOrderNo
}

func (r *TmallChannelTradeDeliverorderGetsRequest) SetOrderStatusList(orderStatusList []int64) error {
    r.orderStatusList = orderStatusList
    r.Set("order_status_list", orderStatusList)
    return nil
}

func (r TmallChannelTradeDeliverorderGetsRequest) GetOrderStatusList() []int64 {
    return r.orderStatusList
}

func (r *TmallChannelTradeDeliverorderGetsRequest) SetIsIncludeSubOrder(isIncludeSubOrder bool) error {
    r.isIncludeSubOrder = isIncludeSubOrder
    r.Set("is_include_sub_order", isIncludeSubOrder)
    return nil
}

func (r TmallChannelTradeDeliverorderGetsRequest) GetIsIncludeSubOrder() bool {
    return r.isIncludeSubOrder
}

func (r *TmallChannelTradeDeliverorderGetsRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TmallChannelTradeDeliverorderGetsRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TmallChannelTradeDeliverorderGetsRequest) SetPageNumber(pageNumber int64) error {
    r.pageNumber = pageNumber
    r.Set("page_number", pageNumber)
    return nil
}

func (r TmallChannelTradeDeliverorderGetsRequest) GetPageNumber() int64 {
    return r.pageNumber
}

func (r *TmallChannelTradeDeliverorderGetsRequest) SetNeedPagination(needPagination bool) error {
    r.needPagination = needPagination
    r.Set("need_pagination", needPagination)
    return nil
}

func (r TmallChannelTradeDeliverorderGetsRequest) GetNeedPagination() bool {
    return r.needPagination
}

func (r *TmallChannelTradeDeliverorderGetsRequest) SetChannel(channel int64) error {
    r.channel = channel
    r.Set("channel", channel)
    return nil
}

func (r TmallChannelTradeDeliverorderGetsRequest) GetChannel() int64 {
    return r.channel
}

