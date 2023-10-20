package topoaid

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmHistoryOuidGetAPIRequest 根据buyerNick获取ouid API请求
// taobao.crm.history.ouid.get
//
// 根据buyerNick获取ouid
type TaobaoCrmHistoryOuidGetAPIRequest struct {
	model.Params
	// 买家淘宝Nick
	_buyerNick string
}

// NewTaobaoCrmHistoryOuidGetRequest 初始化TaobaoCrmHistoryOuidGetAPIRequest对象
func NewTaobaoCrmHistoryOuidGetRequest() *TaobaoCrmHistoryOuidGetAPIRequest {
	return &TaobaoCrmHistoryOuidGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoCrmHistoryOuidGetAPIRequest) Reset() {
	r._buyerNick = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCrmHistoryOuidGetAPIRequest) GetApiMethodName() string {
	return "taobao.crm.history.ouid.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCrmHistoryOuidGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoCrmHistoryOuidGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBuyerNick is BuyerNick Setter
// 买家淘宝Nick
func (r *TaobaoCrmHistoryOuidGetAPIRequest) SetBuyerNick(_buyerNick string) error {
	r._buyerNick = _buyerNick
	r.Set("buyer_nick", _buyerNick)
	return nil
}

// GetBuyerNick BuyerNick Getter
func (r TaobaoCrmHistoryOuidGetAPIRequest) GetBuyerNick() string {
	return r._buyerNick
}

var poolTaobaoCrmHistoryOuidGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoCrmHistoryOuidGetRequest()
	},
}

// GetTaobaoCrmHistoryOuidGetRequest 从 sync.Pool 获取 TaobaoCrmHistoryOuidGetAPIRequest
func GetTaobaoCrmHistoryOuidGetAPIRequest() *TaobaoCrmHistoryOuidGetAPIRequest {
	return poolTaobaoCrmHistoryOuidGetAPIRequest.Get().(*TaobaoCrmHistoryOuidGetAPIRequest)
}

// ReleaseTaobaoCrmHistoryOuidGetAPIRequest 将 TaobaoCrmHistoryOuidGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoCrmHistoryOuidGetAPIRequest(v *TaobaoCrmHistoryOuidGetAPIRequest) {
	v.Reset()
	poolTaobaoCrmHistoryOuidGetAPIRequest.Put(v)
}
