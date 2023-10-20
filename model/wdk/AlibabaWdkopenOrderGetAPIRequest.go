package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkopenordergetAPIRequest 五道口商户订单获取 API请求
// alibaba.wdkopen.order.get
//
// 商户通过五道口订单id获取订单信息
type AlibabawdkopenordergetAPIRequest struct {
	model.Params
	// 经营店id
	_storeId string
	// 外部主订单ID
	_outOrderId string
	// 五道口主订单id
	_bizOrderId int64
}

// NewAlibabawdkopenordergetRequest 初始化AlibabawdkopenordergetAPIRequest对象
func NewAlibabawdkopenordergetRequest() *AlibabawdkopenordergetAPIRequest {
	return &AlibabawdkopenordergetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkopenordergetAPIRequest) GetApiMethodName() string {
	return "alibaba.wdkopen.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkopenordergetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkopenordergetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreId is StoreId Setter
// 经营店id
func (r *AlibabawdkopenordergetAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r AlibabawdkopenordergetAPIRequest) GetStoreId() string {
	return r._storeId
}

// SetOutOrderId is OutOrderId Setter
// 外部主订单ID
func (r *AlibabawdkopenordergetAPIRequest) SetOutOrderId(_outOrderId string) error {
	r._outOrderId = _outOrderId
	r.Set("out_order_id", _outOrderId)
	return nil
}

// GetOutOrderId OutOrderId Getter
func (r AlibabawdkopenordergetAPIRequest) GetOutOrderId() string {
	return r._outOrderId
}

// SetBizOrderId is BizOrderId Setter
// 五道口主订单id
func (r *AlibabawdkopenordergetAPIRequest) SetBizOrderId(_bizOrderId int64) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// GetBizOrderId BizOrderId Getter
func (r AlibabawdkopenordergetAPIRequest) GetBizOrderId() int64 {
	return r._bizOrderId
}
