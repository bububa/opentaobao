package qimen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenSingleitemSynchronizeAPIRequest 商品同步接口 API请求
// taobao.qimen.singleitem.synchronize
//
// taobao.qimen.singleitem.synchronize
type TaobaoQimenSingleitemSynchronizeAPIRequest struct {
	model.Params
	//
	_request *ItemSynRequest
}

// NewTaobaoQimenSingleitemSynchronizeRequest 初始化TaobaoQimenSingleitemSynchronizeAPIRequest对象
func NewTaobaoQimenSingleitemSynchronizeRequest() *TaobaoQimenSingleitemSynchronizeAPIRequest {
	return &TaobaoQimenSingleitemSynchronizeAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoQimenSingleitemSynchronizeAPIRequest) Reset() {
	r._request = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenSingleitemSynchronizeAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.singleitem.synchronize"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenSingleitemSynchronizeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenSingleitemSynchronizeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoQimenSingleitemSynchronizeAPIRequest) SetRequest(_request *ItemSynRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenSingleitemSynchronizeAPIRequest) GetRequest() *ItemSynRequest {
	return r._request
}

var poolTaobaoQimenSingleitemSynchronizeAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoQimenSingleitemSynchronizeRequest()
	},
}

// GetTaobaoQimenSingleitemSynchronizeRequest 从 sync.Pool 获取 TaobaoQimenSingleitemSynchronizeAPIRequest
func GetTaobaoQimenSingleitemSynchronizeAPIRequest() *TaobaoQimenSingleitemSynchronizeAPIRequest {
	return poolTaobaoQimenSingleitemSynchronizeAPIRequest.Get().(*TaobaoQimenSingleitemSynchronizeAPIRequest)
}

// ReleaseTaobaoQimenSingleitemSynchronizeAPIRequest 将 TaobaoQimenSingleitemSynchronizeAPIRequest 放入 sync.Pool
func ReleaseTaobaoQimenSingleitemSynchronizeAPIRequest(v *TaobaoQimenSingleitemSynchronizeAPIRequest) {
	v.Reset()
	poolTaobaoQimenSingleitemSynchronizeAPIRequest.Put(v)
}
