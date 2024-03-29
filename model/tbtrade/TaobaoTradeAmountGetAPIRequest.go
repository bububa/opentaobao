package tbtrade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTradeAmountGetAPIRequest 交易帐务查询 API请求
// taobao.trade.amount.get
//
// 卖家查询该笔交易的资金帐务相关的数据；
// 1. 只供卖家使用，买家不可使用
// 2. 可查询所有的状态的交易，但不同状态时交易的相关数据可能会有不同
type TaobaoTradeAmountGetAPIRequest struct {
	model.Params
	// 订单帐务详情需要返回的字段信息，可选值如下：1. TradeAmount中可指定的fields：tid,alipay_no,created,pay_time,end_time,total_fee,payment,post_fee,cod_fee,commission_fee,buyer_obtain_point_fee2. OrderAmount中可指定的fields：order_amounts.oid,order_amounts.title,order_amounts.num_iid,order_amounts.sku_properties_name,order_amounts.sku_id,order_amounts.num,order_amounts.price,order_amounts.discount_fee,order_amounts.adjust_fee,order_amounts.payment,order_amounts.promotion_name3. order_amounts(返回OrderAmount的所有内容)4. promotion_details(指定该值会返回主订单的promotion_details中除id之外的所有字段)
	_fields string
	// 交易编号
	_tid int64
}

// NewTaobaoTradeAmountGetRequest 初始化TaobaoTradeAmountGetAPIRequest对象
func NewTaobaoTradeAmountGetRequest() *TaobaoTradeAmountGetAPIRequest {
	return &TaobaoTradeAmountGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTradeAmountGetAPIRequest) Reset() {
	r._fields = ""
	r._tid = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTradeAmountGetAPIRequest) GetApiMethodName() string {
	return "taobao.trade.amount.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTradeAmountGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTradeAmountGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 订单帐务详情需要返回的字段信息，可选值如下：1. TradeAmount中可指定的fields：tid,alipay_no,created,pay_time,end_time,total_fee,payment,post_fee,cod_fee,commission_fee,buyer_obtain_point_fee2. OrderAmount中可指定的fields：order_amounts.oid,order_amounts.title,order_amounts.num_iid,order_amounts.sku_properties_name,order_amounts.sku_id,order_amounts.num,order_amounts.price,order_amounts.discount_fee,order_amounts.adjust_fee,order_amounts.payment,order_amounts.promotion_name3. order_amounts(返回OrderAmount的所有内容)4. promotion_details(指定该值会返回主订单的promotion_details中除id之外的所有字段)
func (r *TaobaoTradeAmountGetAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoTradeAmountGetAPIRequest) GetFields() string {
	return r._fields
}

// SetTid is Tid Setter
// 交易编号
func (r *TaobaoTradeAmountGetAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoTradeAmountGetAPIRequest) GetTid() int64 {
	return r._tid
}

var poolTaobaoTradeAmountGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTradeAmountGetRequest()
	},
}

// GetTaobaoTradeAmountGetRequest 从 sync.Pool 获取 TaobaoTradeAmountGetAPIRequest
func GetTaobaoTradeAmountGetAPIRequest() *TaobaoTradeAmountGetAPIRequest {
	return poolTaobaoTradeAmountGetAPIRequest.Get().(*TaobaoTradeAmountGetAPIRequest)
}

// ReleaseTaobaoTradeAmountGetAPIRequest 将 TaobaoTradeAmountGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoTradeAmountGetAPIRequest(v *TaobaoTradeAmountGetAPIRequest) {
	v.Reset()
	poolTaobaoTradeAmountGetAPIRequest.Put(v)
}
