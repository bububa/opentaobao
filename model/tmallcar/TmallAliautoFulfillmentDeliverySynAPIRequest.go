package tmallcar

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallAliautoFulfillmentDeliverySynAPIRequest) GetApiMethodName() string {
	return "tmall.aliauto.fulfillment.delivery.syn"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallAliautoFulfillmentDeliverySynAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
