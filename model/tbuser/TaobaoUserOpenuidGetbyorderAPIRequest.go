package tbuser

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUserOpenuidGetbyorderAPIRequest 根据订单获取买家openuid API请求
// taobao.user.openuid.getbyorder
//
// 根据订单获取买家openuid，最大查询30个
type TaobaoUserOpenuidGetbyorderAPIRequest struct {
	model.Params
	// 买家订单列表
	_tidInfos *BuyerOrder
}

// NewTaobaoUserOpenuidGetbyorderRequest 初始化TaobaoUserOpenuidGetbyorderAPIRequest对象
func NewTaobaoUserOpenuidGetbyorderRequest() *TaobaoUserOpenuidGetbyorderAPIRequest {
	return &TaobaoUserOpenuidGetbyorderAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUserOpenuidGetbyorderAPIRequest) Reset() {
	r._tidInfos = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUserOpenuidGetbyorderAPIRequest) GetApiMethodName() string {
	return "taobao.user.openuid.getbyorder"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUserOpenuidGetbyorderAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUserOpenuidGetbyorderAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTidInfos is TidInfos Setter
// 买家订单列表
func (r *TaobaoUserOpenuidGetbyorderAPIRequest) SetTidInfos(_tidInfos *BuyerOrder) error {
	r._tidInfos = _tidInfos
	r.Set("tid_infos", _tidInfos)
	return nil
}

// GetTidInfos TidInfos Getter
func (r TaobaoUserOpenuidGetbyorderAPIRequest) GetTidInfos() *BuyerOrder {
	return r._tidInfos
}

var poolTaobaoUserOpenuidGetbyorderAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUserOpenuidGetbyorderRequest()
	},
}

// GetTaobaoUserOpenuidGetbyorderRequest 从 sync.Pool 获取 TaobaoUserOpenuidGetbyorderAPIRequest
func GetTaobaoUserOpenuidGetbyorderAPIRequest() *TaobaoUserOpenuidGetbyorderAPIRequest {
	return poolTaobaoUserOpenuidGetbyorderAPIRequest.Get().(*TaobaoUserOpenuidGetbyorderAPIRequest)
}

// ReleaseTaobaoUserOpenuidGetbyorderAPIRequest 将 TaobaoUserOpenuidGetbyorderAPIRequest 放入 sync.Pool
func ReleaseTaobaoUserOpenuidGetbyorderAPIRequest(v *TaobaoUserOpenuidGetbyorderAPIRequest) {
	v.Reset()
	poolTaobaoUserOpenuidGetbyorderAPIRequest.Put(v)
}
