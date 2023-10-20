package trade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWttTradeServiceGetAPIRequest 获取网厅号卡垂直标信息 API请求
// taobao.wtt.trade.service.get
//
// 查询网厅订单信息
type TaobaoWttTradeServiceGetAPIRequest struct {
	model.Params
	// 订单ID
	_bizOrder int64
}

// NewTaobaoWttTradeServiceGetRequest 初始化TaobaoWttTradeServiceGetAPIRequest对象
func NewTaobaoWttTradeServiceGetRequest() *TaobaoWttTradeServiceGetAPIRequest {
	return &TaobaoWttTradeServiceGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoWttTradeServiceGetAPIRequest) Reset() {
	r._bizOrder = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWttTradeServiceGetAPIRequest) GetApiMethodName() string {
	return "taobao.wtt.trade.service.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWttTradeServiceGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWttTradeServiceGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizOrder is BizOrder Setter
// 订单ID
func (r *TaobaoWttTradeServiceGetAPIRequest) SetBizOrder(_bizOrder int64) error {
	r._bizOrder = _bizOrder
	r.Set("biz_order", _bizOrder)
	return nil
}

// GetBizOrder BizOrder Getter
func (r TaobaoWttTradeServiceGetAPIRequest) GetBizOrder() int64 {
	return r._bizOrder
}

var poolTaobaoWttTradeServiceGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoWttTradeServiceGetRequest()
	},
}

// GetTaobaoWttTradeServiceGetRequest 从 sync.Pool 获取 TaobaoWttTradeServiceGetAPIRequest
func GetTaobaoWttTradeServiceGetAPIRequest() *TaobaoWttTradeServiceGetAPIRequest {
	return poolTaobaoWttTradeServiceGetAPIRequest.Get().(*TaobaoWttTradeServiceGetAPIRequest)
}

// ReleaseTaobaoWttTradeServiceGetAPIRequest 将 TaobaoWttTradeServiceGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoWttTradeServiceGetAPIRequest(v *TaobaoWttTradeServiceGetAPIRequest) {
	v.Reset()
	poolTaobaoWttTradeServiceGetAPIRequest.Put(v)
}
