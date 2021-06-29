package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
交易帐务查询 API请求
taobao.trade.amount.get

卖家查询该笔交易的资金帐务相关的数据；
1. 只供卖家使用，买家不可使用
2. 可查询所有的状态的交易，但不同状态时交易的相关数据可能会有不同
*/
type TaobaoTradeAmountGetRequest struct {
    model.Params
    // 交易编号
    _tid   int64
    // 订单帐务详情需要返回的字段信息，可选值如下：1. TradeAmount中可指定的fields：tid,alipay_no,created,pay_time,end_time,total_fee,payment,post_fee,cod_fee,commission_fee,buyer_obtain_point_fee2. OrderAmount中可指定的fields：order_amounts.oid,order_amounts.title,order_amounts.num_iid,order_amounts.sku_properties_name,order_amounts.sku_id,order_amounts.num,order_amounts.price,order_amounts.discount_fee,order_amounts.adjust_fee,order_amounts.payment,order_amounts.promotion_name3. order_amounts(返回OrderAmount的所有内容)4. promotion_details(指定该值会返回主订单的promotion_details中除id之外的所有字段)
    _fields   string
}

// 初始化TaobaoTradeAmountGetRequest对象
func NewTaobaoTradeAmountGetRequest() *TaobaoTradeAmountGetRequest{
    return &TaobaoTradeAmountGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTradeAmountGetRequest) GetApiMethodName() string {
    return "taobao.trade.amount.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTradeAmountGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Tid Setter
// 交易编号
func (r *TaobaoTradeAmountGetRequest) SetTid(_tid int64) error {
    r._tid = _tid
    r.Set("tid", _tid)
    return nil
}

// Tid Getter
func (r TaobaoTradeAmountGetRequest) GetTid() int64 {
    return r._tid
}
// Fields Setter
// 订单帐务详情需要返回的字段信息，可选值如下：1. TradeAmount中可指定的fields：tid,alipay_no,created,pay_time,end_time,total_fee,payment,post_fee,cod_fee,commission_fee,buyer_obtain_point_fee2. OrderAmount中可指定的fields：order_amounts.oid,order_amounts.title,order_amounts.num_iid,order_amounts.sku_properties_name,order_amounts.sku_id,order_amounts.num,order_amounts.price,order_amounts.discount_fee,order_amounts.adjust_fee,order_amounts.payment,order_amounts.promotion_name3. order_amounts(返回OrderAmount的所有内容)4. promotion_details(指定该值会返回主订单的promotion_details中除id之外的所有字段)
func (r *TaobaoTradeAmountGetRequest) SetFields(_fields string) error {
    r._fields = _fields
    r.Set("fields", _fields)
    return nil
}

// Fields Getter
func (r TaobaoTradeAmountGetRequest) GetFields() string {
    return r._fields
}
