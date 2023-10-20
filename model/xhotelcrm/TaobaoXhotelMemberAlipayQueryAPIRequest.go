package xhotelcrm

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelMemberAlipayQueryAPIRequest 希尔顿会员查询 API请求
// taobao.xhotel.member.alipay.query
//
// 希尔顿会员查询
type TaobaoXhotelMemberAlipayQueryAPIRequest struct {
	model.Params
	// 渠道
	_fpt string
	// 支付宝id
	_alipayUserId int64
}

// NewTaobaoXhotelMemberAlipayQueryRequest 初始化TaobaoXhotelMemberAlipayQueryAPIRequest对象
func NewTaobaoXhotelMemberAlipayQueryRequest() *TaobaoXhotelMemberAlipayQueryAPIRequest {
	return &TaobaoXhotelMemberAlipayQueryAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoXhotelMemberAlipayQueryAPIRequest) Reset() {
	r._fpt = ""
	r._alipayUserId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelMemberAlipayQueryAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.member.alipay.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelMemberAlipayQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelMemberAlipayQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFpt is Fpt Setter
// 渠道
func (r *TaobaoXhotelMemberAlipayQueryAPIRequest) SetFpt(_fpt string) error {
	r._fpt = _fpt
	r.Set("fpt", _fpt)
	return nil
}

// GetFpt Fpt Getter
func (r TaobaoXhotelMemberAlipayQueryAPIRequest) GetFpt() string {
	return r._fpt
}

// SetAlipayUserId is AlipayUserId Setter
// 支付宝id
func (r *TaobaoXhotelMemberAlipayQueryAPIRequest) SetAlipayUserId(_alipayUserId int64) error {
	r._alipayUserId = _alipayUserId
	r.Set("alipay_user_id", _alipayUserId)
	return nil
}

// GetAlipayUserId AlipayUserId Getter
func (r TaobaoXhotelMemberAlipayQueryAPIRequest) GetAlipayUserId() int64 {
	return r._alipayUserId
}

var poolTaobaoXhotelMemberAlipayQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoXhotelMemberAlipayQueryRequest()
	},
}

// GetTaobaoXhotelMemberAlipayQueryRequest 从 sync.Pool 获取 TaobaoXhotelMemberAlipayQueryAPIRequest
func GetTaobaoXhotelMemberAlipayQueryAPIRequest() *TaobaoXhotelMemberAlipayQueryAPIRequest {
	return poolTaobaoXhotelMemberAlipayQueryAPIRequest.Get().(*TaobaoXhotelMemberAlipayQueryAPIRequest)
}

// ReleaseTaobaoXhotelMemberAlipayQueryAPIRequest 将 TaobaoXhotelMemberAlipayQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoXhotelMemberAlipayQueryAPIRequest(v *TaobaoXhotelMemberAlipayQueryAPIRequest) {
	v.Reset()
	poolTaobaoXhotelMemberAlipayQueryAPIRequest.Put(v)
}
