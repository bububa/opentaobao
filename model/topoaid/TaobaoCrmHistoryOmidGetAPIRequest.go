package topoaid

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmHistoryOmidGetAPIRequest 根据buyerNick获取omid API请求
// taobao.crm.history.omid.get
//
// 根据buyerNick获取ouid
type TaobaoCrmHistoryOmidGetAPIRequest struct {
	model.Params
	// 买家淘宝Nick
	_buyerNick string
}

// NewTaobaoCrmHistoryOmidGetRequest 初始化TaobaoCrmHistoryOmidGetAPIRequest对象
func NewTaobaoCrmHistoryOmidGetRequest() *TaobaoCrmHistoryOmidGetAPIRequest {
	return &TaobaoCrmHistoryOmidGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoCrmHistoryOmidGetAPIRequest) Reset() {
	r._buyerNick = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCrmHistoryOmidGetAPIRequest) GetApiMethodName() string {
	return "taobao.crm.history.omid.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCrmHistoryOmidGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoCrmHistoryOmidGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBuyerNick is BuyerNick Setter
// 买家淘宝Nick
func (r *TaobaoCrmHistoryOmidGetAPIRequest) SetBuyerNick(_buyerNick string) error {
	r._buyerNick = _buyerNick
	r.Set("buyer_nick", _buyerNick)
	return nil
}

// GetBuyerNick BuyerNick Getter
func (r TaobaoCrmHistoryOmidGetAPIRequest) GetBuyerNick() string {
	return r._buyerNick
}

var poolTaobaoCrmHistoryOmidGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoCrmHistoryOmidGetRequest()
	},
}

// GetTaobaoCrmHistoryOmidGetRequest 从 sync.Pool 获取 TaobaoCrmHistoryOmidGetAPIRequest
func GetTaobaoCrmHistoryOmidGetAPIRequest() *TaobaoCrmHistoryOmidGetAPIRequest {
	return poolTaobaoCrmHistoryOmidGetAPIRequest.Get().(*TaobaoCrmHistoryOmidGetAPIRequest)
}

// ReleaseTaobaoCrmHistoryOmidGetAPIRequest 将 TaobaoCrmHistoryOmidGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoCrmHistoryOmidGetAPIRequest(v *TaobaoCrmHistoryOmidGetAPIRequest) {
	v.Reset()
	poolTaobaoCrmHistoryOmidGetAPIRequest.Put(v)
}
