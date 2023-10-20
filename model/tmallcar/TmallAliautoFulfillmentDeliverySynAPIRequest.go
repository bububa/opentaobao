package tmallcar

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallAliautoFulfillmentDeliverySynAPIRequest 交付状态及物流信息同步 API请求
// tmall.aliauto.fulfillment.delivery.syn
//
// 交付状态及物流信息同步
type TmallAliautoFulfillmentDeliverySynAPIRequest struct {
	model.Params
	// 入参
	_req *SyncInfoReq
}

// NewTmallAliautoFulfillmentDeliverySynRequest 初始化TmallAliautoFulfillmentDeliverySynAPIRequest对象
func NewTmallAliautoFulfillmentDeliverySynRequest() *TmallAliautoFulfillmentDeliverySynAPIRequest {
	return &TmallAliautoFulfillmentDeliverySynAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallAliautoFulfillmentDeliverySynAPIRequest) Reset() {
	r._req = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallAliautoFulfillmentDeliverySynAPIRequest) GetApiMethodName() string {
	return "tmall.aliauto.fulfillment.delivery.syn"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallAliautoFulfillmentDeliverySynAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallAliautoFulfillmentDeliverySynAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReq is Req Setter
// 入参
func (r *TmallAliautoFulfillmentDeliverySynAPIRequest) SetReq(_req *SyncInfoReq) error {
	r._req = _req
	r.Set("req", _req)
	return nil
}

// GetReq Req Getter
func (r TmallAliautoFulfillmentDeliverySynAPIRequest) GetReq() *SyncInfoReq {
	return r._req
}

var poolTmallAliautoFulfillmentDeliverySynAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallAliautoFulfillmentDeliverySynRequest()
	},
}

// GetTmallAliautoFulfillmentDeliverySynRequest 从 sync.Pool 获取 TmallAliautoFulfillmentDeliverySynAPIRequest
func GetTmallAliautoFulfillmentDeliverySynAPIRequest() *TmallAliautoFulfillmentDeliverySynAPIRequest {
	return poolTmallAliautoFulfillmentDeliverySynAPIRequest.Get().(*TmallAliautoFulfillmentDeliverySynAPIRequest)
}

// ReleaseTmallAliautoFulfillmentDeliverySynAPIRequest 将 TmallAliautoFulfillmentDeliverySynAPIRequest 放入 sync.Pool
func ReleaseTmallAliautoFulfillmentDeliverySynAPIRequest(v *TmallAliautoFulfillmentDeliverySynAPIRequest) {
	v.Reset()
	poolTmallAliautoFulfillmentDeliverySynAPIRequest.Put(v)
}
