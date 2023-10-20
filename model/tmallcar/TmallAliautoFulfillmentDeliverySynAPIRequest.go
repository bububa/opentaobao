package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallaliautofulfillmentdeliverysynAPIRequest 交付状态及物流信息同步 API请求
// tmall.aliauto.fulfillment.delivery.syn
//
// 交付状态及物流信息同步
type TmallaliautofulfillmentdeliverysynAPIRequest struct {
	model.Params
	// 入参
	_req *SyncInfoReq
}

// NewTmallaliautofulfillmentdeliverysynRequest 初始化TmallaliautofulfillmentdeliverysynAPIRequest对象
func NewTmallaliautofulfillmentdeliverysynRequest() *TmallaliautofulfillmentdeliverysynAPIRequest {
	return &TmallaliautofulfillmentdeliverysynAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallaliautofulfillmentdeliverysynAPIRequest) GetApiMethodName() string {
	return "tmall.aliauto.fulfillment.delivery.syn"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallaliautofulfillmentdeliverysynAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallaliautofulfillmentdeliverysynAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReq is Req Setter
// 入参
func (r *TmallaliautofulfillmentdeliverysynAPIRequest) SetReq(_req *SyncInfoReq) error {
	r._req = _req
	r.Set("req", _req)
	return nil
}

// GetReq Req Getter
func (r TmallaliautofulfillmentdeliverysynAPIRequest) GetReq() *SyncInfoReq {
	return r._req
}
