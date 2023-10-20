package tbuser

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUserSellerGetAPIRequest 查询卖家用户信息 API请求
// taobao.user.seller.get
//
// 查询卖家用户信息（只能查询有店铺的用户） 只能卖家类应用调用。
type TaobaoUserSellerGetAPIRequest struct {
	model.Params
	// 需要返回的字段列表，可选值为返回示例值中的可以看到的字段
	_fields []string
}

// NewTaobaoUserSellerGetRequest 初始化TaobaoUserSellerGetAPIRequest对象
func NewTaobaoUserSellerGetRequest() *TaobaoUserSellerGetAPIRequest {
	return &TaobaoUserSellerGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUserSellerGetAPIRequest) Reset() {
	r._fields = r._fields[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUserSellerGetAPIRequest) GetApiMethodName() string {
	return "taobao.user.seller.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUserSellerGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUserSellerGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 需要返回的字段列表，可选值为返回示例值中的可以看到的字段
func (r *TaobaoUserSellerGetAPIRequest) SetFields(_fields []string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoUserSellerGetAPIRequest) GetFields() []string {
	return r._fields
}

var poolTaobaoUserSellerGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUserSellerGetRequest()
	},
}

// GetTaobaoUserSellerGetRequest 从 sync.Pool 获取 TaobaoUserSellerGetAPIRequest
func GetTaobaoUserSellerGetAPIRequest() *TaobaoUserSellerGetAPIRequest {
	return poolTaobaoUserSellerGetAPIRequest.Get().(*TaobaoUserSellerGetAPIRequest)
}

// ReleaseTaobaoUserSellerGetAPIRequest 将 TaobaoUserSellerGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoUserSellerGetAPIRequest(v *TaobaoUserSellerGetAPIRequest) {
	v.Reset()
	poolTaobaoUserSellerGetAPIRequest.Put(v)
}
