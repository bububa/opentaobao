package tbtrade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTradeCloseAPIRequest 卖家关闭一笔交易 API请求
// taobao.trade.close
//
// 关闭一笔订单，可以是主订单或子订单。当订单从创建到关闭时间小于10s的时候，会报“CLOSE_TRADE_TOO_FAST”错误。
type TaobaoTradeCloseAPIRequest struct {
	model.Params
	// 交易关闭原因。可以选择的理由有：1.未及时付款2、买家不想买了3、买家信息填写错误，重新拍4、恶意买家/同行捣乱5、缺货6、买家拍错了7、同城见面交易
	_closeReason string
	// 主订单或子订单编号。
	_tid int64
}

// NewTaobaoTradeCloseRequest 初始化TaobaoTradeCloseAPIRequest对象
func NewTaobaoTradeCloseRequest() *TaobaoTradeCloseAPIRequest {
	return &TaobaoTradeCloseAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTradeCloseAPIRequest) Reset() {
	r._closeReason = ""
	r._tid = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTradeCloseAPIRequest) GetApiMethodName() string {
	return "taobao.trade.close"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTradeCloseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTradeCloseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCloseReason is CloseReason Setter
// 交易关闭原因。可以选择的理由有：1.未及时付款2、买家不想买了3、买家信息填写错误，重新拍4、恶意买家/同行捣乱5、缺货6、买家拍错了7、同城见面交易
func (r *TaobaoTradeCloseAPIRequest) SetCloseReason(_closeReason string) error {
	r._closeReason = _closeReason
	r.Set("close_reason", _closeReason)
	return nil
}

// GetCloseReason CloseReason Getter
func (r TaobaoTradeCloseAPIRequest) GetCloseReason() string {
	return r._closeReason
}

// SetTid is Tid Setter
// 主订单或子订单编号。
func (r *TaobaoTradeCloseAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoTradeCloseAPIRequest) GetTid() int64 {
	return r._tid
}

var poolTaobaoTradeCloseAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTradeCloseRequest()
	},
}

// GetTaobaoTradeCloseRequest 从 sync.Pool 获取 TaobaoTradeCloseAPIRequest
func GetTaobaoTradeCloseAPIRequest() *TaobaoTradeCloseAPIRequest {
	return poolTaobaoTradeCloseAPIRequest.Get().(*TaobaoTradeCloseAPIRequest)
}

// ReleaseTaobaoTradeCloseAPIRequest 将 TaobaoTradeCloseAPIRequest 放入 sync.Pool
func ReleaseTaobaoTradeCloseAPIRequest(v *TaobaoTradeCloseAPIRequest) {
	v.Reset()
	poolTaobaoTradeCloseAPIRequest.Put(v)
}
