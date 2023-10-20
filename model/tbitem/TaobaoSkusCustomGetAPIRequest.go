package tbitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSkusCustomGetAPIRequest 根据外部ID取商品SKU API请求
// taobao.skus.custom.get
//
// 跟据卖家设定的Sku的外部id获取商品，如果一个outer_id对应多个Sku会返回所有符合条件的sku &lt;br/&gt;这个Sku所属卖家从传入的session中获取，需要session绑定(注：iid标签里是num_iid的值，可以用作num_iid使用)
type TaobaoSkusCustomGetAPIRequest struct {
	model.Params
	// 需返回的字段列表。可选值：Sku结构体中的所有字段；字段之间用“,”隔开
	_fields string
	// Sku的外部商家ID
	_outerId string
}

// NewTaobaoSkusCustomGetRequest 初始化TaobaoSkusCustomGetAPIRequest对象
func NewTaobaoSkusCustomGetRequest() *TaobaoSkusCustomGetAPIRequest {
	return &TaobaoSkusCustomGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSkusCustomGetAPIRequest) Reset() {
	r._fields = ""
	r._outerId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSkusCustomGetAPIRequest) GetApiMethodName() string {
	return "taobao.skus.custom.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSkusCustomGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSkusCustomGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 需返回的字段列表。可选值：Sku结构体中的所有字段；字段之间用“,”隔开
func (r *TaobaoSkusCustomGetAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoSkusCustomGetAPIRequest) GetFields() string {
	return r._fields
}

// SetOuterId is OuterId Setter
// Sku的外部商家ID
func (r *TaobaoSkusCustomGetAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TaobaoSkusCustomGetAPIRequest) GetOuterId() string {
	return r._outerId
}

var poolTaobaoSkusCustomGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSkusCustomGetRequest()
	},
}

// GetTaobaoSkusCustomGetRequest 从 sync.Pool 获取 TaobaoSkusCustomGetAPIRequest
func GetTaobaoSkusCustomGetAPIRequest() *TaobaoSkusCustomGetAPIRequest {
	return poolTaobaoSkusCustomGetAPIRequest.Get().(*TaobaoSkusCustomGetAPIRequest)
}

// ReleaseTaobaoSkusCustomGetAPIRequest 将 TaobaoSkusCustomGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoSkusCustomGetAPIRequest(v *TaobaoSkusCustomGetAPIRequest) {
	v.Reset()
	poolTaobaoSkusCustomGetAPIRequest.Put(v)
}
