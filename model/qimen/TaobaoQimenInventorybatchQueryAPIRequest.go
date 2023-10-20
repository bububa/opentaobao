package qimen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenInventorybatchQueryAPIRequest 商品单仓批次库存查询接口 API请求
// taobao.qimen.inventorybatch.query
//
// ERP 通过该接口查询指定商品的单仓批次库存
type TaobaoQimenInventorybatchQueryAPIRequest struct {
	model.Params
	// request
	_request *TaobaoQimenInventorybatchQueryRequest
}

// NewTaobaoQimenInventorybatchQueryRequest 初始化TaobaoQimenInventorybatchQueryAPIRequest对象
func NewTaobaoQimenInventorybatchQueryRequest() *TaobaoQimenInventorybatchQueryAPIRequest {
	return &TaobaoQimenInventorybatchQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoQimenInventorybatchQueryAPIRequest) Reset() {
	r._request = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenInventorybatchQueryAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.inventorybatch.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenInventorybatchQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenInventorybatchQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
// request
func (r *TaobaoQimenInventorybatchQueryAPIRequest) SetRequest(_request *TaobaoQimenInventorybatchQueryRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenInventorybatchQueryAPIRequest) GetRequest() *TaobaoQimenInventorybatchQueryRequest {
	return r._request
}

var poolTaobaoQimenInventorybatchQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoQimenInventorybatchQueryRequest()
	},
}

// GetTaobaoQimenInventorybatchQueryRequest 从 sync.Pool 获取 TaobaoQimenInventorybatchQueryAPIRequest
func GetTaobaoQimenInventorybatchQueryAPIRequest() *TaobaoQimenInventorybatchQueryAPIRequest {
	return poolTaobaoQimenInventorybatchQueryAPIRequest.Get().(*TaobaoQimenInventorybatchQueryAPIRequest)
}

// ReleaseTaobaoQimenInventorybatchQueryAPIRequest 将 TaobaoQimenInventorybatchQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoQimenInventorybatchQueryAPIRequest(v *TaobaoQimenInventorybatchQueryAPIRequest) {
	v.Reset()
	poolTaobaoQimenInventorybatchQueryAPIRequest.Put(v)
}
