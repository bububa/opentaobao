package tmallnr

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
区域零售订单获取取件码 API请求
tmall.nr.order.logis.info

区域零售订单获取取件码，方便商家系统接入，获取取件码打印信息进行打印。
*/
type TmallNrOrderLogisInfoRequest struct {
    model.Params
    // 卖家ID
    sellerId   int64
    // 主订单号
    mainOrderIds   []int64
    // 来源标识
    channel   string
}

// 初始化TmallNrOrderLogisInfoRequest对象
func NewTmallNrOrderLogisInfoRequest() *TmallNrOrderLogisInfoRequest{
    return &TmallNrOrderLogisInfoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallNrOrderLogisInfoRequest) GetApiMethodName() string {
    return "tmall.nr.order.logis.info"
}

// IRequest interface 方法, 获取API参数
func (r TmallNrOrderLogisInfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SellerId Setter
// 卖家ID
func (r *TmallNrOrderLogisInfoRequest) SetSellerId(sellerId int64) error {
    r.sellerId = sellerId
    r.Set("seller_id", sellerId)
    return nil
}

// SellerId Getter
func (r TmallNrOrderLogisInfoRequest) GetSellerId() int64 {
    return r.sellerId
}
// MainOrderIds Setter
// 主订单号
func (r *TmallNrOrderLogisInfoRequest) SetMainOrderIds(mainOrderIds []int64) error {
    r.mainOrderIds = mainOrderIds
    r.Set("main_order_ids", mainOrderIds)
    return nil
}

// MainOrderIds Getter
func (r TmallNrOrderLogisInfoRequest) GetMainOrderIds() []int64 {
    return r.mainOrderIds
}
// Channel Setter
// 来源标识
func (r *TmallNrOrderLogisInfoRequest) SetChannel(channel string) error {
    r.channel = channel
    r.Set("channel", channel)
    return nil
}

// Channel Getter
func (r TmallNrOrderLogisInfoRequest) GetChannel() string {
    return r.channel
}
