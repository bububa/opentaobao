package qimen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenItemsSynchronizeAPIRequest 商品同步接口 (批量) API请求
// taobao.qimen.items.synchronize
//
// ERP调用奇门的接口,批量同步商品信息给WMS
type TaobaoQimenItemsSynchronizeAPIRequest struct {
	model.Params
	//
	_request *ItemsSynchronizeRequest
}

// NewTaobaoQimenItemsSynchronizeRequest 初始化TaobaoQimenItemsSynchronizeAPIRequest对象
func NewTaobaoQimenItemsSynchronizeRequest() *TaobaoQimenItemsSynchronizeAPIRequest {
	return &TaobaoQimenItemsSynchronizeAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoQimenItemsSynchronizeAPIRequest) Reset() {
	r._request = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenItemsSynchronizeAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.items.synchronize"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenItemsSynchronizeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenItemsSynchronizeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoQimenItemsSynchronizeAPIRequest) SetRequest(_request *ItemsSynchronizeRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenItemsSynchronizeAPIRequest) GetRequest() *ItemsSynchronizeRequest {
	return r._request
}

var poolTaobaoQimenItemsSynchronizeAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoQimenItemsSynchronizeRequest()
	},
}

// GetTaobaoQimenItemsSynchronizeRequest 从 sync.Pool 获取 TaobaoQimenItemsSynchronizeAPIRequest
func GetTaobaoQimenItemsSynchronizeAPIRequest() *TaobaoQimenItemsSynchronizeAPIRequest {
	return poolTaobaoQimenItemsSynchronizeAPIRequest.Get().(*TaobaoQimenItemsSynchronizeAPIRequest)
}

// ReleaseTaobaoQimenItemsSynchronizeAPIRequest 将 TaobaoQimenItemsSynchronizeAPIRequest 放入 sync.Pool
func ReleaseTaobaoQimenItemsSynchronizeAPIRequest(v *TaobaoQimenItemsSynchronizeAPIRequest) {
	v.Reset()
	poolTaobaoQimenItemsSynchronizeAPIRequest.Put(v)
}
