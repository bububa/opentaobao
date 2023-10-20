package qimen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenInventoryruleCreateAPIRequest 渠道间库存规则设置接口 API请求
// taobao.qimen.inventoryrule.create
//
// 渠道间库存规则设置
type TaobaoQimenInventoryruleCreateAPIRequest struct {
	model.Params
	//
	_request *RequestDo
}

// NewTaobaoQimenInventoryruleCreateRequest 初始化TaobaoQimenInventoryruleCreateAPIRequest对象
func NewTaobaoQimenInventoryruleCreateRequest() *TaobaoQimenInventoryruleCreateAPIRequest {
	return &TaobaoQimenInventoryruleCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoQimenInventoryruleCreateAPIRequest) Reset() {
	r._request = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenInventoryruleCreateAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.inventoryrule.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenInventoryruleCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenInventoryruleCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoQimenInventoryruleCreateAPIRequest) SetRequest(_request *RequestDo) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenInventoryruleCreateAPIRequest) GetRequest() *RequestDo {
	return r._request
}

var poolTaobaoQimenInventoryruleCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoQimenInventoryruleCreateRequest()
	},
}

// GetTaobaoQimenInventoryruleCreateRequest 从 sync.Pool 获取 TaobaoQimenInventoryruleCreateAPIRequest
func GetTaobaoQimenInventoryruleCreateAPIRequest() *TaobaoQimenInventoryruleCreateAPIRequest {
	return poolTaobaoQimenInventoryruleCreateAPIRequest.Get().(*TaobaoQimenInventoryruleCreateAPIRequest)
}

// ReleaseTaobaoQimenInventoryruleCreateAPIRequest 将 TaobaoQimenInventoryruleCreateAPIRequest 放入 sync.Pool
func ReleaseTaobaoQimenInventoryruleCreateAPIRequest(v *TaobaoQimenInventoryruleCreateAPIRequest) {
	v.Reset()
	poolTaobaoQimenInventoryruleCreateAPIRequest.Put(v)
}
