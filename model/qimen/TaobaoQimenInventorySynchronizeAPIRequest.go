package qimen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenInventorySynchronizeAPIRequest 库存状态同步接口 API请求
// taobao.qimen.inventory.synchronize
//
// ERP通过该接口同步指定商品的库存信息
type TaobaoQimenInventorySynchronizeAPIRequest struct {
	model.Params
	//
	_request *TaobaoQimenInventorySynchronizeRequest
}

// NewTaobaoQimenInventorySynchronizeRequest 初始化TaobaoQimenInventorySynchronizeAPIRequest对象
func NewTaobaoQimenInventorySynchronizeRequest() *TaobaoQimenInventorySynchronizeAPIRequest {
	return &TaobaoQimenInventorySynchronizeAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoQimenInventorySynchronizeAPIRequest) Reset() {
	r._request = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenInventorySynchronizeAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.inventory.synchronize"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenInventorySynchronizeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenInventorySynchronizeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoQimenInventorySynchronizeAPIRequest) SetRequest(_request *TaobaoQimenInventorySynchronizeRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenInventorySynchronizeAPIRequest) GetRequest() *TaobaoQimenInventorySynchronizeRequest {
	return r._request
}

var poolTaobaoQimenInventorySynchronizeAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoQimenInventorySynchronizeRequest()
	},
}

// GetTaobaoQimenInventorySynchronizeRequest 从 sync.Pool 获取 TaobaoQimenInventorySynchronizeAPIRequest
func GetTaobaoQimenInventorySynchronizeAPIRequest() *TaobaoQimenInventorySynchronizeAPIRequest {
	return poolTaobaoQimenInventorySynchronizeAPIRequest.Get().(*TaobaoQimenInventorySynchronizeAPIRequest)
}

// ReleaseTaobaoQimenInventorySynchronizeAPIRequest 将 TaobaoQimenInventorySynchronizeAPIRequest 放入 sync.Pool
func ReleaseTaobaoQimenInventorySynchronizeAPIRequest(v *TaobaoQimenInventorySynchronizeAPIRequest) {
	v.Reset()
	poolTaobaoQimenInventorySynchronizeAPIRequest.Put(v)
}
