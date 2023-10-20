package tbtrade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTradePostageUpdateAPIRequest 修改交易邮费价格 API请求
// taobao.trade.postage.update
//
// 修改订单邮费接口，通过传入订单编号和邮费价格，修改订单的邮费，返回修改时间modified,邮费post_fee,总费用total_fee。
// &lt;br/&gt; &lt;span style=&#34;color:red&#34;&gt; API取消加邮费功能通知：http://open.taobao.com/support/announcement_detail.htm?tid=24750&lt;/span&gt;
type TaobaoTradePostageUpdateAPIRequest struct {
	model.Params
	// 邮费价格(邮费单位是元）
	_postFee string
	// 主订单编号
	_tid int64
}

// NewTaobaoTradePostageUpdateRequest 初始化TaobaoTradePostageUpdateAPIRequest对象
func NewTaobaoTradePostageUpdateRequest() *TaobaoTradePostageUpdateAPIRequest {
	return &TaobaoTradePostageUpdateAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTradePostageUpdateAPIRequest) Reset() {
	r._postFee = ""
	r._tid = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTradePostageUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.trade.postage.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTradePostageUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTradePostageUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPostFee is PostFee Setter
// 邮费价格(邮费单位是元）
func (r *TaobaoTradePostageUpdateAPIRequest) SetPostFee(_postFee string) error {
	r._postFee = _postFee
	r.Set("post_fee", _postFee)
	return nil
}

// GetPostFee PostFee Getter
func (r TaobaoTradePostageUpdateAPIRequest) GetPostFee() string {
	return r._postFee
}

// SetTid is Tid Setter
// 主订单编号
func (r *TaobaoTradePostageUpdateAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoTradePostageUpdateAPIRequest) GetTid() int64 {
	return r._tid
}

var poolTaobaoTradePostageUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTradePostageUpdateRequest()
	},
}

// GetTaobaoTradePostageUpdateRequest 从 sync.Pool 获取 TaobaoTradePostageUpdateAPIRequest
func GetTaobaoTradePostageUpdateAPIRequest() *TaobaoTradePostageUpdateAPIRequest {
	return poolTaobaoTradePostageUpdateAPIRequest.Get().(*TaobaoTradePostageUpdateAPIRequest)
}

// ReleaseTaobaoTradePostageUpdateAPIRequest 将 TaobaoTradePostageUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoTradePostageUpdateAPIRequest(v *TaobaoTradePostageUpdateAPIRequest) {
	v.Reset()
	poolTaobaoTradePostageUpdateAPIRequest.Put(v)
}
