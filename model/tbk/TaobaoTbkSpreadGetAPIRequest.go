package tbk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkSpreadGetAPIRequest 淘宝客-公用-长链转短链 API请求
// taobao.tbk.spread.get
//
// 输入一个原始的链接，转换得到指定的传播方式，如二维码，淘口令，短连接；
// 现阶段只支持短连接。
type TaobaoTbkSpreadGetAPIRequest struct {
	model.Params
	// 请求列表，内部包含多个url
	_requests []TbkSpreadRequest
}

// NewTaobaoTbkSpreadGetRequest 初始化TaobaoTbkSpreadGetAPIRequest对象
func NewTaobaoTbkSpreadGetRequest() *TaobaoTbkSpreadGetAPIRequest {
	return &TaobaoTbkSpreadGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTbkSpreadGetAPIRequest) Reset() {
	r._requests = r._requests[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTbkSpreadGetAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.spread.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTbkSpreadGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTbkSpreadGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequests is Requests Setter
// 请求列表，内部包含多个url
func (r *TaobaoTbkSpreadGetAPIRequest) SetRequests(_requests []TbkSpreadRequest) error {
	r._requests = _requests
	r.Set("requests", _requests)
	return nil
}

// GetRequests Requests Getter
func (r TaobaoTbkSpreadGetAPIRequest) GetRequests() []TbkSpreadRequest {
	return r._requests
}

var poolTaobaoTbkSpreadGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTbkSpreadGetRequest()
	},
}

// GetTaobaoTbkSpreadGetRequest 从 sync.Pool 获取 TaobaoTbkSpreadGetAPIRequest
func GetTaobaoTbkSpreadGetAPIRequest() *TaobaoTbkSpreadGetAPIRequest {
	return poolTaobaoTbkSpreadGetAPIRequest.Get().(*TaobaoTbkSpreadGetAPIRequest)
}

// ReleaseTaobaoTbkSpreadGetAPIRequest 将 TaobaoTbkSpreadGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoTbkSpreadGetAPIRequest(v *TaobaoTbkSpreadGetAPIRequest) {
	v.Reset()
	poolTaobaoTbkSpreadGetAPIRequest.Put(v)
}
